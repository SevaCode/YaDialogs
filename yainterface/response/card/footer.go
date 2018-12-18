package card // import "github.com/SevaCode/YaDialogs/yainterface/response/card"

// Кнопки под галереей изображений.
type Footer struct {
	// Текст первой кнопки, обязательное свойство. Максимум 64 символа.
	Text string `json:"text"`

	// Дополнительная кнопка для галереи изображений.
	Button Button `json:"button"`
}
