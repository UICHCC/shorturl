FROM golang:1.22-alpine as builder

ADD ./ /app

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app
RUN GOARCH=amd64 GOOS=linux go build


FROM alpine

WORKDIR /app
COPY --from=builder /app/shorturl-go /app

CMD ["./shorturl-go"]
EXPOSE 3000
