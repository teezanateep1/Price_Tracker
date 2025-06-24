package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func InitDB(filepath string) (*sql.DB, error) {
    db, err := sql.Open("sqlite3", filepath)
    if err != nil {
        return nil, err
    }

    createProductTable := `
    CREATE TABLE IF NOT EXISTS products (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        url TEXT,
        price INTEGER
    );`
    _, err = db.Exec(createProductTable)
    if err != nil {
        return nil, err
    }

    createSubsTable := `
    CREATE TABLE IF NOT EXISTS subscriptions (
        alert_type TEXT DEFAULT 'any_change',
        alert_threshold INTEGER DEFAULT 10,
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        product_id INTEGER,
        line_token TEXT,
        notify_interval TEXT
    );`
    _, err = db.Exec(createSubsTable)
    if err != nil {
        return nil, err
    }

    
    createNotificationsTable := `
    CREATE TABLE IF NOT EXISTS notifications (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        user_id TEXT,
        product_name TEXT,
        message TEXT,
        is_read BOOLEAN DEFAULT 0,
        created_at DATETIME DEFAULT CURRENT_TIMESTAMP
    );`
    _, err = db.Exec(createNotificationsTable)
    if err != nil {
        return nil, err
    }
    
    return db, nil
}
