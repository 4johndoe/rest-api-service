migrate-up:
	migrate -database "postgres://127.0.0.1:5434/postgres?sslmode=disable&user=postgres&password=postgres" -path ./migrations up
