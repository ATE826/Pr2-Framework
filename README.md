Запуск
docker-compose up --build


Остановка
docker-compose down

Для тестирования:
Переходим в папку
cd .\service-orders\

Запускаем тест:
go test ./internal/services -v

Swagger:
http://localhost:8080/swagger/index.html
