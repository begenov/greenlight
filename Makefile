postgres:
	sudo docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createbd: 
	sudo docker exec -it postgres createdb --username=root --owner=root greenlight

.PHONY: postgres