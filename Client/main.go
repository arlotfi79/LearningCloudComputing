package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	requestURL := fmt.Sprintf("http://Server:%d/fileWithCheckSum", 8081)
	res, err := http.Get(requestURL)
	if err != nil {
		fmt.Printf("-> Error making http request: %s\n", err)
		os.Exit(1)
	}

	if res.StatusCode == 200 {
		log.Println("-> File received successfully")

		var data = FileWithCheckSum{}
		json.NewDecoder(res.Body).Decode(&data)

		if err := os.Mkdir("clientdata", os.ModePerm); err != nil {
			log.Fatal(err)
		}

		createFileInClient("clientdata/test.txt", data.FileInfo)

		if data.Checksum == calculateCheckSum("clientdata/test.txt") {
			log.Println("-> File saved successfully")
		}
	}

}
