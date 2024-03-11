FROM golang:alpine
# RUN apk update && apk add --no-cache git && apk add --no-cache bash && apk add build-base

WORKDIR /usr/app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod tidy

CMD ["air", "./src/cmd/main.go", "-b", "0.0.0.0"]