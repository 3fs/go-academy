# Readme

```
go test -cover -v -race $(go list ./...)

go test -cover -v -race -coverprofile=coverage.out github.com/3fs/go-academy/03-log-vendor/04-tests/pkg/greeter

go tool cover -func=coverage.out

go tool cover -html=coverage.out

```
