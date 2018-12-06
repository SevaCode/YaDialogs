package request // import "github.com/SevaCode/YaDialogs/yainterface/request"

import "strings"

// Тип ввода, обязательное свойство. Возможные значения:
// * SimpleUtterance — голосовой ввод;
// * ButtonPressed — нажатие кнопки.
type Type int

const (
	SimpleUtterance Type = iota
	ButtonPressed
)

var types = [...]string{
	"SimpleUtterance",
	"ButtonPressed",
}

func (t *Type) String() string {
	return types[*t]
}

func (t *Type) UnmarshalJSON(b []byte) error {
	str := strings.Trim(string(b), `"`)

	switch str {
	case "SimpleUtterance":
		*t = SimpleUtterance

	case "ButtonPressed":
		*t = ButtonPressed
	}

	return nil
}
