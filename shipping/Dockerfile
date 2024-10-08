# FROM golang:1.22.3 AS builder
# WORKDIR /usr/src/app
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o shipping ./cmd/main.go

# FROM scratch
# COPY --from=builder /usr/src/app/shipping ./shipping
# CMD ["./shipping"]

FROM alpine:latest
RUN mkdir /app
COPY shipping /app
CMD [ "/app/shipping"]