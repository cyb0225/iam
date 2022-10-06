## create a executable file.
FROM golang:1.18.4-alpine as build

RUN mkdir /app

WORKDIR /app

ENV GOPROXY=https://goproxy.cn,direct

COPY . .

RUN go build -o main ./cmd/apiserver/main.go

FROM alpine:latest as build_1

RUN apk --no-cache add ca-certificates

COPY --from=build /app/main ./
ENV PORT=12345 RUN_MODE=release

EXPOSE 12345

VOLUME ["/storage", "/config"]

CMD ["./main", "--port=12345", "--mode=$RUN_MODE"]