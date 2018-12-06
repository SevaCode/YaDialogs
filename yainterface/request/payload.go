package request // import "github.com/SevaCode/YaDialogs/yainterface/request"

// JSON, полученный с нажатой кнопкой от обработчика навыка
// (в ответе на предыдущий запрос), максимум 4096 байт.
type Payload interface{}
