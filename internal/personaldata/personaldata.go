package personaldata

import "fmt"

// Personal хранит персональные данные пользователя.
// Структура встраивается в Training и DaySteps.
type Personal struct {
	Name   string
	Weight float64 // вес в килограммах
	Height float64 // рост в метрах
}

// Print выводит данные пользователя на экран.
func (p Personal) Print() {
	fmt.Printf("Имя: %s\nВес: %.2f\nРост: %.2f\n", p.Name, p.Weight, p.Height)
}
