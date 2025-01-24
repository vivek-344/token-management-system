include app.env

postgres:
	docker run --name postgres17 -p 5432:5432 -e POSTGRES_USER=${POSTGRES_USER} -e POSTGRES_PASSWORD=${POSTGRES_PASSWORD} -d postgres

createdb:
	docker exec -it postgres17 createdb --username=${POSTGRES_USER} --owner=${POSTGRES_USER} ${POSTGRES_DB}

dropdb:
	docker exec -it postgres17 dropdb ${POSTGRES_DB}

migrateup:
	migrate -path db/migration -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/${POSTGRES_DB}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://$POSTGRES_USER:$POSTGRES_PASSWORD@localhost:5432/${POSTGRES_DB}?sslmode=disable" -verbose down

.PHONY: postgres createdb dropdb migrateup migratedown