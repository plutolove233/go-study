etcd-start:
	etcd

go-mod-init:
	go mod tidy

user-rpc-start:
	go run ./user/rpc/user.go -f ./user/rpc/etc/user.yaml

order-api-start:
	go run ./order/api/order.go -f ./order/api/etc/order.yaml

send-request:
	curl -i -X GET http://localhost:8888/api/order/get/1
