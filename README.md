# Todo API для ecom.tech

Простой и эффективный HTTP-сервер на Go, реализующий CRUD-операции для управления задачами (Todo list).  
Данные хранятся в памяти, без внешних баз данных.  
Проект выполнен в рамках тестового задания на стажировку - используется только стандартная библиотека Go.

[![Go](https://img.shields.io/badge/Go-1.25+-00ADD8?logo=go&logoColor=white)](https://go.dev/)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen)](https://github.com/Wastedyouth225/ecom-tech)
[![Tests](https://img.shields.io/badge/tests-passing-brightgreen)](https://github.com/Wastedyouth225/ecom-tech)
[![Docker](https://img.shields.io/badge/docker-ready-2496ED?logo=docker&logoColor=white)](https://github.com/Wastedyouth225/ecom-tech/blob/main/Dockerfile)

## Технологии
- [Go](https://go.dev/) (стандартная библиотека)
- net/http
- encoding/json
- sync (для thread-safe хранилища)
- Docker 

## Использование

Сервер предоставляет REST API для работы с задачами.

### Эндпоинты

| Метод  | Путь             | Описание                           | Тело запроса (JSON)                                                                 |
|--------|------------------|------------------------------------|-------------------------------------------------------------------------------------|
| POST   | `/todos`         | Создать новую задачу               | `{ "title": string*, "description": string, "completed": bool }`                    |
| GET    | `/todos`         | Получить список всех задач         | —                                                                                   |
| GET    | `/todos/{id}`    | Получить задачу по ID              | —                                                                                   |
| PUT    | `/todos/{id}`    | Обновить задачу по ID              | `{ "title": string*, "description": string, "completed": bool }`                    |
| DELETE | `/todos/{id}`    | Удалить задачу по ID               | —                                                                                   |

`*` — поле `title` обязательно и не может быть пустым.

### Структура задачи
```json
{
  "id": 1,
  "title": "string",
  "description": "string",
  "completed": false
}
```
## Примеры использования
### Создать задачу
```sh
curl -v -X POST http://localhost:8080/todos \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить курицу","description":"2 кг"}'
```
## Получить все задачи
```
curl -v http://localhost:8080/todos
```
## Получить задачу по ID
```
curl -v http://localhost:8080/todos/1
```
##  Обновить задачу
```
curl -X PUT http://localhost:8080/todos/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Купить хлеб","completed":true}'
```
##  Удалить задачу
```
curl -v -X DELETE http://localhost:8080/todos/1
```

## Разработка
Требования:
Go 1.25 или выше
## Установка зависимостей
```
go mod tidy
```
## Запуск Development сервера
```
go run ./cmd/server
```
## Запуск в Docker
```
docker build -t ecom-tech-todo .
docker run -p 8080:8080 ecom-tech-todo
```
## Тестирование
Проект покрыт unit-тестами, проверяющими:

- создание задач
- валидацию
- обновление и удаление
- обработку ошибок (400, 404)
- отсутствие дублирования ID

## Запуск тестов
```
go test ./... -v
```
## Структура проекта
```
ecom-tech/
├── README.md  
├── Dockerfile
├── go.mod
├── go.sum
├── cmd/server/main.go
├── internal/http/handler.go
├── internal/http/middleware.go
├── internal/todo/model.go
├── internal/todo/service.go
├── internal/todo/store.go
├── internal/todo/todo_test.go
````

