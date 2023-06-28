hello:
	echo "hello"

gen-api:
	oapi-codegen -generate types,fiber -package api ./api/api-spec.yml > ./api/server.gen.go

gen-entity:
	go generate ./ent