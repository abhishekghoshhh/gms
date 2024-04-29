# GMS setup and testing


## Unit tetsing
```
visit gomock repository by the following link:
https://github.com/uber-go/mock

install the mockgen tool:
go install go.uber.org/mock/mockgen@latest

To ensure it was installed correctly, use:
mockgen -version

If that fails, make sure your GOPATH/bin is in your PATH. You can add it with:
export PATH=$PATH:$(go env GOPATH)/bin

create mocks using the following command
mockgen -source=pkg/iam/client.go -destination=mocks/mock_iam_client.go -package=mocks
mockgen -source=pkg/http/custom_client.go -destination=mocks/mock_http_client.go -package=mocks

```


## Live reloading
```
visit air repository by the following link:
https://github.com/cosmtrek/air

install the air tool with go:
go install github.com/cosmtrek/air@latest

initialize project:
air init

start project with existing toml file
air -c .air.toml
```