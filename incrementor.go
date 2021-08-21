//Package incrementor
/*
	Реализует инкрементор числа
	Создаётся функцией New()

	Содержит методы:
	- GetNumber()		- получить число из инкрментора
	- IncrementNumber()	- увеличить число в инкременторе
	- SetMaximumValue()	- Установить максимально значение числа (по умолчанию - max int)

	При достижении максимального значения инкрементор сбрасывается на 0
*/
package incrementor

import (
	"errors"
)

const (
	DefaultMaxValue = 1<<(sizeInt-1) - 1 // Максимальное значение по умолчанию
)

// New - создаёт экземпляр инкрементора
func New() *Incrementor {
	return &Incrementor{maxNum: DefaultMaxValue}
}

// SetMaximumValue - Задать максимальное значение инкрементора
// На вход идёт любое число в диапазоне [0..maxInt]
func (inc *Incrementor) SetMaximumValue(maxValue int) error {
	if maxValue < 0 {
		return errors.New("error! The maximum value must not be lower 0")
	}
	inc.maxNum = maxValue
	return nil
}

// GetNumber - получить текущее значение инкрементора
func (inc *Incrementor) GetNumber() int {
	return inc.num
}

// IncrementNumber - Увеличить значение инкрементора на 1
func (inc *Incrementor) IncrementNumber() {
	inc.num++
	if inc.num >= inc.maxNum {
		inc.num = 0
	}
}
