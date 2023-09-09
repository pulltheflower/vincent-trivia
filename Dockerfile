FROM golang:1.19.0

WORKDIR /usr/src/app

RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod tidy