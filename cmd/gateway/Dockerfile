FROM alpine:latest

ADD / .

EXPOSE 8888 

ENTRYPOINT ["./gateway", "run"]

# CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v