package main

import (
	"fmt"

	client "github.com/nextbillion-ai/nb-optimization-sdk/src"
)

func main() {
	client := client.NewClient("https://sgpstg.nextbillion.io/optimise-mvrp?key=plaintesting") // Replace with the actual base URL of your server

	// Make a request (example: GET /optimise-mvrp/result)
	resp, err := client.MakeRequest("/optimise-mvrp/result", "GET")
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}
