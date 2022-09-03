package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := os.Mkdir("serverdata", os.ModePerm); err != nil {
		log.Fatal(err)
	}

	createRandomFile("serverdata/test.txt")
	log.Println("-> File created in the server")

	router := gin.Default()
	router.GET("/fileWithCheckSum", getFile)

	err := router.Run(":8081")
	if err != nil {
		return
	}
}

func getFile(c *gin.Context) {

	fileInfo, err := os.ReadFile("serverdata/test.txt")
	checksum := calculateCheckSum("serverdata/test.txt")

	if err == nil {
		var sendFile = FileWithCheckSum{FileInfo: string(fileInfo), Checksum: checksum}

		c.IndentedJSON(http.StatusOK, sendFile)
	}
}
