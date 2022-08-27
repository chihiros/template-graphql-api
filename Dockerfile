FROM golang:1.18-bullseye AS docker
WORKDIR /app
COPY /app .
RUN go install github.com/cosmtrek/air@latest
CMD air
