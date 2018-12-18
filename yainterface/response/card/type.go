package card // import "github.com/SevaCode/YaDialogs/yainterface/response/card"

import "strings"

// Тип карточки. Поддерживаемые значения:
// * BigImage — одно изображение.
// * ItemsList — галерея изображений (от 1 до 5).
// Требуемый формат ответа зависит от типа карточки.
type Type int

const (
	BigImage Type = iota
	ItemsList
)

var types = [...]string{
	"BigImage",
	"ItemsList",
}

func (t *Type) String() string {
	return types[*t]
}

func (t *Type) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case "BigImage":
		*t = BigImage

	case "ItemsList":
		*t = ItemsList
	}

	return nil
}
