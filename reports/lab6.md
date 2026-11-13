# Лабораторная работа №6 — вариант 5

## Цель
Реализовать сервис с использованием Go kit и двух транспортов: HTTP и gRPC.

## Предметная область
Ресторан: `ID`, `Name`, `Address`, `Rating float32`, `Open bool`.

Методы сервиса:
- `Create(ctx, restaurant) (string, error)`
- `GetByID(ctx, id) (Restaurant, error)`
- `ChangeRating(ctx, id, rating) error`

## Архитектура
- `internal/restaurant`: модель, интерфейс репозитория, потокобезопасное in-memory хранилище и бизнес-сервис.
- `internal/transport`: Go kit endpoints, HTTP-декодирование и gRPC transport.
- `cmd/server`: одновременный запуск HTTP (:8080) и gRPC (:9090).
- `api/restaurant/v1/restaurant.proto`: описание gRPC API.

## Запуск и проверка
```bash
go mod tidy
go test ./...
go run ./cmd/server
```

HTTP:
```bash
Invoke-RestMethod `
  -Uri "http://localhost:8080/restaurants" `
  -Method POST `
  -ContentType "application/json" `
  -Body $body
curl.exe "http://localhost:8080/restaurants/r1"
curl.exe -X PATCH "http://localhost:8080/restaurants/r1" `
  -H "Content-Type: application/json" `
  --data-raw '{\"rating\":4.8}'
```

gRPC (при установленном grpcurl):
```bash
grpcurl -plaintext localhost:9090 list
Get-Content -Raw .\request.json | grpcurl -plaintext `
  -import-path . `
  -proto api/restaurant/v1/restaurant.proto `
  -d "@" `
  localhost:9090 restaurant.v1.RestaurantService/Create
Get-Content -Raw .\get.json | grpcurl -plaintext `
  -import-path . `
  -proto api/restaurant/v1/restaurant.proto `
  -d '@' `
  localhost:9090 restaurant.v1.RestaurantService/GetByID
Get-Content -Raw .\change.json | grpcurl -plaintext `
  -import-path . `
  -proto api/restaurant/v1/restaurant.proto `
  -d '@' `
  localhost:9090 restaurant.v1.RestaurantService/ChangeRating
```

## Тестирование
Добавлены unit-тесты жизненного цикла ресторана и проверки валидации/ошибки отсутствующей записи. Автоматический запуск тестов будет выполнен CI после push/создания PR. Результаты реального запуска в этом отчёте не заявляются заранее.
