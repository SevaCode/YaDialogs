package yainterface // import "github.com/SevaCode/YaDialogs/yainterface"

// Session содержит данные о сессии.
type Session struct {

	// Признак новой сессии. Возможные значения:
	// * true — пользователь начал новый разговор с навыком;
	// * false — запрос отправлен в рамках уже начатого разговора.
	New bool `json:"new"`

	// Идентификатор сообщения в рамках сессии, максимум 8 символов.
	// Инкрементируется с каждым следующим запросом.
	MessageID int64 `json:"message_id"`

	// Уникальный идентификатор сессии, 64 байта.
	SessionID string `json:"session_id"`

	// Идентификатор вызываемого навыка, присвоенный при создании.
	// Чтобы узнать идентификатор своего навыка, откройте его в
	// личном кабинете — идентификатор будет в адресе
	// страницы, https://.../developer/skills/<идентификатор>/
	SkillID string `json:"skill_id"`

	// Идентификатор экземпляра приложения, в котором пользователь
	// общается с Алисой, максимум 64 символа.
	// Даже если пользователь авторизован с одним и тем же аккаунтом
	// в приложении Яндекс для Android и iOS,
	// Яндекс.Диалоги присвоят отдельный user_id каждому из этих приложений.
	UserID string `json:"user_id"`
}
