# A `Makefile` is a file used to define commands for building, running, testing, and managing a project using the make command.
# Instead of typing long commands, you run:
# make run
# make build
# make test
# make tidy

run:
	go run cmd/server/main.go

build:
	go build -o bin/server cmd/server/main.go

test:
	go test ./...

tidy:
	go mod tidy

exec-project:
	go mod tidy & go build -o bin/server cmd/server/main.go & \
	go run cmd/server/main.go
