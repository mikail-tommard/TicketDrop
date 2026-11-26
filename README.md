# TicketDrop

TicketDrop — учебный сервер на Go для продажи билетов на концерт с ограниченным количеством мест. По замыслу покупатель сможет забронировать место на две минуты и получить билет после подтверждения. Просроченная бронь освобождает место, а при одновременных заявках последнее место достаётся одному покупателю.

## Что уже работает

HTTP-сервер на стандартном `net/http` слушает порт `8080`. При запуске он создаёт в памяти один демоконцерт с ID `1` и остатком в 10 мест. Бронирование и выдача билетов ещё в разработке.

| Маршрут | Ответ |
| --- | --- |
| `GET /livez` | `200 OK` с JSON-строкой `"Server is live"` |
| `GET /concerts/{id}` | `200 OK` с карточкой концерта, `404 Not Found` для неизвестного ID, `400 Bad Request` для нечислового ID |

## Локальный запуск

### PostgreSQL и миграции

Для базы понадобится Docker с Compose. Создай `.env` в корне проекта рядом с `docker-compose.yaml` и подставь свой локальный пароль.

```dotenv
POSTGRES_USER=ticketdrop
POSTGRES_PASSWORD=change_me
POSTGRES_DB=ticketdrop
```

Из корня проекта запусти PostgreSQL и миграции.

```sh
docker compose up -d
docker compose ps -a
docker compose logs migrate
```

Сервис `migrate` ждёт готовности PostgreSQL и запускает `goose up` для SQL-файлов из `migration/`. Дождись статуса `Exited (0)` у `migrate`. Он означает, что миграции завершились успешно. При повторном запуске Goose пропускает уже применённые версии.

Миграции создают таблицы `concerts` и `reservations` и добавляют демоконцерт с ID `1` и 10 доступными местами. С хоста PostgreSQL доступен по адресу `localhost:5433`, а из сети Compose по адресу `postgres:5432`.

PostgreSQL создаёт базу из `POSTGRES_DB` при первой инициализации пустого тома. Если том уже использовался и Goose сообщает `database "ticketdrop" does not exist`, создай базу в текущем томе и повтори запуск.

```sh
docker compose exec postgres sh -c 'createdb -U "$POSTGRES_USER" ticketdrop'
docker compose up -d
docker compose logs migrate
```

### HTTP API

Понадобится Go 1.26.5 или новее. Текущая версия API читает концерты из памяти и может запускаться отдельно от Compose. Из корня проекта запусти сервер.

```sh
go run ./cmd/api
```

В другом терминале проверь API:

```sh
curl -i http://localhost:8080/livez
curl -i http://localhost:8080/concerts/1
curl -i http://localhost:8080/concerts/999
```

Карточка концерта с ID `1`:

```json
{
  "title": "Вечер джаза",
  "description": "Концерт в пятницу",
  "available": 10
}
```

## Как устроен проект

- `cmd/api/` — точка входа и настройки HTTP-сервера.
- `internal/adapters/httpapi/` — маршруты, обработчики и запись JSON-ответов.
- `internal/adapters/memory/` — хранилище концертов в памяти.
- `internal/adapters/repository/` — PostgreSQL-репозиторий для чтения концерта.
- `internal/domain/` — модель концерта и проверка её данных.
- `internal/usecase/` — сценарий получения карточки концерта через интерфейс хранилища.
- `migration/` — SQL-миграции Goose для концертов и броней.

В планируемой архитектуре PostgreSQL хранит остаток мест, брони и билеты. Redis кэширует карточку концерта. Подтверждение брони сохраняет событие в PostgreSQL, а фоновый процесс доставляет его через Kafka для выдачи билета. Prometheus и Grafana показывают нагрузку на API и задержку выдачи.
