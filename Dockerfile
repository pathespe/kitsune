FROM golang:1.16.5 AS builder
ENV GOOS linux
ENV CGO_ENABLED 0
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN cd cmd && go build -o .

FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder app .
EXPOSE 4000
CMD ./cmd/cmd