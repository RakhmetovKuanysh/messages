## Микросервис для отправки и получения сообщений

Микросервис позволяет получать и отправлять сообщения для проекта social-network Otus.

- - -

### Техническое описание

|||
|-----------------------|---------------------|
| Язык программирования | `Golang`            |
| База данных           | `MySQL`             |
| Протокол обмена       | `HTTP`              |
| Формат обмена         | `JSON`              |

- - -

### Методы

#### 1. Healthcheck [GET]

- URL:    `/health`
- Method: `GET`

#### 2. Получить сообщения [GET]

+ URL:        `/messages`
+ Method:     `GET`
+ Body:       `application/json`
+ Params:

|||
|--------------|--------|
|`receiverId`  | 1      |
|`senderId`    | 2      |

**Успешный ответ**
```json
{
    "message": "Found",
    "code": 200,
    "messages": [
        {
            "id": "3",
            "text": "Hello",
            "receiver_id": 1,
            "sender_id": 2,
            "created_at": "2020-04-12T11:15:41Z"
        }
    ]
}
```

#### 3. Добавить сообщений в базу [POST]

+ URL:        `/message`
+ Method:     `POST`
+ Body:       `application/json`
+ Params:

|||
|--------------|--------|
|`receiverId`  | 1      |
|`senderId`    | 2      |
|`text`        | Hi!    |

**Успешный ответ**
```json
{
    "message": "Added",
    "code": 200
}
```
