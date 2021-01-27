build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ./cmd/web

db_migrate_down:
	go run cmd/migrate/migrate.go down

db_drop:
	go run cmd/migrate/migrate.go drop

db_show_version:
	go run cmd/migrate/migrate.go version

db_migrate_up:
	go run cmd/migrate/migrate.go up

run:
	go run cmd/web/main.go -debug true
