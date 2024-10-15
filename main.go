package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/414-n/go_final_project/server/server"
	"github.com/joho/godotenv"
)

var webDir = "web"

func main() {

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("TODO_PORT")
	if port == "" {
		port = "7540"
	}

	dbFile := os.Getenv("TODO_DBFILE")
	if dbFile == "" {
		log.Fatal("TODO_DBFILE environment variable is not set")
	}

	db, err := server.InitDB(dbFile)
	if err != nil {
		log.Fatal("Error initializing database:", err)
	}
	defer db.Close()

	s := server.NewServer(db)

	http.Handle("/", http.FileServer(http.Dir(webDir)))
	http.HandleFunc("/api/tasks", s.HandleTasks)

	fmt.Printf("Server is listening on port %s\n", port)
	fmt.Printf("Open http://localhost:%s/\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
