@echo off

cd ../
go env -w CGO_ENABLED=1

@rem Windows
go env -w GOOS=windows
go mod tidy
go build -o ./dist/server.exe src/main.go

@rem Linux
@rem go env -w CGO_ENABLED=0
@rem go env -w GOOS=linux
@rem go mod tidy
@rem go build -o ./dist/server src/main.go

@rem docker pull golang:1.21.0-bullseye
@rem docker run -it --rm -v /root/server:/app golang:1.21.0-bullseye sh -c 'cd /app && cd ./script && ./linux.sh'