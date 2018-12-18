package yainterface // import "github.com/SevaCode/YaDialogs/yainterface"

import (
	"github.com/SevaCode/YaDialogs/yainterface/common"
	"github.com/SevaCode/YaDialogs/yainterface/request"
)

// Request содержит данные, полученные от пользователя.
type Request struct {

	// Запрос, который был передан вместе с командой активации навыка.
	Command string `json:"command"`

	// Полный текст пользовательского запроса, максимум 1024 символа.
	OriginalUtterance string `json:"original_utterance"`

	// Тип ввода, обязательное свойство. Возможные значения:
	// * SimpleUtterance — голосовой ввод;
	// * ButtonPressed — нажатие кнопки.
	Type request.Type `json:"type"`

	// Формальные характеристики реплики, которые удалось выделить Яндекс.Диалогам.
	// Отсутствует, если ни одно из вложенных свойств не применимо.
	Markup request.Markup `json:"markup,omitempty"`

	// JSON, полученный с нажатой кнопкой от обработчика навыка
	// (в ответе на предыдущий запрос), максимум 4096 байт.
	Payload common.Payload `json:"payload,omitempty"`

	// Слова и именованные сущности, которые Диалоги извлекли
	// из запроса пользователя.
	// Подробное описание поддерживаемых типов сущностей
	// см. в разделе Именованные сущности в запросах.
	// https://tech.yandex.ru/dialogs/alice/doc/nlu-docpage/
	Nlu request.Nlu `json:"nlu"`
}
