FROM golang:1.15-alpine as builder
WORKDIR /go/src/svc1
COPY ./cmd/app .
RUN ls -a
RUN go build -o webserver .
FROM alpine
WORKDIR /app
COPY --from=builder /go/src/svc1/ /app/
EXPOSE 5001
CMD ["./webserver"]