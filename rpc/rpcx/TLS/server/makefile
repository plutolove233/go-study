genRSA:
	openssl genrsa -out server.key 2048

genCSR:
	openssl req -new -key server.key -out server.csr

genCRT:
	 openssl x509 -req -days 3650 -in server.csr -signkey server.key -out server.pem
