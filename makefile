postgres:
	docker run --name postgres-clothes -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres-clothes createdb --username=postgres --owner=postgres clothesapp

dropdb:
	docker exec -it postgres-clothes dropdb -U postgres clothesapp

postgres-terminal:
	docker exec -it postgres-clothes psql -U postgres clothesapp

.PHONY: postgres createdb dropdb postgres-terminal 

