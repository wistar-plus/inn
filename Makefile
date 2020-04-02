.PHONY: *

run-user:
	go run cmd/user/main.go
run-msg:
	go run cmd/message/main.go
run-gw:
	cd cmd/gateway && go run main.go

docker-build:
	cd cmd/user && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v
	cd cmd/message && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v
	cd cmd/gateway && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v
docker-stop:
	docker stop user-service message-service gateway-service message-service2
docker-rmi:
	docker rmi inn_user-service inn_message-service inn_gateway-service inn_message-service2
etcdkey:
	ETCDCTL_API=3 ./etcdctl --endpoints=http://127.0.0.1:2379 get / --prefix --keys-only
docker-run: docker-build
	docker-compose up -d
