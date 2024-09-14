generate-ent:
	go run entgo.io/ent/cmd/ent generate ./ent/schema --feature sql/upsert

build:
	go build -o bin/myapp main.go

run:
	go run main.go