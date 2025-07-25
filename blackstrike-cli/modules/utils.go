package modules

import (
    "fmt"
    "os"
)

func SaveToFile(path string, data []string) {
    f, _ := os.Create(path)
    defer f.Close()
    for _, line := range data {
        f.WriteString(line + "\n")
    }
    fmt.Println("Saved data to", path)
}
