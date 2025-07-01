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
    go run main.go handlers.go db.go models.go
    ```
Сервер будет доступен на `http://localhost:8080`.

## API Методы

- `POST /user/register` – регистрация  
- `POST /login` – авторизация  
- `GET /user/get/{id}` – получить анкету

## Тестирование

Импортируйте `postman_collection.json` в Postman.

1. Выполните POST-запрос Social Network API -> Register User :

2. Выполните POST-запрос Social Network API -> Login :

3. Выполните GET-запрос Social Network API -> Get User by ID :

