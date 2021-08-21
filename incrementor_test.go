package incrementor_test

import (
	"github.com/LapinDmitry/Incrementor"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

// Тест получения числа
func TestIncrementor_GetNumber(t *testing.T) {
	// Создадим объект
	inc := incrementor.New()

	// Получим число (прогнозируемое значение - 0)
	want := 0
	actual := inc.GetNumber()

	// Вывод и проверка
	t.Log("Actual: ", actual)
	require.Equal(t, want, actual)
}

// Тест инкремента
func TestIncrementor_IncrementNumber(t *testing.T) {
	inc := incrementor.New()

	// После вызова метода инкремента, начально значение должно увеличиться на один
	want1 := 0
	want2 := want1 + 1

	// Производим действия
	actual1 := inc.GetNumber()
	inc.IncrementNumber()
	actual2 := inc.GetNumber()

	// Вывод и проверка
	t.Log("Actual1: ", actual1)
	t.Log("Actual2: ", actual2)
	require.Equal(t, want1, actual1)
	require.Equal(t, want2, actual2)
}

// Проверка установки максимального значения
func TestIncrementor_SetMaximumValue(t *testing.T) {
	inc := incrementor.New()

	// После установки максимального значения 2 и двух инкрементов, значение должно сброситься до нуля
	want := 0
	argSet := 2

	// Установим значение (аргумент не должен быть отрицательным)
	err := inc.SetMaximumValue(argSet)
	assert.NoError(t, err)

	// Два инкремента
	inc.IncrementNumber()
	inc.IncrementNumber()

	// Берём число
	actual := inc.GetNumber()

	// Вывод и проверка
	t.Log("Actual: ", actual)
	assert.Equal(t, want, actual)
}

// Комплексные тесты
func TestIncrementor_ComplexTests(t *testing.T) {
	tests := []struct {
		countInc1 int  // Количество инкрементов в 1 раз
		set       int  // Установка максимума
		wantErr   bool // Прогнозируемое наличие ошибки
		countInc2 int  // Количество инкрементов в 2 раз
		result    int  // Прогнозируемый результат
	}{
		{0, 10, false, 5, 5},   // Просто 5 инкрементов
		{0, 5, false, 5, 0},    // Переполнение
		{0, 5, false, 7, 2},    // Инкремент после переполнения
		{0, 5, false, 12, 2},   // Двойное переполнение
		{0, 0, false, 100, 0},  // Максимально значение - 0
		{20, 10, false, 0, 20}, // После установки максимального, число не меняется без инкремента, даже если оно больше нового максимума
		{20, 10, false, 1, 0},  // Инкремент после установки максимального, число сразу сбрасывается
		{2, -10, true, 6, 8},   // Отрицательный аргумент в максимальном
		//{0, incrementor.DefaultMaxValue,false,incrementor.DefaultMaxValue,0},	// Проверка сброса на максимальных значениях
	}

	for _, test := range tests {
		inc := incrementor.New()

		// Первые инкременты
		for i := 0; i < test.countInc1; i++ {
			inc.IncrementNumber()
		}

		// Установка максимума
		err := inc.SetMaximumValue(test.set)
		assert.Equal(t, test.wantErr, err != nil)

		// Вторые инкременты
		for i := 0; i < test.countInc2; i++ {
			inc.IncrementNumber()
		}

		// Конечный результат
		result := inc.GetNumber()
		assert.Equal(t, test.result, result)

		// Вывод
		t.Logf("Actual Error:%v Result:%d", err, result)
	}
}
