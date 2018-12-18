package response // import "github.com/SevaCode/YaDialogs/yainterface/response"

import "github.com/SevaCode/YaDialogs/yainterface/response/card"

// Описание карточки — сообщения с поддержкой изображений.
// Если приложению удается отобразить карточку для пользователя,
// свойство response.text не используется.
type Card struct {
	// Тип карточки. Поддерживаемые значения:
	// * BigImage — одно изображение.
	// * ItemsList — галерея изображений (от 1 до 5).
	// Требуемый формат ответа зависит от типа карточки.
	Type card.Type `json:"type"`

	// Идентификатор изображения, который возвращается в ответ на запрос загрузки.
	// https://tech.yandex.ru/dialogs/alice/doc/resource-upload-docpage/
	// Необходимо указывать для карточки типа BigImage, для типа ItemsList игнорируется.
	ImageId string `json:"image_id,omitempty"`

	// Заголовок для изображения.
	// Игнорируется для карточки типа ItemsList.
	Title string `json:"title,omitempty"`

	// Описание изображения.
	// Игнорируется для карточки типа ItemsList.
	Description string `json:"description,omitempty"`

	// Свойства изображения, на которое можно нажать.
	// Игнорируется для карточки типа ItemsList.
	Button card.Button `json:"button,omitempty"`

	// Заголовок галереи изображений.
	// Игнорируется для карточки типа BigImage.
	Header card.Header `json:"header,omitempty"`

	// Набор изображений для галереи — не меньше 1, не больше 5.
	// Игнорируется для карточки типа BigImage.
	Items []card.Item `json:"items,omitempty"`

	// Кнопки под галереей изображений.
	Footer card.Footer `json:"footer,omitempty"`
}
