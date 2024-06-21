package m_file

import (
	"bytes"
	"io"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/handy-golang/go-tools/m_encrypt"
	"github.com/handy-golang/go-tools/m_path"
)

type DownFileOpt struct {
	Url      string
	SavePath string
	SaveName string
}

func DownFile(opt DownFileOpt) (resData string, resErr error) {
	Url := opt.Url

	SavePath := opt.SavePath
	if len(SavePath) < 1 {
		SavePath = "."
	}
	isLogPath := m_path.Exists(SavePath)
	if !isLogPath {
		os.Mkdir(SavePath, 0o777)
	}

	SavePath, _ = filepath.Abs(SavePath)

	SaveName := opt.SaveName
	if len(SaveName) < 2 {
		SaveName = GetSaveFileName(GetNameOpt{
			FileName: m_encrypt.TimeID(),
			SavePath: SavePath,
			RandName: true,
		})
	}

	c := colly.NewCollector()
	c.WithTransport(&http.Transport{
		Proxy: http.ProxyFromEnvironment,
		DialContext: (&net.Dialer{
			Timeout: 100 * time.Second,
		}).DialContext,
		MaxIdleConns:          0,
		IdleConnTimeout:       0,
		TLSHandshakeTimeout:   0,
		ExpectContinueTimeout: 0,
	})

	c.OnResponse(func(r *colly.Response) {
		fileName := SaveName
		extName := path.Ext(SaveName) // 后缀名
		if len(extName) < 1 {
			extName = ContentToExtName(r.Headers.Get("Content-Type"))
			if len(extName) > 0 {
				extName = "." + extName
			}
			fileName = fileName + extName
		}

		SaveFile := SavePath + "/" + fileName
		f, err := os.Create(SaveFile)
		if err != nil {
			resErr = err
		}
		io.Copy(f, bytes.NewReader(r.Body))

		resData = SaveFile
	})
	c.OnError(func(r *colly.Response, err error) {
		if err != nil {
			resErr = err
		}
	})
	c.Visit(Url)

	return
}
