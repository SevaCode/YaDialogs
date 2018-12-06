package meta // import "github.com/SevaCode/YaDialogs/yainterface/meta"

import "github.com/SevaCode/YaDialogs/yainterface/meta/interfaces"

// Интерфейсы, доступные на устройстве пользователя.
type Interfaces struct {

	// Пользователь может видеть ответ навыка на экране
	// и открывать ссылки в браузере.
	Screen interfaces.Screen `json:"screen"`
}
