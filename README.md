# AvitoChatAssignment

### Общая Информация
- Основные требования находятся по ссылке: https://github.com/avito-tech/backend-trainee-assignment
- Используется: Docker, Make, Golang, PostgreSQL
- Внешние библиотеки: ZapLogger, GorrilaMux - в качестве мультиплексора
- Сервер работает на 9000 порту
- путь до конфига /internal/app/config.go
- Выполнил: Андронов Дмитрий

### How to Run | Build

**Run Mode**
- Для запуска используется технология Docker
- Склонировать репозиторий и перейти в корень
- sudo docker build -t dev/avito-chat .
- sudo docker run -p 9000:9000 --name dev -t dev/avito-chat


