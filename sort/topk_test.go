package sort

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestTopK(t *testing.T) {
	a := []int{}

	for i := 0; i < 500; i++ { // 随机生成 500 个元素, 从 0 到 50000
		a = append(a, rand.Intn(50000))
	}

	top := TopK(a, 10) // 获取元素中最大的 10 个

	fmt.Println(top)
}
