package modules

import "fmt"

func FindSubdomains(domain string) []string {
    fmt.Println("Finding subdomains for", domain)
    return []string{"sub1." + domain, "sub2." + domain}
}
