
ecom-tech - Todo API на Go

HTTP-сервер на Go для работы с Todo задачами. Все данные хранятся в памяти (без внешней базы данных).

Эндпоинты
- POST /todos - создать новую задачу
- GET /todos - получить список всех задач
- GET /todos/{id} - получить задачу по ID
- PUT /todos/{id} - обновить задачу по ID
- DELETE /todos/{id} - удалить задачу по ID

Структура задачи
- id - числовой, автоматически присваивается
- title - строка, обязательное поле
- description - строка
- completed - bool, признак завершённости

Валидация
- title не может быть пустым при создании или обновлении
- Ошибка валидации → 400 Bad Request
- Если задача с указанным ID не найдена → 404 Not Found

Технологии
- Go 1.25.5
- Стандартная библиотека
- Docker

Запуск

Локально:
git clone https://github.com/Wastedyouth225/ecom-tech.git
cd ecom-tech
go run ./cmd/server
Сервер запускается на порту 8080.

Через Docker:
docker build -t ecom-tech .
docker run -p 8080:8080 ecom-tech

Unit-тесты
Расположены в internal/todo/todo_test.go
Запуск:
go test ./... -v

Примеры запросов

Создание задачи:
curl -X POST http://localhost:8080/todos -H "Content-Type: application/json" -d '{"title":"Test","description":"Demo"}'

Получение всех задач:
curl http://localhost:8080/todos

Получение задачи по ID:
curl http://localhost:8080/todos/1

Обновление задачи:
curl -X PUT http://localhost:8080/todos/1 -H "Content-Type: application/json" -d '{"title":"Updated","description":"Updated","completed":true}'

Удаление задачи:
curl -X DELETE http://localhost:8080/todos/1

Логирование
Сервер выводит в консоль:
Server started at :8080

Структура проекта:

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
├── pkg/
