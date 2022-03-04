FROM golang:1.16-alpine3.13 as builder

WORKDIR .
COPY ./ ./
ENV GOPATH=

RUN go build -mod=vendor -ldflags "-s -w" -o main main.go

ENV CGO_ENABLES=0 \
	GOOS=linux \
	GOARCH=amd64

ENTRYPOINT ["./main"]
