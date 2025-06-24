package jobs

import (
    "database/sql"
    "fmt"
    "shopee-price-alert/utils"
)

func CheckPrices(db *sql.DB) {
    rows, _ := db.Query(`
        SELECT p.id, p.name, p.url, p.price, s.line_token, s.alert_type, s.alert_threshold, s.user_id 
        FROM products p 
        JOIN subscriptions s ON p.id = s.product_id
    `)
    defer rows.Close()

    for rows.Next() {
        var productID int
        var name, url string
        var oldPrice int
        var token, alertType string
        var threshold int
        var userID string

        rows.Scan(&productID, &name, &url, &oldPrice, &token, &alertType, &threshold, &userID)

        // mock current price for demo (replace with real scraping or API)
        currentPrice := oldPrice - 100 // ลดไป 100 บาท

        notify := false

        if alertType == "any_change" && currentPrice < oldPrice {
            notify = true
        } else if alertType == "percent" {
            drop := float64(oldPrice-currentPrice) / float64(oldPrice) * 100
            if int(drop) >= threshold {
                notify = true
            }
        }

        if notify {
            msg := fmt.Sprintf("📉 '%s' ลดราคาจาก %d → %d บาท\n%s", name, oldPrice, currentPrice, url)
            utils.SendLineNotify(token, msg)
            utils.CreateNotification(db, userID, name, msg)
            db.Exec("UPDATE products SET price = ? WHERE id = ?", currentPrice, productID)
        }
    }
}
