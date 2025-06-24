
package routes

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
    "strings"
)

func GetSubscriptions(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Header.Get("X-User-ID")
        rows, _ := db.Query("SELECT s.id, p.name, p.price, s.alert_type, s.alert_threshold FROM subscriptions s JOIN products p ON s.product_id = p.id WHERE s.user_id = ?", userID)
        var subs []map[string]interface{}
        for rows.Next() {
            var id, price, threshold int
            var name, typ string
            rows.Scan(&id, &name, &price, &typ, &threshold)
            subs = append(subs, map[string]interface{}{
                "id": id, "product_name": name, "current_price": price, "alert_type": typ, "alert_threshold": threshold,
            })
        }
        json.NewEncoder(w).Encode(subs)
    }
}

func UpdateSubscription(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        idStr := strings.TrimPrefix(r.URL.Path, "/api/subscriptions/")
        id, _ := strconv.Atoi(idStr)
        var body map[string]interface{}
        json.NewDecoder(r.Body).Decode(&body)
        db.Exec("UPDATE subscriptions SET alert_type = ?, alert_threshold = ? WHERE id = ?", body["alert_type"], body["alert_threshold"], id)
        w.WriteHeader(http.StatusOK)
    }
}
