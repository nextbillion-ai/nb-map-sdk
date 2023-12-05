package main

import (
	"fmt"
	"os"

	// structs "github.com/nextbillion-ai/nb-optimization-interface"
	"github.com/nextbillion-ai/nb-optimization-interface/structs"
	client "github.com/nextbillion-ai/nb-optimization-sdk/src"
	// validations "github.com/nextbillion-ai/nb-optimization-interface/validations"
)


func main() {
	apikey:=os.Getenv("APIKEY")
	sdkClient := client.NewClient("https://api.nextbillion.io", apikey)
	// Create an instance of OptimizationPostInput
	input := structs.OptimizationPostInput{
		Locations: structs.Locations{
			Id:       50,
			Location: "19.38972030,-99.14369731|19.42566371,-99.17081981|19.42436859,-99.13374095",
		},
		Vehicles: []structs.Vehicle{
			{
				Id:         1,
				StartIndex: 0,
				EndIndex:   2,
				Capacity:   []int64{4},
				TimeWindow: []uint64{1656500400, 1656502860},
			},
			{
				Id:         2,
				StartIndex: 1,
				EndIndex:   2,
				Capacity:   []int64{4},
				TimeWindow: []uint64{1656500400, 1656502860},
			},
		},
		Jobs: []structs.Job{
			{
				Id:            1,
				LocationIndex: 0,
				Service:       300,
				TimeWindows:   [][]uint64{{1656493200, 1656496800}},
			},
			{
				Id:            2,
				LocationIndex: 1,
				Service:       300,
				TimeWindows:   [][]uint64{{1656502200, 1656507600}},
			},
		},
	}

	// Make a POST request
	resp, err := sdkClient.MakeRequest(fmt.Sprintf("/optimise-mvrp?key=%s",apikey),"POST", input)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("Response status:", resp.Status)
}

