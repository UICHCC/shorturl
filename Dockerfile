FROM golang:1.22-alpine AS builder

ADD ./ /app

RUN go env -w GOPROXY=https://goproxy.cn,direct

WORKDIR /app
RUN GOARCH=amd64 GOOS=linux go build


FROM alpine

COPY --from=builder /app/ /app
WORKDIR /app
RUN rm -rf controller model service util .gitignore main.go go.* *.md *.sh Dockerfile LICENSE .circleci .git


CMD ["./shorturl-go"]
EXPOSE 3000
