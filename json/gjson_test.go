package test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"
)

const rawJson = `
{
    "name": "jszhou",
    "age": 20,
    "nickname": ["a","b","c","d"],
    "friends": [
        {
            "name": "boy1",
            "age": 20,
            "sex": "m"
        },
        {
            "name": "boy2",
            "age": 22,
            "sex": "m"
        },
        {
            "name": "girl1",
            "age": 21,
            "sex": "w"
        }
    ],
    "add": {
        "country": "china",
        "city": "sz",
        "town": "dongzhu"
    }
}
`

func TestGetInt(t *testing.T) {
	var expect int64 = 20
	actual := GetInt(rawJson, "age")
	assert.Equal(t, expect, actual)

	expect = 22
	actual = GetInt(rawJson, "friends.1.age")
	assert.Equal(t, expect, actual)
}

func TestGetString(t *testing.T) {
	expect := "jszhou"
	actual := GetString(rawJson, "name")
	assert.Equal(t, expect, actual)

	expect = "girl1"
	actual = GetString(rawJson, "friends.2.name")
	assert.Equal(t, expect, actual)
}

func TestGetMany(t *testing.T) {
	aa := gjson.Get(rawJson, "friends.#.name")

	fmt.Println(aa)
	fmt.Println(aa.Type)

	bb := gjson.Get(rawJson, "add")
	fmt.Println(bb)
	fmt.Println(bb.Type)
	fmt.Println(bb.Map())

	cc := gjson.Get(rawJson, `friends.#(name%"boy*")`)
	fmt.Println(cc)

	dd := gjson.Get(rawJson, `friends.#(name*,age)`).Array()
	fmt.Println(dd)
}
