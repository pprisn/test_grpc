go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u google.golang.org/grpc

���� ��� �������� ������� ������� �� REST ������
curl -XGET http://localhost:8090/api/hello/Mike

���������
{"message":"���, wer, ����� ����� 5, ����� � ��������� ������� %!d(float64=0.628)"}