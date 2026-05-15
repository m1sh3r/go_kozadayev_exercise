# Задание 10: Платформа техобслуживания

## Реализация
- Глобальный Changelog через gRPC Unary Interceptor.
- Экспорт данных в формат Excel с использованием библиотеки `excelize/v2`.
- Полный цикл REST + gRPC.

## Запуск
```bash
go run task10/main.go
```

## Примеры запросов
### Создание заказ-наряда
```bash
curl -X POST http://localhost:8080/v1/workorders -d '{"car_id": "F123", "description": "Oil change", "total_cost": 150.0}'
```
### Экспорт в Excel
```bash
curl http://localhost:8080/v1/export
```
