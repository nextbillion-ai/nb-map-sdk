package main

import (
    "fmt"
    "github.com/nextbillion-ai/nb-optimization-sdk/src"
)

func main() {
    // Initialize the client
    client := client.NewClient("http://localhost:8080") // Replace with the actual base URL of your server

    // Make a request (example: GET /optimise-mvrp/result)
    resp, err := client.MakeRequest("/optimise-mvrp/result", "GET")
    if err != nil {
        fmt.Println("Error making request:", err)
        return
    }
    defer resp.Body.Close()

    // Handle the response
    fmt.Println("Response status:", resp.Status)
    // Add more logic to handle the response body as needed
}
