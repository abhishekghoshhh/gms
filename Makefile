.SILENT:

.PHONY: build test image

IMAGE_NAME = abhishek1009/gms-go
DEFAULT_TAG = latest

RUN_ARGS = IAM_HOST=https://ska-iam.stfc.ac.uk PASSWORD_GRANT_FLOW_USERNAME=fake-name PASSWORD_GRANT_FLOW_PASSWORD=fake-password PASSWORD_GRANT_FLOW_CLIENT_ID=fake-client-id PASSWORD_GRANT_FLOW_CLIENT_SECRET=fake-client-secret

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