FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o booking_api_testcase ./main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/booking_api_testcase /app/booking_api_testcase
COPY --from=builder /app/configs /app/configs
EXPOSE 3000
CMD ["/app/booking_api_testcase"]