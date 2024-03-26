.SILENT:

.PHONY: build test image

IMAGE_NAME = abhishek1009/gms-go
DEFAULT_TAG = latest


say_hello:
	cat ./Makefile

build:
	go build cmd/gms/main.go

run:
	go run cmd/gms/main.go

test:
	go test ./...

image:
	docker build -t $(IMAGE_NAME):$(DEFAULT_TAG) -f deployment/Dockerfile .

air:
	air -c .air.toml