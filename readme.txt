По материалам блога Andrew Minkin
https://blog.maddevs.io/go-rest-%D0%B8%D0%BB%D0%B8-grpc-f5d52d7ffff6
Проверка на практике описанного решения

Пакеты требуемые для установки
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc



Тест для проверки отклика сервера на REST запрос
curl -XGET http://localhost:8090/api/hello/Mike

Результат
HTTP/1.1 200 OK
Content-Length: 119
Content-Type: application/json
Date: Tue, 16 Apr 2019 07:17:00 GMT
Grpc-Metadata-Content-Type: application/grpc

{
    "message": "Имя, Mike, целое число 5, число с плавающей запятой %!d(float64=0.628)"
}

Тест для проверки отклика сервера на REST запрос
curl -XGET http://localhost:8090/api/tim/123
HTTP/1.1 200 OK
Content-Length: 45
Content-Type: application/json
Date: Tue, 16 Apr 2019 07:15:39 GMT
Grpc-Metadata-Content-Type: application/grpc

{
    "message": "Вы ввели число  123"
}
