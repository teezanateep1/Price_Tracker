
package handlers

import (
    "encoding/json"
    "fmt"
    "net/http"
    "shopee-price-alert/utils"
    "shopee-price-alert/db"
    "net/url"
    "strconv"
)

type Product struct {
    Name         string `json:"name"`
    URL          string `json:"url"`
    Price        int    `json:"price"`
    IsMall       bool   `json:"is_mall"`
    IsPreferred  bool   `json:"is_preferred"`
    HasVoucher   bool   `json:"has_voucher"`
}

type ShopeeResponse struct {
    Items []struct {
        ItemBasic struct {
            ItemID                  int    `json:"itemid"`
            ShopID                  int    `json:"shopid"`
            Name                    string `json:"name"`
            Price                   int    `json:"price"`
            ShowOfficialShopLabel  bool   `json:"show_official_shop_label"`
            ShowShopeeVerifiedLabel bool  `json:"show_shopee_verified_label"`
            HasVoucher              bool   `json:"has_voucher"`
        } `json:"item_basic"`
    } `json:"items"`
}

func SearchHandler(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query().Get("q")
    if query == "" {
        http.Error(w, "Missing query", http.StatusBadRequest)
        return
    }

    shopeeURL := "https://shopee.co.th/api/v4/search/search_items"
    params := url.Values{}
    params.Set("by", "relevancy")
    params.Set("keyword", query)
    params.Set("limit", "10")
    params.Set("newest", "0")
    params.Set("order", "desc")
    params.Set("page_type", "search")

    resp, err := http.Get(shopeeURL + "?" + params.Encode())
    if err != nil {
        http.Error(w, "Failed to fetch from Shopee", http.StatusInternalServerError)
        return
    }
    defer resp.Body.Close()

    var shopeeRes ShopeeResponse
    err = json.NewDecoder(resp.Body).Decode(&shopeeRes)
    if err != nil {
        http.Error(w, "Failed to parse response", http.StatusInternalServerError)
        return
    }

    database, _ := db.InitDB("data.db")

    var results []Product
    for _, item := range shopeeRes.Items {
        basic := item.ItemBasic

        if basic.ShowOfficialShopLabel || basic.ShowShopeeVerifiedLabel || basic.HasVoucher {
            // check if cached
            var affLink string
            row := database.QueryRow("SELECT url FROM products WHERE name = ? AND price = ?", basic.Name, basic.Price/100000)
            row.Scan(&affLink)

            if affLink == "" {
                // generate and cache
                affLink, err = utils.GetAffiliateLink(basic.ShopID, basic.ItemID)
                if err != nil {
                    continue
                }
                database.Exec("INSERT INTO products(name, url, price) VALUES (?, ?, ?)", basic.Name, affLink, basic.Price/100000)
            }

            p := Product{
                Name:        basic.Name,
                URL:         affLink,
                Price:       basic.Price / 100000,
                IsMall:      basic.ShowOfficialShopLabel,
                IsPreferred: basic.ShowShopeeVerifiedLabel,
                HasVoucher:  basic.HasVoucher,
            }
            results = append(results, p)
        }
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(results)
}
