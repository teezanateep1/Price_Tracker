package utils

import (
    "bytes"
    "net/http"
    "net/url"
)

func SendLineNotify(token, message string) error {
    endpoint := "https://notify-api.line.me/api/notify"
    data := url.Values{}
    data.Set("message", message)

    req, err := http.NewRequest("POST", endpoint, bytes.NewBufferString(data.Encode()))
    if err != nil {
        return err
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Authorization", "Bearer " + token)

    client := &http.Client{}
    _, err = client.Do(req)
    return err
}


package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

func NotifyExpiredCookie() {
    token := os.Getenv("LINE_NOTIFY_TOKEN")
    if token == "" {
        fmt.Println("LINE_NOTIFY_TOKEN not set")
        return
    }
    message := "⚠️ Shopee Affiliate Cookie หมดอายุ หรือใช้ไม่ได้แล้ว\nกรุณาอัปเดต .env ใหม่"
    SendLineNotify(token, message)
}


import "database/sql"

func CreateNotification(db *sql.DB, userID string, productName string, message string) {
    db.Exec("INSERT INTO notifications(user_id, product_name, message) VALUES (?, ?, ?)", userID, productName, message)
}
