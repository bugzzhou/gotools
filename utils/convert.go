package utils

import "math/rand"

// 判断string是否存在于指定切片中
func IsInSlice(in string, target []string) bool {
	for _, v := range target {
		if in == v {
			return true
		}
	}
	return false
}

// 随机从指定切片中取出n个元素
func RandomSelect(slice []string, n int) ([]string, []string) {
	if n > len(slice) {
		return nil, nil
	}

	selected := make([]string, n)
	copy(selected, slice)

	// 打乱选中的元素顺序
	rand.Shuffle(len(selected), func(i, j int) {
		selected[i], selected[j] = selected[j], selected[i]
	})
	return selected[:n], selected[n:]
}
