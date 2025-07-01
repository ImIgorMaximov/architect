# Social Network - Монолитное приложение

### Требования

- Go 1.21+
- PostgreSQL

### Установка

1. Клонируйте репозиторий.
2. Создайте базу данных:
    ```
    createdb social_network
    psql -U imx -d social_network -f schema.sql
    ```
3. Запустите сервер:
    ```
    go run cmd/main.go internal/handlers.go internal/db.go internal/models.go
    ```
Сервер будет доступен на `http://localhost:8080`.

## API Методы

- `POST /user/register` – регистрация  
- `POST /login` – авторизация  
- `GET /user/get/{id}` – получить анкету

## Тестирование

Импортируйте `postman_collection.json` в Postman.

1. Выполните POST-запрос Social Network API -> Register User (images/RegiserUser.png)
2. Выполните POST-запрос Social Network API -> Login  (images/Login.png)
3. Выполните GET-запрос Social Network API -> Get User by ID (images/GetUserByID.png)

