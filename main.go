//Тест работы описания приммеров создания grpc и RESTAPI
//https://blog.maddevs.io/go-rest-или-grpc-f5d52d7ffff6
//Структура проекта :
//pb пакет для хранения протофайлов и сгенерированного кода
// main.go сервис, который будет конфигурировать и запускать наше приложение
//server пакет для реализации логики нашего приложения

//Описание Protocol Buffers
//https://developers.google.com/protocol-buffers/docs/proto3
//
//Makefile мейкфайл для автоматизации рутины

package main

//import "github.com/maddevsio/grpc-rest-api-example/server"
import "./server"

func main() {
	g := server.New()
	g.Start()
        //без ожидания завершения горутин сервер проваливается
	g.WaitStop()
}