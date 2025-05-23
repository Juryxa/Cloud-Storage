# Запуск проекта через Docker Compose

В этом проекте используются два микросервиса:
- **auth_service** (аутентификация)
- **file_upload_service** (загрузка файлов)

а также вспомогательные сервисы:
- PostgreSQL
- MinIO
- Redis

Перед запуском необходимо создать ключи и настроить файлы окружения.

---

## 1. Генерация ключей

Если у вас нет ключей, выполните в терминале следующие команды:

```bash
# Сгенерировать приватный ключ
openssl genpkey -algorithm RSA -out private_key.pem -pkeyopt rsa_keygen_bits:2048

# Сгенерировать публичный ключ
openssl rsa -pubout -in private_key.pem -out public_key.pem
```

Поместите файлы `private_key.pem` и `public_key.pem` в папку auth_service. Далее !СКОПИРУЙТЕ! `public_key.pem` и поместите его в file_upload_service.

---

## 2. Настройка файлов окружения

Создайте файлы с переменными окружения:

### Для auth_service нужно создать файл .env такого вида
```ini
POSTGRES_USER=postgres
POSTGRES_PASSWORD=dima15042004
POSTGRES_DB=auth-service
STORAGE_PATH=postgres://postgres:dima15042004@auth_db:5432/auth-service?sslmode=disable
HTTP_SERVER_ADDRESS=0.0.0.0:8080
QUOTA_SERVICE_URL=http://file-upload-service:8081
ACCESS_TOKEN_TTL=15m
REFRESH_TOKEN_TTL=720h
PUBLIC_KEY_PATH=public_key.pem
PRIVATE_KEY_PATH=private_key.pem
```

### Для file_upload_service нужно создать файл .env такого вида
```ini
HTTP_SERVER_ADDRESS=0.0.0.0:8081
JWT_PUBLIC_KEY_PATH=public_key.pem
MINIO_PORT=localhost:9000
MINIO_ROOT_USER=minioadmin
MINIO_ROOT_PASSWORD=minioadmin
MINIO_USE_SSL=false
MINIO_URL_LIFETIME=8h
REDIS_URL_LIFETIME=8h
REDIS_PORT=redis:6379
POSTGRES_USER=postgres
POSTGRES_PASSWORD=dima15042004
POSTGRES_DB=file_upload_service
STORAGE_PATH=postgres://postgres:dima15042004@file_upload_db:5432/file_upload_service?sslmode=disable
```

---

## 3. Запуск Docker Compose

Запустите в терминале команду:
```bash
docker-compose up
```

Сервисы будут запущены со следующими адресами:
- **auth_service:** http://localhost:8080
- **file_upload_service:** http://localhost:8081
- **PostgreSQL:** порт 5433-auth_db и 5434-file_upload_db(на хосте)
- **MinIO:** API – http://localhost:9000, консоль – http://localhost:9001
- **Redis:** порт 6379

---

## Документация, чтобы создать ``` swag init -g cmd/server/main.go -o docs ``` документация к каждому сервису находится в папке docs

### ЕСЛИ НУЖНО ЗАПУСТИТЬ ПРИЛОЖЕНИЕ ЛОКАЛЬНО(БЕЗ ДОКЕРА)ИЩИТЕ ОПИСАНИЕ В ПАПКАХ С ТАКИМ НАЗВАНИЕМ В КОНЦЕ ..._service
