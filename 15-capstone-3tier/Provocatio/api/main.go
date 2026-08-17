package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func main() {
	dsn := os.Getenv("DATABASE_URL") // 예: postgres://app:secret@db:5432/app?sslmode=disable
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Postgres 컨테이너가 준비될 때까지 잠깐 재시도
	for i := 0; i < 30; i++ {
		if err = db.Ping(); err == nil {
			break
		}
		log.Println("DB 대기 중...", err)
		time.Sleep(2 * time.Second)
	}
	if err != nil {
		log.Fatal("DB 연결 실패: ", err)
	}

	if _, err := db.Exec(`CREATE TABLE IF NOT EXISTS counter (id int PRIMARY KEY, n int)`); err != nil {
		log.Fatal("init: ", err)
	}
	db.Exec(`INSERT INTO counter (id, n) VALUES (1, 0) ON CONFLICT (id) DO NOTHING`)

	http.HandleFunc("/next", func(w http.ResponseWriter, r *http.Request) {
		var n int
		if err := db.QueryRow(`UPDATE counter SET n = n + 1 WHERE id = 1 RETURNING n`).Scan(&n); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%d", n)
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		// DB 연결까지 확인하고 싶다면:
		if err := db.Ping(); err != nil {
			w.WriteHeader(http.StatusServiceUnavailable)
			fmt.Fprintln(w, "DB unreachable")
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, "OK")
	})

	log.Println("api listening on :6000")
	http.ListenAndServe(":6000", nil)
}
