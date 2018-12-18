package card // import "github.com/SevaCode/YaDialogs/yainterface/response/card"

// Заголовок галереи изображений.
// Игнорируется для карточки типа BigImage.
type Header struct {
	// Текст заголовка, обязателен, если передается свойство header.
	// Максимум 64 символа.
	Text string `json:"text"`
}
