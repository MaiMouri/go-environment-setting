FROM golang:1.14.4-alpine

ENV GO111MODULE=on


WORKDIR /go/src/github.com/s-moteki/go-playground/environment-setting-file/app

COPY go.mod go.mod
COPY go.sum go.sum

EXPOSE 8080
CMD ["sh", "./start_app.sh"]
