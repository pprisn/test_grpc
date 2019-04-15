go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc

Тест для проверки отклика сервера на REST запрос
curl -XGET http://localhost:8090/api/hello/Mike

Результат
{"message":"Имя, wer, целое число 5, число с плавающей запятой %!d(float64=0.628)"}