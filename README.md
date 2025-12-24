# Ecom Tech TODO API

## Запуск локально
go run ./cmd/server

## Запуск через Docker
docker build -t ecom-tech .
docker run -p 8080:8080 ecom-tech

## Эндпоинты
- POST /todos — создать задачу
- GET /todos — список задач
- GET /todos/{id} — получить задачу по ID
- PUT /todos/{id} — обновить задачу
- DELETE /todos/{id} — удалить задачу
