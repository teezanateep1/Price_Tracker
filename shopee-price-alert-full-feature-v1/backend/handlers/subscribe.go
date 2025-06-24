package handlers

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "shopee-price-alert/db"
)

type SubscribeRequest struct {
    Name     string `json:"name"`
    URL      string `json:"url"`
    Price    int    `json:"price"`
    Token    string `json:"token"`
    Interval string `json:"interval"`
}

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
    var req SubscribeRequest
    err := json.NewDecoder(r.Body).Decode(&req)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    database, _ := db.InitDB("data.db")

    res, _ := database.Exec("INSERT INTO products(name, url, price) VALUES (?, ?, ?)", req.Name, req.URL, req.Price)
    productID, _ := res.LastInsertId()

    _, err = database.Exec("INSERT INTO subscriptions(product_id, line_token, notify_interval) VALUES (?, ?, ?)", productID, req.Token, req.Interval)
    if err != nil {
        http.Error(w, "Subscribe failed", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Subscribed!"))
}
