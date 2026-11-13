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
curl -X POST localhost:8080/restaurants -H 'Content-Type: application/json' -d '{"id":"r1","name":"Cafe","address":"Main 1","rating":4.5,"open":true}'
curl localhost:8080/restaurants/r1
curl -X PATCH localhost:8080/restaurants/r1 -H 'Content-Type: application/json' -d '{"rating":4.8}'
```

gRPC (при установленном grpcurl):
```bash
grpcurl -plaintext localhost:9090 list
grpcurl -plaintext -d '{"method":"Create","id":"r2","name":"Bistro","address":"Main 2","rating":4.1,"open":true}' localhost:9090 restaurant.v1.RestaurantService/Create
grpcurl -plaintext -d '{"method":"GetByID","id":"r2"}' localhost:9090 restaurant.v1.RestaurantService/GetByID
grpcurl -plaintext -d '{"method":"ChangeRating","id":"r2","rating":4.9}' localhost:9090 restaurant.v1.RestaurantService/ChangeRating
```

## Тестирование
Добавлены unit-тесты жизненного цикла ресторана и проверки валидации/ошибки отсутствующей записи. Автоматический запуск тестов будет выполнен CI после push/создания PR. Результаты реального запуска в этом отчёте не заявляются заранее.
