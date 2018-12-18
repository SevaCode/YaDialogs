// main
package main

import (
	"encoding/json"
	"fmt"

	"github.com/SevaCode/YaDialogs/yainterface"
)

func main() {
	fmt.Println("Yandex Dialogs")

	requestExample()
	responseExample()
}

func requestExample() {
	var jsonStr = `
{
  "meta": {
    "locale": "ru-RU",
    "timezone": "Europe/Moscow",
    "client_id": "ru.yandex.searchplugin/5.80 (Samsung Galaxy; Android 4.4)",
    "interfaces": {
      "screen": { }
    }
  },
  "request": {
    "command": "закажи пиццу на улицу льва толстого, 16 на завтра",
    "original_utterance": "закажи пиццу на улицу льва толстого, 16 на завтра",
    "type": "SimpleUtterance",
    "markup": {
      "dangerous_context": true
    },
    "payload": {},
    "nlu": {
      "tokens": [
        "закажи",
        "пиццу",
        "на",
        "льва",
        "толстого",
        "16",
        "на",
        "завтра"
      ],
      "entities": [
        {
          "tokens": {
            "start": 2,
            "end": 6
          },
          "type": "YANDEX.GEO",
          "value": {
            "house_number": "16",
            "street": "льва толстого"
          }
        },
        {
          "tokens": {
            "start": 3,
            "end": 5
          },
          "type": "YANDEX.FIO",
          "value": {
            "first_name": "лев",
            "last_name": "толстой"
          }
        },
        {
          "tokens": {
            "start": 5,
            "end": 6
          },
          "type": "YANDEX.NUMBER",
          "value": 16
        },
        {
          "tokens": {
            "start": 6,
            "end": 8
          },
          "type": "YANDEX.DATETIME",
          "value": {
            "day": 1,
            "day_is_relative": true
          }
        }
      ]
    }
  },
  "session": {
    "new": true,
    "message_id": 4,
    "session_id": "2eac4854-fce721f3-b845abba-20d60",
    "skill_id": "3ad36498-f5rd-4079-a14b-788652932056",
    "user_id": "AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC"
  },
  "version": "1.0"
}`

	var request yainterface.RootRequest

	err := json.Unmarshal([]byte(jsonStr), &request)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Request: %#v\n", request)
	}
}

func responseExample() {
	jsonStrs := []string{`
{
  "response": {
    "text": "Здравствуйте! Это мы, хороводоведы.",
    "tts": "Здравствуйте! Это мы, хоров+одо в+еды.",
    "buttons": [
        {
            "title": "Надпись на кнопке",
            "payload": {},
            "url": "https://example.com/",
            "hide": true
        }
    ],
    "end_session": false
  },
  "session": {
    "session_id": "2eac4854-fce721f3-b845abba-20d60",
    "message_id": 4,
    "user_id": "AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC"
  },
  "version": "1.0"
}`, `
{
  "response": {
    "text": "Здравствуйте! Это мы, хороводоведы.",
    "tts": "Здравствуйте! Это мы, хоров+одо в+еды.",
    "card": {
      "type": "BigImage",
      "image_id": "1027858/46r960da47f60207e924",
      "title": "Заголовок для изображения",
      "description": "Описание изображения.",
      "button": {
        "text": "Надпись на кнопке",
        "url": "http://example.com/",
        "payload": {}
      }
    },
    "buttons": [
      {
        "title": "Надпись на кнопке",
        "payload": {},
        "url": "https://example.com/",
        "hide": true
      }
    ],
    "end_session": false
  },
  "session": {
    "session_id": "2eac4854-fce721f3-b845abba-20d60",
    "message_id": 4,
    "user_id": "AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC"
  },
  "version": "1.0"
}`, `
{
  "response": {
    "text": "Здравствуйте! Это мы, хороводоведы.",
    "tts": "Здравствуйте! Это мы, хоров+одо в+еды.",
    "card": {
      "type": "ItemsList",
	  "header": {
        "text": "Заголовок галереи изображений"
      },
	  "items": [
        {
          "image_id": "<image_id>",
          "title": "Заголовок для изображения.",
          "description": "Описание изображения.",
          "button": {
            "text": "Надпись на кнопке",
            "url": "http://example.com/",
            "payload": {}
          }
        }
      ],
      "footer": {
        "text": "Текст блока под изображением.",
        "button": {
          "text": "Надпись на кнопке",
          "url": "https://example.com/",
          "payload": {}
        }
      }
    },
    "buttons": [
      {
        "title": "Надпись на кнопке",
        "payload": {},
        "url": "https://example.com/",
        "hide": true
      }
    ],
    "end_session": false
  },
  "session": {
    "session_id": "2eac4854-fce721f3-b845abba-20d60",
    "message_id": 4,
    "user_id": "AC9WC3DF6FCE052E45A4566A48E6B7193774B84814CE49A922E163B8B29881DC"
  },
  "version": "1.0"
}`,
	}

	for i, jsonStr := range jsonStrs {
		var response yainterface.RootResponse
		err := json.Unmarshal([]byte(jsonStr), &response)

		if err != nil {
			fmt.Println("Error: ", err)
		} else {
			fmt.Printf("Response[%d]: %#v\n", i, response)
		}
	}
}
