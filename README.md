# Booking-api-service

#### [Ссылка на задачу](https://github.com/CyberZoneDev/test-tasks/blob/master/backend-golang.md)

### Описание

API на Golang для управлением пользователями и бронированиями. Реализованы следующие функции:

- CRUD для пользователей (получение списка пользователей, получение пользователя по id, обновление пользователя по id, удаление пользователя по id)
- CRUD для бронирований (получение списка бронирований, получение бронирования по id, обновление бронирования по id, удаление бронирования по id)
- Шифрование паролей
- Запуск через Docker Compose

### Стек
- Golang
    - Gorilla/Mux (маршрутизация)
    - Viper (конфигурация)
- Swagger (документация к API)
- PostgreSQL (база данных)
- Docker

### Установка
Клонируйте репозиторий:

```bash
git clone https://github.com/theinlaoq/booking_api_testcase.git
cd booking_api_testcase
```
У вас должны быть установлены: Make, Docker.

### Запуск проекта
Обычный запуск проекта:

```bash
make up
```

API станет доступен по адресу http://localhost:3000

Документация к API доступна по адресу http://localhost:3000/swagger/index.html

Если возникают ошибки, попробуйте пересобрать проект.

Остановка проекта:

```bash
make down
```

#### Контакты автора
tg: @theinlaoq