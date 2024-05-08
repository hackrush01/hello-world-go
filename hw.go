package main

import "fmt"
import "os"
import "flag"


func main() {
    fmt.Println("Hello World!")

    filePath := flag.String("filePath", "", "path of the json report")
    flag.Parse()

    data, err := os.ReadFile(*filePath)
    if err != nil {
        fmt.Printf("os.ReadFile(%s): %v\n", *filePath, err)
        os.Exit(1)
    }
    fmt.Println(string(data))
}
