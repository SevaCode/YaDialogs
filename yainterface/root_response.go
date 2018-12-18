package yainterface // import "github.com/SevaCode/YaDialogs/yainterface"

// Формат ответа
type RootResponse struct {

	// Данные для ответа пользователю.
	Response Response `json:"response"`

	// Данные о сессии.
	Session Session `json:"session"`

	// Версия протокола. Текущая версия — 1.0.
	Version string `json:"version"`
}
