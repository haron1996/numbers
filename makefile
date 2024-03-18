run:
	go run main.go

tidy:
	go mod tidy

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

deduplicate:
	sort -u random.csv -o random.csv

countcsv:
	wc -l random.csv