.SILENT:

.PHONY: build test image

IMAGE_NAME = gms-go
DEFAULT_TAG = latest

RUN_ARGS = IAM_HOST=https://ska-iam.stfc.ac.uk

say_hello:
	cat ./Makefile

build:
	go build cmd/gms/main.go

run:
	$(RUN_ARGS) go run cmd/gms/main.go

test:
	go test ./...

image:
	docker build -t $(IMAGE_NAME):$(DEFAULT_TAG) -f deployment/Dockerfile .

air:
	air -c .air.toml