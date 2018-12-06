package entities // import "github.com/SevaCode/YaDialogs/yainterface/request/nlu/entities"

import "strings"

// Тип именованной сущности. Возможные значения:
// YANDEX.DATETIME — дата и время, абсолютные или относительные.
// YANDEX.FIO — фамилия, имя и отчество.
// YANDEX.GEO — местоположение (адрес или аэропорт).
// YANDEX.NUMBER — число, целое или с плавающей точкой.
type Type int

const (
	YANDEX_DATETIME Type = iota
	YANDEX_FIO
	YANDEX_GEO
	YANDEX_NUMBER
)

var types = [...]string{
	"YANDEX.DATETIME",
	"YANDEX.FIO",
	"YANDEX.GEO",
	"YANDEX.NUMBER",
}

func (t *Type) String() string {
	return types[*t]
}

func (t *Type) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case "SYANDEX.DATETIME":
		*t = YANDEX_DATETIME

	case "YANDEX.FIO":
		*t = YANDEX_FIO

	case "SYANDEX.GEO":
		*t = YANDEX_GEO

	case "YANDEX.NUMBER":
		*t = YANDEX_NUMBER
	}

	return nil
}
