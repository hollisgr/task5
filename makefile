EXEC := task5
SRC := cmd/app/main.go

GIN := github.com/gin-gonic/gin
MOCK := github.com/stretchr/testify/mock

all: build run

build: clean
	go build -o $(EXEC) $(SRC)

run:
	./$(EXEC)

clean:
	rm -f $(EXEC)

mod:
	go mod init $(EXEC)

get:
	go get \
		$(GIN) \
		$(MOCK)
