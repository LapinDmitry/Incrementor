package incrementor

import (
	"unsafe"
)

const (
	sizeInt = unsafe.Sizeof(int(0)) * 8 // Размер инта в текущей системе в битах
)

type Incrementor struct {
	num    int
	maxNum int
}
