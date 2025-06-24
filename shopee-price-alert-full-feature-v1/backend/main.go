package main

import (
    "log"
    "net/http"
    "shopee-price-alert/handlers"
    "shopee-price-alert/db"
    "shopee-price-alert/jobs"
    "time"
)

func main() {
    database, err := db.InitDB("data.db")
    if err != nil {
        log.Fatal(err)
    }

    go func() {
        for {
            jobs.CheckPrices(database)
            time.Sleep(3600 * time.Second)
        }
    }()

    http.HandleFunc("/api/search", handlers.SearchHandler)
    http.HandleFunc("/api/subscribe", handlers.SubscribeHandler)

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
