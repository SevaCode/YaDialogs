package entity // import "github.com/SevaCode/YaDialogs/yainterface/request/nlu/entity"

// Обозначение начала и конца именованной сущности в массиве слов.
// Нумерация слов в массиве начинается с 0.
type Tokens struct {

	// Первое слово именованной сущности.
	Start uint16 `json:"start"`

	// Первое слово после именованной сущности.
	End uint16 `json:"end"`
}
