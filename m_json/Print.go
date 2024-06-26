package m_json

import (
	"log"
)

// struct 进行 json 格式 打印
func Println(data any) string {
	json := ToJson(data)
	jsonStr := JsonFormat(json)
	log.Println(jsonStr)
	return jsonStr
}

// byte 进行 json 格式 打印
func PrintlnForByte(jsonByte []byte) string {
	jsonStr := JsonFormat(jsonByte)
	log.Println(jsonStr)
	return jsonStr
}
