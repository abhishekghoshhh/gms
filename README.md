# GMS setup and testing


## Unit tetsing
```
visit gomock repository by the following link:
https://github.com/uber-go/mock

nstall the mockgen tool:
go install go.uber.org/mock/mockgen@latest

To ensure it was installed correctly, use:
mockgen -version

If that fails, make sure your GOPATH/bin is in your PATH. You can add it with:
export PATH=$PATH:$(go env GOPATH)/bin

create mocks using the following command
mockgen -source=pkg/lib/build_capabilities.go -destination=mocks/mock_build_capabilities.go -package=mocks
mockgen -source=pkg/lib/get_groups.go -destination=mocks/mock_get_groups.go -package=mocks
```