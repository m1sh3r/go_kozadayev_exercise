# Задание 5: Умный дом

## Реализация
- Потоковые данные (gRPC Server Streaming) через метод `MonitorReadings`.
- Унифицированная авторизация (Middleware) для gRPC Interceptors и HTTP.
- Проверка статического API-ключa `secret-key`.

## Запуск
```bash
go run task5/main.go
```

## Примеры запросов
### Добавление показаний
```bash
curl -X POST http://localhost:8080/v1/readings -H "x-api-key: secret-key" -d '{"device_id": "temp-1", "value": 22.5}'
```
