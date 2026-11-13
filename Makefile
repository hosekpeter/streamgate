STREAM ?= match1
SIZE   ?= 1280x720
RTMP   ?= rtmp://localhost:1935

.PHONY: run build test vet fmt up down logs stream

## run: run streamgate locally (needs MediaMTX on localhost:8888)
run:
	go run .

## build: build the binary into ./streamgate
build:
	go build -o streamgate .

## test: run tests with the race detector
test:
	go test -race ./...

## vet: run go vet
vet:
	go vet ./...

## fmt: format the code
fmt:
	gofmt -w .

## up: start streamgate and MediaMTX in Docker
up:
	docker compose up --build -d

## down: stop the Docker stack
down:
	docker compose down

## logs: follow logs of the Docker stack
logs:
	docker compose logs -f

## stream: send a test stream to MediaMTX (Ctrl+C to stop), e.g. make stream STREAM=match2 SIZE=1920x1080
stream:
	ffmpeg -re -f lavfi -i testsrc=size=$(SIZE):rate=30 -f lavfi -i sine \
		-c:v libx264 -preset veryfast -g 60 -c:a aac \
		-f flv $(RTMP)/$(STREAM)
