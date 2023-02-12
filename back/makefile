postgres:
	docker run --name postgres-clothes -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres-clothes createdb --username=postgres --owner=postgres clothes-app

dropdb:
	docker exec -it postgres-clothes dropdb -U postgres clothes-app

postgres-terminal:
	docker exec -it postgres-clothes psql -U postgres

.PHONY: postgres createdb dropdb postgres-terminal

