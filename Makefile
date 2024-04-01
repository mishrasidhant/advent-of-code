ADVENT_YEAR?=2023
ADVENT_PUZZLE?=2
BINARY_NAME?=adventPuzzle.out

all: build test

build:
	go build -o ${BINARY_NAME} ./${ADVENT_YEAR}/${ADVENT_PUZZLE}/main.go
run:
	go build -o ${BINARY_NAME} ./${ADVENT_YEAR}/${ADVENT_PUZZLE}/main.go
	./${BINARY_NAME}
test:
	go test -v ./${ADVENT_YEAR}/${ADVENT_PUZZLE}/main.go
clean:
	go clean
	rm ${BINARY_NAME}
	
# Example on how to automate fetching build dependencies
# deps:
#    go get github.com/gorilla/websocket