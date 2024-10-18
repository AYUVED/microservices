# FROM golang:1.22.3 AS builder
# WORKDIR /usr/src/app
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o eventEmitter ./cmd/main.go

# FROM scratch
# COPY --from=builder /usr/src/app/eventEmitter ./eventEmitter
# CMD ["./eventEmitter"]

FROM alpine:latest
RUN mkdir /app
COPY eventEmitter /app
CMD [ "/app/eventEmitter"]