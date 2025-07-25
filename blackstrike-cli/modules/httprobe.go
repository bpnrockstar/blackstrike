package modules

import "fmt"

func HTTPProbe(subdomains []string) []string {
    fmt.Println("Probing HTTP endpoints")
    return subdomains
}
