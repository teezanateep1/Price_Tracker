package utils

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
)

type GraphQLRequest struct {
    OperationName string                 `json:"operationName"`
    Query         string                 `json:"query"`
    Variables     map[string]interface{} `json:"variables"`
}

type GraphQLResponse struct {
    Data struct {
        ProductOfferLinks []struct {
            ItemID           string `json:"itemId"`
            ShopID           int    `json:"shopId"`
            ProductOfferLink string `json:"productOfferLink"`
        } `json:"productOfferLinks"`
    } `json:"data"`
}

func GetAffiliateLink(shopID int, itemID int) (string, error) {
    cookie := os.Getenv("SHOPEE_COOKIE")
    if cookie == "" {
        return "", fmt.Errorf("Shopee cookie not set in env")
    }

    gql := GraphQLRequest{
        OperationName: "batchGetProductOfferLink",
        Query: `
      query batchGetProductOfferLink (
        $sourceCaller: SourceCaller!
        $productOfferLinkParams: [ProductOfferLinkParam!]!
        $advancedLinkParams: AdvancedLinkParams
      ){
        productOfferLinks(
          productOfferLinkParams: $productOfferLinkParams,
          sourceCaller: $sourceCaller,
          advancedLinkParams: $advancedLinkParams
        ) {
          itemId
          shopId
          productOfferLink
        }
      }`,
        Variables: map[string]interface{}{
            "productOfferLinkParams": []map[string]interface{}{
                {
                    "itemId": fmt.Sprintf("%d", itemID),
                    "shopId": shopID,
                    "trace":  "{}",
                },
            },
            "sourceCaller": "WEB_SITE_CALLER",
            "advancedLinkParams": map[string]string{
                "subId1": "",
                "subId2": "",
                "subId3": "",
                "subId4": "",
                "subId5": "",
            },
        },
    }

    body, _ := json.Marshal(gql)

    req, err := http.NewRequest("POST", "https://affiliate.shopee.co.th/api/v3/gql?q=batchGetProductOfferLink", bytes.NewBuffer(body))
    if err != nil {
        return "", err
    }

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("X-API-SOURCE", "AFFILIATE_DASHBOARD")
    req.Header.Set("X-Requested-With", "XMLHttpRequest")
    req.Header.Set("Cookie", cookie)

    client := &http.Client{}
    resp, err := client.Do(req)
    if resp.StatusCode == 401 || resp.StatusCode == 403 {
        NotifyExpiredCookie()
    }
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    resBody, _ := ioutil.ReadAll(resp.Body)
    var gqlResp GraphQLResponse
    err = json.Unmarshal(resBody, &gqlResp)
    if err != nil {
        return "", err
    }

    if len(gqlResp.Data.ProductOfferLinks) == 0 {
        return "", fmt.Errorf("no affiliate link returned")
    }

    return gqlResp.Data.ProductOfferLinks[0].ProductOfferLink, nil
}
