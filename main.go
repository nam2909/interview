// counter_app.go
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

const counterFile = "counter.txt"

var (
	mu    sync.Mutex
	count int
)

func main() {
	// Load initial count
	c, err := loadCount(counterFile)
	if err != nil {
		log.Printf("Could not load counter, starting at 0: %v", err)
	}
	count = c

	http.HandleFunc("/", viewHandler)
	http.HandleFunc("/increment", incrementHandler)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// viewHandler writes a simple HTML page with the current count and a button.
func viewHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, err := fmt.Fprintf(w, `<html>
	<head><title>Counter</title></head>
	<body>
	  <h1>Count: %d</h1>
	  <form action="/increment" method="POST">
	    <button type="submit">+1</button>
	  </form>
	</body>
	</html>`, count)
	if err != nil {
		return
	}
}

// incrementHandler increases the counter, saves it, and redirects back.
func incrementHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	mu.Lock()
	count++
	if err := saveCount(counterFile, count); err != nil {
		log.Printf("Error saving count: %v", err)
	}
	mu.Unlock()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// loadCount reads the counter from a file.
// If the file doesn't exist, returns 0.
func loadCount(filename string) (int, error) {
	data, err := os.ReadFile(filename)
	if os.IsNotExist(err) {
		return 0, nil
	} else if err != nil {
		return 0, err
	}
	c, err := strconv.Atoi(string(data))
	if err != nil {
		return 0, err
	}
	return c, nil
}

// saveCount writes the counter to a file atomically.
func saveCount(filename string, c int) error {
	temp := []byte(strconv.Itoa(c))
	return os.WriteFile(filename, temp, 0644)
}
