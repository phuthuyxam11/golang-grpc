proto:
	rm -f pb/*.go
	#rm -f doc/swagger/*.swagger.json
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto
evans:
	evans --host localhost --port 1999 -r repl

.PHONY: proto evans