package nlu // import "github.com/SevaCode/YaDialogs/yainterface/request/nlu"

import "github.com/SevaCode/YaDialogs/yainterface/request/nlu/entity"

// Именованная сущность.
type Entity struct {

	// Первое слово именованной сущности.
	Tokens entity.Tokens `json:"tokens"`

	// Тип именованной сущности. Возможные значения:
	// YANDEX.DATETIME — дата и время, абсолютные или относительные.
	// YANDEX.FIO — фамилия, имя и отчество.
	// YANDEX.GEO — местоположение (адрес или аэропорт).
	// YANDEX.NUMBER — число, целое или с плавающей точкой.
	Type entity.Type `json:"type"`

	// Формальное описание именованной сущности.
	// Формат этого поля для всех поддерживаемых типов сущностей приведен в
	// разделе Именованные сущности в запросах.
	// https://tech.yandex.ru/dialogs/alice/doc/nlu-docpage/
	Value entity.Value `json:"value"`
}
