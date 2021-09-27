# Пример gRPC клиента и сервера

## [Установка зависимостей](https://grpc.io/docs/languages/go/quickstart/#prerequisites)

### Запуск
1) ```go run cmd/server/main.go```
2) ```go run cmd/client/main.go```

### Генерация .pb.go файлов
```protoc --go_out=plugins=grpc:proto proto/notification.proto```