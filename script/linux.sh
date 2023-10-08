cd ../

go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct

go env -w GOOS=linux
go mod tidy
go build -o ./dist/server src/main.go