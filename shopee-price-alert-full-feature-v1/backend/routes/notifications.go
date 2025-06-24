
package routes

import (
    "database/sql"
    "encoding/json"
    "net/http"
    "strconv"
)

func GetNotifications(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        userID := r.Header.Get("X-User-ID")
        rows, _ := db.Query("SELECT id, product_name, message, is_read, created_at FROM notifications WHERE user_id = ? ORDER BY created_at DESC", userID)
        var notifs []map[string]interface{}
        for rows.Next() {
            var id int
            var name, msg, created string
            var isRead bool
            rows.Scan(&id, &name, &msg, &isRead, &created)
            notifs = append(notifs, map[string]interface{}{
                "id": id, "product_name": name, "message": msg, "is_read": isRead, "created_at": created,
            })
        }
        json.NewEncoder(w).Encode(notifs)
    }
}

func MarkNotificationRead(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        idStr := r.URL.Path[len("/api/notifications/"):]
        id, _ := strconv.Atoi(idStr[:len(idStr)-5]) // remove /read
        db.Exec("UPDATE notifications SET is_read = 1 WHERE id = ?", id)
        w.WriteHeader(http.StatusOK)
    }
}
