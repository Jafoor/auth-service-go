MAIN:=./
TARGET:=main
WIN_TARGET:=auth-service.exe
SERVER_CMD:=./${TARGET} serve

build-proto:
	protoc ${PROTOC_FLAGS} ${ORDER_RESELLER_PROTO_FILES}

run-server:
	${SERVER_CMD}

vendor:
	go mod vendor

install-dev-deps:
	go install github.com/air-verse/air@latest

install-deps:
	go mod download


prepare: install-dev-deps install-deps vendor build-proto

dev: prepare
	air serve

build: install-deps
	go build -o ${TARGET} ${MAIN}

start: build run-server

dev-win: prepare
	air -c .air.win.toml


build-win: install-deps
	go build -o ${WIN_TARGET} ${MAIN}


start-win: build-win run-server


# Format using both gofumpt and golines
fmt:
	@echo "Formatting Go files with gofumpt and golines..."
	@gofumpt -w . && golines -w .
	@echo "All Go files formatted!"

# test all
test:
	go test -v ./...

postgres:
	sudo docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres:17-alpine
postgres-start:
	sudo docker start postgres17
createdb:
	sudo docker exec -it postgres17 createdb --username=postgres --owner=postgres resale_orders