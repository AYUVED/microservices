# FROM golang:1.22.3 AS builder
# WORKDIR /usr/src/app
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o logservice ./cmd/main.go

# FROM scratch
# COPY --from=builder /usr/src/app/logservice ./logservice
# CMD ["./logservice"]

FROM alpine:latest
RUN mkdir /app
COPY logservice /app
CMD [ "/app/logservice"]