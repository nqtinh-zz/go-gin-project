init:
	go mod init
	go mod tidy

dev:
	go run cmd/server/main.go

generate-test:
	go generate ./...

test:
	go test ./...

setup-db:
	docker-compose -f "docker-compose.yaml" up -d --build

clean:
	docker-compose -f "docker-compose.yaml" down

migration-db:
	export DATABASE_URL="postgres://postgres:postgres@localhost:5434/api?sslmode=disable"
	dbmate -d ./db/init up

serve-swagger:
	which swagger || (GO111MODULE=off go get -u github.com/go-swagger/go-swagger/cmd/swagger)
	GO111MODULE=off swagger generate spec -o ./swagger.yaml --scan-models
	swagger serve -F=swagger swagger.yaml