hello:
	echo "hello"

gen-api:
	oapi-codegen -generate types,fiber -package api ./api/api-spec.yml > ./api/types.generated.go

gen-entity:
	go generate ./ent

run-db:
	docker run -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

atlas-schema:
	atlas schema inspect \
	-u "ent://ent/schema" \
	--dev-url "postgresql://postgres:password@localhost/postgres?search_path=public&sslmode=disable" \
	--format '{{ sql . "  " }}'