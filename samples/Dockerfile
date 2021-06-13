FROM golang:1.16.5-alpine as builder

WORKDIR /go/src/github.com/songrgg/testservice/
COPY main.go .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/github.com/songrgg/testservice/app .
CMD ["./app"]