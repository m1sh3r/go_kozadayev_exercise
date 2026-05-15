# Задание 1: Базовый гараж

## Реализация
- gRPC сервер с методами `CreateCar` и `GetCar`.
- REST-обертка через **grpc-gateway** (кодогенерация).
- Хранение данных в памяти (in-memory map) с использованием RWMutex.

## Запуск
```bash
go run task1/main.go
```

## Примеры запросов
### Создание машины
```bash
curl -X POST http://localhost:8080/v1/cars -d '{"vin": "F123", "brand": "Ford", "model": "Focus", "year": 2010}'
```
### Получение машины
```bash
curl http://localhost:8080/v1/cars/F123
```
