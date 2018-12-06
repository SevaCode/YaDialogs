package yainterface // import "github.com/SevaCode/YaDialogs/yainterface"

import "github.com/SevaCode/YaDialogs/yainterface/meta"

// Meta содержит информацию об устройстве, с помощью которого пользователь
// разговаривает с Алисой.
type Meta struct {

	// Язык в POSIX-формате.
	Locale string `json:"locale"`

	// Название часового пояса, включая алиасы.
	TimeZone string `json:"timezone"`

	// Идентификатор устройства и приложения, в котором идет разговор,
	// максимум 1024 символа.
	//
	// __Примечание__. Не рекомендуется к использованию.
	// Интерфейсы, доступные на клиентском устройстве,
	// перечислены в свойстве interfaces.
	ClientID string `json:"client_id"`

	// Интерфейсы, доступные на устройстве пользователя.
	Interfaces meta.Interfaces `json:"interfaces"`
}
