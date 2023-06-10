FROM golang:1.20.5-alpine AS builder

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o ./app cmd/service/main.go

FROM alpine
RUN apk update
RUN apk upgrade
RUN apk add --no-cache ffmpeg
COPY --from=builder /src/app /usr/bin/app

ARG VIDEO_FILE_PATH
ENV VIDEO_FILE_PATH=$VIDEO_FILE_PATH

CMD app $VIDEO_FILE_PATH
