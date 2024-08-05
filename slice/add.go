package slice

import "github.com/zpapple77/gokit/internal/slice"

/*
Add 在 index 处添加元素
index 范围应为[0,len(src)]
如果index == len(src)，则添加到末尾
*/

func Add[Src any](src []Src, element Src, index int) ([]Src, error) {
	res, err := slice.Add(src, element, index)
	return res, err
}
