package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/mathcale/goexpert-course/mutithreading/internal/pkg/httpclient"
)

var ZIP_CODE string = "26700000"

func main() {
	viaCepChan := make(chan interface{})
	apiCepChan := make(chan interface{})

	go func() {
		client := httpclient.NewHttpClient("https://viacep.com.br")
		endpoint := fmt.Sprintf("/ws/%s/json/", ZIP_CODE)

		if err := client.Get(endpoint, viaCepChan); err != nil {
			panic(err)
		}
	}()

	go func() {
		client := httpclient.NewHttpClient("https://cdn.apicep.com")
		endpoint := fmt.Sprintf("/file/apicep/%s.json", ZIP_CODE)

		if err := client.Get(endpoint, apiCepChan); err != nil {
			panic(err)
		}
	}()

	select {
	case response := <-viaCepChan:
		fmt.Println("Got response from [viacep] API:")
		fmt.Printf("%s\n", formatResponse(response))
	case response := <-apiCepChan:
		fmt.Println("Got response from [apicep] API:")
		fmt.Printf("%s\n", formatResponse(response))
	case <-time.After(1000 * time.Millisecond):
		fmt.Println("Timeout")
	}
}

func formatResponse(response interface{}) string {
	responseAsJsonStr, marshalErr := json.MarshalIndent(response, "", "  ")

	if marshalErr != nil {
		panic(marshalErr)
	}

	return string(responseAsJsonStr)
}
