# Задание 2: Сервис-агрегатор документов

## Реализация
- Сущности: Пользователь (User) и Документ (Document) со связью One-to-Many.
- gRPC: `CreateUser`, `AddDocumentToUser`.
- REST: `GET /users/{id}/documents`.
- Хранение: PostgreSQL с использованием `sqlc`.
- Миграции: Находятся в `db/migrations`.

## Запуск
```bash
go run task2/main.go
```

## Примеры запросов
### Получение документов пользователя
```bash
curl http://localhost:8080/users/1/documents
```
