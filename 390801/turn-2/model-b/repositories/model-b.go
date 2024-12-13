package main

func main() {
	// Before
	userRepo := repositories.NewSQLUserRepository(db)

	// After
	userRepo, err := repositories.NewMongoUserRepository("mongodb://localhost:27017", "myapp")
	if err != nil {
		panic(fmt.Errorf("failed to initialize MongoDB repository: %w", err))
	}
}
