run:
	go run .
protoc:
	protoc --proto_path=./proto --go_out=./proto --go_opt=paths=source_relative --go-grpc_out=./proto --go-grpc_opt=paths=source_relative ./proto/*.proto

certificate: self
	openssl req -new -x509 -key server.key -out server.crt --addext "subjectAltName = DNS:localhost" -days 365 -subj "/CN=localhost"
self:
	openssl genpkey -algorithm RSA -out server.key