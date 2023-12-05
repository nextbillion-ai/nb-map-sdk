package main

import (
	"fmt"
	"io"
	"os"

	// structs "github.com/nextbillion-ai/nb-optimization-interface"
	"github.com/nextbillion-ai/nb-optimization-interface/structs"
	client "github.com/nextbillion-ai/nb-optimization-sdk/src"
	// validations "github.com/nextbillion-ai/nb-optimization-interface/validations"
)

func mkUint64Ptr(value uint64) *uint64 {
	return &value
}

func main() {
	apikey := os.Getenv("APIKEY")
	sdkClient := client.NewClient("https://api.nextbillion.io", apikey)
	/*
			{
		  "locations": {
		    "id": 50,
		    "description": "Test",
		    "location": "19.38972030,-99.14369731|19.42566371,-99.17081981|19.42436859,-99.13374095"
		  },
		  "vehicles": [{
		    "id": 1,
		    "start_index": 0,
		    "end_index": 2
		  }, {
		    "id": 2,
		    "start_index": 1,
		    "end_index": 2
		  }
		  ],
		  "jobs": [{
		    "id": 1,
		    "location_index": 0
		  },{
		    "id": 2,
		    "location_index": 1
		  }
		  ]
		}
	*/
	// Create an instance of OptimizationPostInput
	input := structs.OptimizationPostInput{
		Locations: structs.Locations{
			Id:              50,
			AnyTypeLocation: "19.38972030,-99.14369731|19.42566371,-99.17081981|19.42436859,-99.13374095",
			Location:        "19.38972030,-99.14369731|19.42566371,-99.17081981|19.42436859,-99.13374095",
		},
		Vehicles: []structs.Vehicle{
			{
				Id:         1,
				StartIndex: mkUint64Ptr(0),
				EndIndex:   mkUint64Ptr(2),
			},
			{
				Id:         2,
				StartIndex: mkUint64Ptr(1),
				EndIndex:   mkUint64Ptr(2),
			},
		},
		Jobs: []structs.Job{
			{
				Id:            1,
				LocationIndex: 0,
			},
			{
				Id:            2,
				LocationIndex: 1,
			},
		},
	}

	// Make a POST request
	resp, err := sdkClient.MakeRequest(fmt.Sprintf("/optimise-mvrp?key=%s", apikey), "POST", input)
	if resp != nil {
		defer resp.Body.Close()
		body, re := io.ReadAll(resp.Body)
		if re != nil {
			fmt.Printf("Error getting response body:%s\n", re)
		}
		fmt.Printf("response body: %s \n", string(body))
	}
	if err != nil {
		fmt.Printf("Error making request:%s\n", err)
		return
	}
	fmt.Println("Response status:", resp.Status)
}
