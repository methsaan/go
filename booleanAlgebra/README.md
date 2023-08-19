## Build

```
go mod init 
go mod tidy

go fmt main.go
golangci-lint run . 

env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -trimpath -ldflags "-s -w"
```
