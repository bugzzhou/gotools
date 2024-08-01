package test

import (
	"github.com/tidwall/gjson"
)

func GetString(data, match string) string {
	raw := gjson.Get(data, match)
	return raw.String()
}

func GetInt(data, match string) int64 {
	raw := gjson.Get(data, match)
	return raw.Int()
}

func Get() {
	return
}
