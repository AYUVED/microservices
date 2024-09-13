# FROM golang:1.22.3 AS builder
# WORKDIR /usr/src/app
# COPY . .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o log ./cmd/main.go

# FROM scratch
# COPY --from=builder /usr/src/app/log ./log
# CMD ["./log"]

FROM alpine:latest
RUN mkdir /app
COPY log /app
CMD [ "/app/log"]