package yainterface // import "github.com/SevaCode/YaDialogs/yainterface"

// Формат запроса
type RootRequest struct {

	// Свойства с метаинформацией о запросе
	//
	// Информация об устройстве, с помощью которого пользователь
	// разговаривает с Алисой.
	Meta Meta `json:"meta"`

	// Данные о сессии.
	Session Session `json:"session"`

	// Версия протокола. Текущая версия — 1.0.
	Version string `json:"version"`

	// Содержимое запроса к навыку

	// Данные, полученные от пользователя.
	Request Request `json:"request"`
}
