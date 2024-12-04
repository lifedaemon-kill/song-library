Реализация онлайн библиотеки песен 🎶

Необходимо реализовать следующее

### 1. Выставить rest методы
  1. Получение данных библиотеки с фильтрацией по всем полям и пагинацией
  2. Получение текста песни с пагинацией по куплетам
  3. Удаление песни
  4. Изменение данных песни
  5. Добавление новой песни в формате ```JSON```

```JSON
{
"group": "Muse",
"song": "Supermassive Black Hole"
}
```

### 2. При добавлении сделать запрос в АПИ, описанного сваггером. Апи, описанный сваггером, будет поднят при проверке тестового задания. Реализовывать его отдельно не нужно

```openapi
openapi: 3.0.3
info:
  title: Music info
  version: 0.0.1
paths:
  /info:
    get:
      parameters:
        - name: group
          in: query
          required: true
          schema:
            type: string
        - name: song
          in: query
          required: true
          schema:
            type: string
        responses:
            '200':
              description: Ok
              content:
                application/json:
                  schema:
                    $ref: '#/components/schemas/SongDetail'
            '400':
              description: Bad request
            '500':
              description: Internal server error
components:
  schemas:
    SongDetail:
      required:
        - releaseDate
        - text
        - link
      type: object
      properties:
        releaseDate:
          type: string
          example: 16.07.2006
        text:
          type: string
          example: Ooh baby, don't you know I suffer?\nOoh baby, can you hear me moan?\nYou caught me under false pretenses\nHow long before you let me go?\n\nOoh\nYou set my soul alight\nOoh\nYou set my soul alight
        link:
          type: string
          example: https://www.youtube.com/watch?v=Xsp3_a-PMTw
```

### 3. Обогащенную информацию положить в БД postgres (структура БД должна быть создана путем миграций при старте сервиса)
### 4. Покрыть код debug- и info-логами
### 5. Вынести конфигурационные данные в .env-файл
### 6. Сгенерировать сваггер на реализованное АПИ
