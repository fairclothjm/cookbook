BINARY_NAME=main.out

all: build test

build: ${BINARY_NAME}

$(BINARY_NAME): typeswitch.go
	go build -o ${BINARY_NAME} typeswitch.go

test:
	go test -v typeswitch.go

run: build
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}
