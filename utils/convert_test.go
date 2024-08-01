package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsInSlice(t *testing.T) {
	expect1 := true
	actual1 := IsInSlice("a", []string{"a", "b", "c"})
	assert.Equal(t, expect1, actual1)

	expect2 := false
	actual2 := IsInSlice("d", []string{"a", "b", "c"})
	assert.Equal(t, expect2, actual2)
}

func TestRandSelect(t *testing.T) {
	raw := []string{"a", "b", "c"}
	expect1 := permutations(raw, 2)
	actual1, _ := RandomSelect(raw, 2)
	assert.Contains(t, expect1, actual1)

	var expect2 []string = nil
	actual2, _ := RandomSelect(raw, 20)
	assert.Equal(t, expect2, actual2)
}

//
//
//
//
//
//
//
//-----------------------------------------------------------------------------------
//测试函数自身所包含的函数，不会用于其它正式逻辑
//
//
//
//
//

func TestPer(t *testing.T) {
	raw := []string{"a", "b", "c"}
	expect := [][]string{
		{"a", "b"},
		{"a", "c"},
		{"b", "a"},
		{"b", "c"},
		{"c", "a"},
		{"c", "b"},
	}
	actual := permutations(raw, 2)

	assert.Equal(t, expect, actual)
}

// 仅用于产生  []string 中 n个不重复元素的全排列。
// 用户assert.Contains中expect部分的产生
func permutations(slice []string, n int) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	var result [][]string
	for i, _ := range slice {
		// 从切片中取出一个元素作为当前元素
		current := slice[i]

		// 从剩余元素中生成排列组合
		remaining := append([]string(nil), slice[:i]...)
		remaining = append(remaining, slice[i+1:]...)
		perms := permutations(remaining, n-1)

		// 将当前元素与每个排列组合合并
		for _, perm := range perms {
			result = append(result, append([]string{current}, perm...))
		}
	}

	return result
}
