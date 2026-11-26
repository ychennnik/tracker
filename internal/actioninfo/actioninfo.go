package actioninfo

import (
	"fmt"
	"log"
)

// DataParser описывает поведение, общее для Training и DaySteps:
// разбор строки с данными и формирование строки с информацией.
type DataParser interface {
	Parse(datastring string) error
	ActionInfo() (string, error)
}

// Info перебирает записи из dataset, разбирает каждую через dp.Parse
// и печатает результат dp.ActionInfo. Ошибки логируются, а обработка
// продолжается со следующей записи.
func Info(dataset []string, dp DataParser) {
	for _, data := range dataset {
		if err := dp.Parse(data); err != nil {
			log.Printf("ошибка разбора данных %q: %v", data, err)
			continue
		}

		info, err := dp.ActionInfo()
		if err != nil {
			log.Printf("ошибка формирования информации для %q: %v", data, err)
			continue
		}

		fmt.Println(info)
	}
}
