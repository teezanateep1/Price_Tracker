
package routes

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
)

func SearchProducts() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        q := r.URL.Query().Get("q")
        api := "https://shopee.co.th/api/v4/search/search_items?by=relevancy&limit=10&newest=0&order=desc&page_type=search&scenario=PAGE_GLOBAL_SEARCH&version=2&keyword=" + url.QueryEscape(q)

        req, _ := http.NewRequest("GET", api, nil)
        req.Header.Set("User-Agent", "Mozilla/5.0")
        req.Header.Set("Referer", "https://shopee.co.th/search?keyword=" + url.QueryEscape(q))

        resp, err := http.DefaultClient.Do(req)
        if err != nil {
            http.Error(w, err.Error(), 500)
            return
        }
        defer resp.Body.Close()

        body, _ := io.ReadAll(resp.Body)
        var data map[string]interface{}
        json.Unmarshal(body, &data)

        items := []map[string]interface{}{}
        if res, ok := data["items"].([]interface{}); ok {
            for _, item := range res {
                it := item.(map[string]interface{})["item_basic"].(map[string]interface{})
                itemId := int(it["itemid"].(float64))
                shopId := int(it["shopid"].(float64))
                name := it["name"].(string)
                price := int(it["price"].(float64) / 100000) // price is in cents
                image := fmt.Sprintf("https://cf.shopee.co.th/file/%s", it["image"].(string))
                affiliate := fmt.Sprintf("https://s.shopee.co.th/6AZVePoJNQ?sub_id1=peerapon-web&itemid=%d&shopid=%d", itemId, shopId)

                items = append(items, map[string]interface{}{
                    "name": name,
                    "price": price,
                    "image": image,
                    "itemid": itemId,
                    "shopid": shopId,
                    "affiliate_url": affiliate,
                })
            }
        }

        json.NewEncoder(w).Encode(items)
    }
}
