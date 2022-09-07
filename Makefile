build:
	go build -o bin/softdare  ./cmd/softdare

run:
	./bin/softdare

start:
	go build -o bin/softdare  ./cmd/softdare && ./bin/softdare
