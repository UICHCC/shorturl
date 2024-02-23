FROM golang:1.22-alpine as builder

ADD ./ /app

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app
RUN GOARCH=amd64 GOOS=linux go build


FROM alpine

COPY --from=builder /app/ /app
WORKDIR /app

CMD ["./shorturl-go"]
EXPOSE 3000
