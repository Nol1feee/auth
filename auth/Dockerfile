FROM golang:1.20.3-alpine AS builder

COPY . /github.com/Nol1feee/CLI-chat/auth/
WORKDIR /github.com/Nol1feee/CLI-chat/auth/

RUN go mod download
RUN go build -o ./bin/service_linux cmd/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/Nol1feee/CLI-chat/auth/bin/service_linux .

CMD ["./service_linux"]