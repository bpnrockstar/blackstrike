package modules

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "blackstrike/modules/utils"
)

func FindSubdomains(domain string) []string {
    fmt.Println("[+] Using crt.sh to find subdomains...")
    url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("[-] Failed to query crt.sh:", err)
        return nil
    }
    defer resp.Body.Close()
    body, _ := ioutil.ReadAll(resp.Body)

    var entries []map[string]interface{}
    _ = json.Unmarshal(body, &entries)

    found := make(map[string]bool)
    for _, entry := range entries {
        if nameValue, ok := entry["name_value"].(string); ok {
            for _, sub := range strings.Split(nameValue, "\n") {
                if strings.HasSuffix(sub, domain) {
                    found[sub] = true
                }
            }
        }
    }

    var result []string
    for sub := range found {
        result = append(result, sub)
    }

    utils.SaveToFile("output/subdomains.txt", result)
    return result
}