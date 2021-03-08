migrateup:
	migrate -path migrations --database "postgresql://postgres:example@localhost:5155/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations --database "postgresql://postgres:example@localhost:5155/postgres?sslmode=disable" -verbose down

formatdb:
	make migratedown && make migrateup
