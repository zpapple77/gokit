package slice

import "github.com/zpapple77/gokit/internal/errs"

func Add[T any](src []T, element T, index int) ([]T, error) {
	length := len(src)
	if index < 0 || index > length {
		return nil, errs.NewErrIndexOutOfRange(length, index)
	}

	//先将src扩展一个元素
	var zeroValue T
	src = append(src, zeroValue)
	for i := length; i > index; i-- {
		src[i] = src[i-1]
	}
	src[index] = element
	return src, nil
}
