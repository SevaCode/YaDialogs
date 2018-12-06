package request // import "github.com/SevaCode/YaDialogs/yainterface/request"

// Формальные характеристики реплики, которые удалось выделить Яндекс.Диалогам.
// Отсутствует, если ни одно из вложенных свойств не применимо.
type Markup struct {
	// Признак реплики, которая содержит криминальный подтекст (самоубийство,
	// разжигание ненависти, угрозы). Вы можете настроить навык на
	// определенную реакцию для таких случаев — например, отвечать "Не
	// понимаю, о чем вы. Пожалуйста, переформулируйте вопрос."
	//
	// Возможно только значение true. Если признак не применим, это свойство
	// не включается в ответ.
	DangerousContext bool `json:"dangerous_context,omitempty"`
}
