run:
	go run cmd/server/main.go

migrate-up:
	goose -dir ./migrations postgres "user=postgres password=postgres dbname=nongki_db sslmode=disable" up

migrate-down:
	goose -dir ./migrations postgres "user=postgres password=postgres dbname=nongki_db sslmode=disable" down

