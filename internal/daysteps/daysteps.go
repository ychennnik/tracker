package daysteps

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/Yandex-Practicum/tracker/internal/personaldata"
	"github.com/Yandex-Practicum/tracker/internal/spentenergy"
)

// DaySteps хранит данные о дневной прогулке и о пользователе.
type DaySteps struct {
	Steps    int           // количество шагов
	Duration time.Duration // длительность прогулки
	personaldata.Personal
}

// Parse разбирает строку формата "678,0h50m" и записывает данные в ds.
// Поля структуры меняются только если вся строка разобрана без ошибок.
func (ds *DaySteps) Parse(datastring string) (err error) {
	parts := strings.Split(datastring, ",")
	if len(parts) != 2 {
		return fmt.Errorf("неверный формат данных %q: ожидалось 2 значения, получено %d", datastring, len(parts))
	}

	steps, err := strconv.Atoi(parts[0])
	if err != nil {
		return fmt.Errorf("не удалось разобрать количество шагов: %w", err)
	}
	if steps <= 0 {
		return errors.New("количество шагов должно быть больше 0")
	}

	duration, err := time.ParseDuration(parts[1])
	if err != nil {
		return fmt.Errorf("не удалось разобрать продолжительность: %w", err)
	}
	if duration <= 0 {
		return errors.New("продолжительность должна быть больше 0")
	}

	ds.Steps = steps
	ds.Duration = duration

	return nil
}

// ActionInfo возвращает строку с информацией о прогулке.
func (ds DaySteps) ActionInfo() (string, error) {
	distance := spentenergy.Distance(ds.Steps, ds.Height)

	calories, err := spentenergy.WalkingSpentCalories(ds.Steps, ds.Weight, ds.Height, ds.Duration)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"Количество шагов: %d.\nДистанция составила %.2f км.\nВы сожгли %.2f ккал.\n",
		ds.Steps, distance, calories,
	), nil
}
