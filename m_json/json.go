package m_json

import (
	"bytes"
	"encoding/json"

	"github.com/handy-golang/go-tools/m_str"
	jsoniter "github.com/json-iterator/go"
)

// struct 转成 json byte
func ToJson(data any) []byte {
	jsonStr, err := jsoniter.Marshal(data)
	if err != nil {
		return []byte{}
	}
	return jsonStr
}

// struct 转成 json string
func ToStr(data any) string {
	byteStr := ToJson(data)
	jsonStr := m_str.ToStr(byteStr)
	return jsonStr
}

// json byte 转成 json string 并格式化
func JsonFormat(jsonByte []byte) string {
	var out bytes.Buffer
	json.Indent(&out, jsonByte, "", " ")
	return out.String()
}

// struct 转成 json string 并格式化
func Format(data any) string {
	jsonByte := ToJson(data)
	return JsonFormat(jsonByte)
}

// struct 转成 map
func StructToMap(val any) (resData map[string]any) {
	jsonByte := ToJson(val)
	err := jsoniter.Unmarshal(jsonByte, &resData)
	if err != nil {
		resData = map[string]any{}
	}
	return
}
