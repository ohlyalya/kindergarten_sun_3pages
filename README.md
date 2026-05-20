# ГБДОУ «Солнце» — сайт с Backend на Go, Swagger и отдельным Frontend

Проект сделан под 3 страницы макета:

1. `index.html` — главная страница;
2. `programs.html` — образовательные программы, группы, занятия и коллектив;
3. `parents.html` — отзывы, мероприятия, новости и события.

Frontend и Backend запускаются отдельно. Frontend обращается к Backend по HTTP-запросам.

## Структура

kindergarten_sun_3pages/
├── backend/
│   ├── cmd/server/main.go
│   ├── docs/openapi.yaml
│   └── internal/
│       ├── handlers/
│       ├── models/
│       ├── service/
│       └── storage/
└── frontend/
    ├── index.html
    ├── programs.html
    ├── parents.html
    ├── style.css
    ├── app.js
    └── assets/
```

## Запуск Backend

cd backend
go run ./cmd/server


Backend запустится на адресе:

http://localhost:8080


Swagger UI:

http://localhost:8080/swagger/


OpenAPI YAML:

http://localhost:8080/docs/openapi.yaml


Если порт 8080 занят, можно запустить на другом порту:

PORT=18080 go run ./cmd/server


## Запуск Frontend

Во втором терминале:

cd frontend
python -m http.server 5500

Открыть в браузере:

http://localhost:5500

## Backend-эндпоинты

- `GET /api/health` — проверка работы сервера;
- `GET /api/home` — данные главной страницы;
- `GET /api/programs` — группы, занятия, педагоги;
- `GET /api/parents` — отзывы, мероприятия, новости;
- `GET /api/search?q=...` — поиск по контенту сайта;
- `POST /api/applications` — отправка заявки на запись;
- `GET /api/applications` — просмотр сохранённых заявок;
- `POST /api/questions` — отправка вопроса администрации;
- `GET /api/questions` — просмотр сохранённых вопросов.

Заявки и вопросы сохраняются в памяти приложения через map. Это не мок: backend валидирует входные данные, присваивает ID, статус, дату создания и хранит записи до перезапуска сервера.