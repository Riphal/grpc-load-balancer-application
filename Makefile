migrateup:
	migrate -path migrations --database "postgresql://postgres:example@localhost:5155/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path migrations --database "postgresql://postgres:example@localhost:5155/postgres?sslmode=disable" -verbose down

formatdb:
	make migratedown && make migrateup

LIST = bankAccount
protogen:
	for t in ${LIST}; do \
  		rm ./common/proto/"$$t"/"$$t".pb.go; \
  		protoc -I=./common/proto/ --go_out=plugins=grpc:common/proto ./common/proto/"$$t"/"$$t".proto; \
  	done
