PROJECT := $(shell git config --local remote.origin.url|sed -n 's#.*/\([^.]*\)\.git#\1#p')

.PHONY:
.SILENT:
.DEFAULT_GOAL := run

build:
	go mod download && CGO_ENABLED=0 go build
	docker build --no-cache=true -t ${PROJECT}:latest -f ./Dockerfile .

run: build
	docker-compose up --remove-orphans
	# ${PROJECT} --config data/example/config/storage-server.toml

stop:
	docker-compose down

swag:
	swag init

lint:
	golangci-lint run

test:
	go test -v ./...

gen:
	mockgen -source=internal/domain/model.go -destination=internal/domain/mocks/model.go

clean:
	rm -rf ${PROJECT}
	docker-compose down -v
