package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	//read api key from env
	apiKey := os.Getenv("TWINROOM_API_KEY")
	twinRoomUrlBase := os.Getenv("TWINROOM_API_URL_BASE")
	r.GET("/stream", func(c *gin.Context) {
		client := &http.Client{}
		query := "Hi! Are you fine?"
		data := map[string]string{
			"session_id": "sample_session_id", // Twinroom users can determine this value
			"user_id":    "sample_user_id",    // Twinroom users can determine this value
			"content":    query,
		}
		jsonData, err := json.Marshal(data)

		if err != nil {
			log.Fatalf("Error occurred during marshaling. Error: %s", err.Error())
		}
		req, err := http.NewRequest("POST", twinRoomUrlBase+"/api/v1/message", bytes.NewBuffer(jsonData))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Add("api-key", apiKey)

		resp, err := client.Do(req)
		fmt.Println(*resp)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		// process SSE response by Twinroom Server
		reader := bufio.NewReader(resp.Body)

		var msg string
		var voice string

		for {
			line, err := reader.ReadBytes('\n')
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}

			msgPart := getTargetPartFromResString(string(line), TargetMsg)
			if msgPart != nil {
				msg += *msgPart
			}
			voicePart := getTargetPartFromResString(string(line), TargetVoice)
			if voicePart != nil {
				voice += *voicePart
				fmt.Println(*voicePart)
				saveBase64EncodedAudio(*voicePart)
			}
		}

		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"msg":   msg,
			"query": query,
		})
	})

	r.Run(":8083") // 8083ポートでリッスン
}
