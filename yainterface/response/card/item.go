package card // import "github.com/SevaCode/YaDialogs/yainterface/response/card"

// Набор изображений для галереи — не меньше 1, не больше 5.
// Игнорируется для карточки типа BigImage.
type Item struct {
	// Идентификатор изображения, который возвращается в ответ на запрос загрузки.
	// https://tech.yandex.ru/dialogs/alice/doc/resource-upload-docpage/
	ImageId string `json:"image_id"`

	// Заголовок для изображения.
	Title string `json:"title"`

	// Описание изображения.
	Description string `json:"description"`

	// Свойства изображения, на которое можно нажать.
	Button Button `json:"button"`
}
