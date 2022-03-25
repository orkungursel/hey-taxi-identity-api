export SERVER_HOST=localhost

run:
	cd ./cmd/ && go run -tags swagger .

build:
	rm -rf ./bin/ &&  mkdir -p ./bin/ && go build -ldflags="-s -w" -trimpath -o bin/main cmd/main.go

swagger: swagger-fmt
	swag init -g ./pkg/server/swagger/doc.go -pd --parseDepth 2

swagger-fmt:
	swag fmt -g ./pkg/server/swagger/doc.go
