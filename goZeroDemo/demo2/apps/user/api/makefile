user-api-gen:
	goctl api go -api user.api -dir .

user-service-start:
	go run user.go -f etc/user-api.yaml

user-service-test:
	curl -i -X POST http://127.0.0.1:8888/user/login -H "Content-Type: application/json" -d "{ \"userName\":\"666\", \"password\":\"123456\" }"

error-test:
	curl -i -X POST http://127.0.0.1:8888/user/login -H "Content-Type: application/json" -d "{ \"userName\":\"1\", \"password\":\"123456\" }"