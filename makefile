EXEC := task5
SRC := cmd/app/main.go

GIN := github.com/gin-gonic/gin
MOCK := github.com/stretchr/testify/mock

all: build run

build:
	go build -o $(EXEC) $(SRC)

run:
	./$(EXEC)

mod:
	go mod init $(EXEC)

get:
	go get \
		$(GIN) \
		$(MOCK)
