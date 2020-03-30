FROM alpine:latest

ADD message .

EXPOSE 8888 

ENTRYPOINT ["./message", "run"]

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v