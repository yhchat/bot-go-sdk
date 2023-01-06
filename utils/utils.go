package utils

import (
	"encoding/json"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

//interface转json byte
func InterfaceToJsonBytes(v interface{}) []byte {
	var json = jsoniter.ConfigCompatibleWithStandardLibrary
	contentBytes, _ := json.Marshal(v)
	return contentBytes
}

//interface转string
func InterfaceToString(unk interface{}) string {
	if unk == nil {
		return ""
	}
	switch i := unk.(type) {
	case string:
		return i
	case float64:
		return strconv.FormatFloat(i, 'f', -1, 64)
	case float32:
		return strconv.FormatFloat(float64(i), 'f', -1, 32)
	case int64:
		return strconv.FormatInt(i, 10)
	case int:
		return strconv.Itoa(i)
	default:
		return ""
	}
}

//interface转int64
func InterfaceToInt64(i interface{}) int64 {
	if i == nil {
		return 0
	}
	switch v := i.(type) {
	case float64:
		return int64(v)
	case float32:
		return int64(v)
	case int64:
		return v
	case int:
		return int64(v)
	case string:
		r, _ := strconv.ParseInt(v, 10, 64)
		return r
	default:
		return 0
	}
}

//json字符串转map
func JsonToMap(str string) (m map[string]interface{}) {
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		return
	}
	return
}
