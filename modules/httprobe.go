package modules

import (
    "fmt"
    "net/http"
    "sync"
    "time"
    "blackstrike/modules/utils"
)

func ProbeHosts(domains []string) []string {
    fmt.Println("[+] Probing for live hosts...")
    var wg sync.WaitGroup
    var mu sync.Mutex
    live := []string{}

    client := &http.Client{
        Timeout: 5 * time.Second,
    }

    for _, domain := range domains {
        wg.Add(1)
        go func(d string) {
            defer wg.Done()
            url := "http://" + d
            resp, err := client.Get(url)
            if err == nil && resp.StatusCode < 500 {
                mu.Lock()
                fmt.Println("[LIVE] " + url)
                live = append(live, url)
                mu.Unlock()
            }
        }(domain)
    }
    wg.Wait()
    utils.SaveToFile("output/live_hosts.txt", live)
    return live
}