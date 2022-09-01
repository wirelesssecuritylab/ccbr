```go
   func InitDB() *sqlx.DB {
	database, err := sqlx.Open("mysql", "mysql:test123789@tcp(127.0.0.1:3306)/policymanager")
	if err != nil {
		fmt.Println("open mysql failed,", err)

	}
	return database
}
   ```
