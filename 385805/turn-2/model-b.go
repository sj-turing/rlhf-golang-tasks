import (
	"database/sql"
	"fmt"
	"log"
	"runtime/pprof"
	_ "github.com/lib/pq"
)

func example_profileTransaction() {
	db, err := sql.Open("postgres", "user=YourDBUser password=YourDBPassword dbname=YourDBName sslmode=disable")
	if err != nil {
		log.Fatalf("Error opening connection: %v", err)
	}
	defer db.Close()

	f, err := os.Create("CPUprofile")
	if err != nil {
		log.Fatal("Could not create profile file:", err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal("Could not start CPU profile:", err)
	}

	defer pprof.StopCPUProfile()

	executeSQLQueries(db) // or, your state managing using SQL queries.
}
