FROM golang:1.19.2 as builder
WORKDIR /
ENV GO111MODULE="on"
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux go build -o application
# マルチステージビルド
FROM alpine:3.8
WORKDIR /
# Multi stage build function of Docker
COPY --from=builder application application
COPY --from=builder view/ view/
CMD ["./application"]