bundle-api-spec:
	scripts/bundle-api-spec.sh

gen-api:
	(cd backend; go generate .)

gen-db:
	(cd backend; go generate ./ent)

gen-all: bundle-api-spec gen-api gen-db

run-db:
	docker run -e POSTGRES_PASSWORD=password -p 5432:5432 -d postgres

atlas-schema:
	atlas schema inspect \
	-u "ent://ent/schema" \
	--dev-url "postgresql://postgres:password@localhost/postgres?search_path=public&sslmode=disable" \
	--format '{{ sql . "  " }}'