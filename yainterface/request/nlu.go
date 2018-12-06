package request // import "github.com/SevaCode/YaDialogs/yainterface/request"

import "github.com/SevaCode/YaDialogs/yainterface/request/nlu"

// Слова и именованные сущности, которые Диалоги извлекли из
// запроса пользователя.
// Подробное описание поддерживаемых типов сущностей
// см. в разделе Именованные сущности в запросах.
// https://tech.yandex.ru/dialogs/alice/doc/nlu-docpage/
type Nlu struct {

	// Массив слов из произнесенной пользователем фразы.
	Tokens []string `json:"tokens"`

	// Массив именованных сущностей.
	Entities []nlu.Entity `json:"entities"`
}
