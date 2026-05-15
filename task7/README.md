# Задание 7: Поисковый движок

## Реализация
- Межсервисное взаимодействие: Content Service дергает Search Service.
- Паттерн **Circuit Breaker** (разрыв цепи) с использованием библиотеки `gobreaker`.
- Статьи создаются даже при недоступности сервиса поиска (fallback).

## Запуск
```bash
go run task7/main.go
```

## Примеры запросов
### Создание статьи
```bash
curl -X POST http://localhost:8080/v1/articles -d '{"title": "Go Patterns", "content": "Circuit breaker...", "author_id": "user-1"}'
```
### Получение статьи
```bash
curl http://localhost:8080/v1/articles/1
```
