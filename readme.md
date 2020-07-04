# Setup databases
1. Setup test database:
	```
	go run ./migrator/main.go db_create test
	go run ./migrator/main.go migrate test
	```

2. Setup development database: `go run ./migrator/main.go db_create development`

# Development
1. Run webpack-dev-server: `yarn run webpack`
2. Run backend server: `go run main.go`
3. Point browser to `http://localhost:8080/decks`

# Typescript
- This project uses TypeScript for static typing and type
  annotation
