# URL Shortener

Cервис для сокращения длинных ссылок, написанный на **Go** с использованием **Gin**, **PostgreSQL**, **Docker** и **Swagger**.

## 🚀 Возможности
- Генерация коротких ссылок
- Перенаправление по сокращённому URL
- REST API с документацией Swagger
- Поддержка PostgreSQL для хранения данных
- Контейнеризация с Docker

## 📦 Технологии
- **Go** (Gin, sqlx, squirrel)
- **PostgreSQL** (хранение ссылок)
- **Docker & Docker Compose** (упрощённый запуск)
- **Swagger** (автодокументация API)

---

## ⚡ Быстрый старт
### 1️⃣ Клонируем репозиторий
```sh
git clone https://github.com/yourusername/url-shortener.git
cd url-shortener
```

### 2️⃣ Запускаем с Docker
```sh
docker-compose up -d --build
```

### 3️⃣ Проверяем работу
Открываем **Swagger UI**:
```
http://localhost:8080/swagger/index.html
```

Пример запроса:
```sh
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://google.com"}'
```
Пример ответа:
```json
{
  "short_url": "http://localhost:8080/abc123",
  "id": 1
}
```

---

## ⚙️ API Документация
Swagger доступен по адресу:
```
http://localhost:8080/swagger/index.html
```
### ✨ Основные эндпоинты:
- **`POST /shorten`** — создание короткой ссылки
- **`GET /:short`** — редирект по короткому URL

---

## 🔧 Локальный запуск url-shortener (go run)
### 1️⃣ Запускаем БД с Docker
```sh
docker-compose up -d postgres pgadmin
```

### 2️⃣ Запускаем сервер
```sh
make run
```

### 3️⃣ Тестируем API
```sh
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://google.com"}'
```

---
## 📌 TODO
- [ ] Добавить поддержку пользователей (авторизация)
- [ ] Аналитика переходов по ссылкам