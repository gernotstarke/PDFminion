# borrowed from https://tutorialedge.net/golang/makefiles-for-go-developers/

hello:
	echo "Hello"

build:
	go build -o bin/ pdfminion.go

run:
	go run pdfminion.go

compile:
	echo "Compiling for current platform only"


compile-all:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOARCH=arm go build -o bin/main-linux-arm main.go
	GOOS=linux GOARCH=arm64 go build -o bin/main-linux-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go

test:
	echo "Testing with go test"
	go test ./...

bdd:
	echo "Executing Cucumber BDD specifications"
	godog
	node ./cucumber-report-index.js

all: hello build