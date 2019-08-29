package convertx

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// timestamp accurate to second
func Timestamp2FormatStr(timestamp int64, format string) string {
	return time.Unix(timestamp, 0).Format(format)
}

func Str2Int64(s string) (int64, error) {
	return strconv.ParseInt(s, 10, 64)
}

func Str2Float64(s string) (float64, error) {
	return strconv.ParseFloat(s, 64)
}

/*
hello_world to helloWorld/HelloWorld
includeFirst 是否有包含首字母
*/
func Snake2Hump(s string, includeFirst bool) string {
	list := strings.Split(s, "_")
	for k := range list {
		if !includeFirst && k == 0 {
			continue
		}
		list[k] = strings.ToUpper(string(list[k][0])) + string(list[k][1:])
	}
	return strings.Join(list, "")
}

/*
HelloWorld to hello_world
*/
func Hump2Snake(s string) string {
	list := FindUpperWord(s)
	for k := range list {
		s = strings.ReplaceAll(s, list[k], "_"+strings.ToLower(list[k]))
	}
	return strings.Trim(s, "_")
}

func FindUpperWord(s string) []string {
	b := []byte(s)
	res := []string{}
	for k := range b {
		if b[k] >= 65 && b[k] <= 90 {
			res = append(res, string(b[k]))
		}
	}
	return res
}

/*
,a,,,,b,,c, to [a,b,c]
*/
func SeniorSplit(s, sub string) []string {
	s = strings.Trim(s, sub)
	res := []string{}
	list := strings.Split(s, sub)
	for k := range list {
		if list[k] != "" {
			res = append(res, list[k])
		}
	}
	return res
}

func GoModel2proto3(i interface{}) string {
	iType := reflect.TypeOf(i)
	iValue := reflect.ValueOf(i)
	res := ""
	format := "   %s %s = %d;\n"
	index := 1
	for k := 0; k < iValue.NumField(); k++ {
		pType := goTypeConvert2proto3(iType.Field(k).Type.String())
		pName := Hump2Snake(iType.Field(k).Name)
		res += fmt.Sprintf(format, pType, pName, index)
		index++
	}
	return fmt.Sprintf("message %s {\n%s}", iType.Name(), res)
}

func goTypeConvert2proto3(gtype string) string {
	switch gtype {
	case "string":
		return "string"
	case "[]byte", "json.RawMessage":
		return "bytes"
	case "int", "int8", "uint8", "int16", "int32", "uint32":
		return "int32"
	case "int64":
		return "int64"
	default:
		return ""
	}
}