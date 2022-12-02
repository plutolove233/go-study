search-api-gen:
	goctl api go -api search.api -dir .

search-service-start:
	go run search.go -f etc/search-api.yaml

search-service-test:
	curl -i -X GET http://127.0.0.1:8889/search/do?name=%E8%A5%BF%E6%B8%B8%E8%AE%B0 -H "Authorization:eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2Njk5NTM5NTgsImlhdCI6MTY2OTk1MDM1OCwidXNlcklkIjoxfQ.Sje6PCk9-Ck9izn48kqq3YUNFJxoW2fbSvg0hpEAj6w"