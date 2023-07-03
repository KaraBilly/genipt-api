package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Role int

const (
	User Role = iota
	Assistant
	System
)

type ConversationRequest struct {
	Message  []ConversationMessage `json:"message"`
	UserName string                `json:"userName"`
}

type ConversationMessage struct {
	Role    Role   `json:"role"`
	Content string `json:"content"`
}

type ConversationResponse struct {
	Message []ConversationMessage `json:"message"`
}

func main() {
	// 1.创建路由
	r := gin.Default()
	// 2.绑定路由规则，执行的函数
	// gin.Context，封装了request和response
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hello World!")
	})

	r.POST("/Conversation", func(c *gin.Context) {
		var msgs ConversationRequest
		if err := c.BindJSON(&msgs); err != nil {
			log.Fatal(err)
		}

		msgs.UserName = "乔盛飞"

		payload, err := json.Marshal(msgs)
		resp, err := ConversationPost(payload)

		if err != nil {
			fmt.Println(err)
		}
		c.JSON(http.StatusOK, resp)
	})
	// 3.监听端口，默认在8080
	// Run("里面不指定端口号默认为8080")
	r.Run(":8000")
}

func ConversationPost(data []byte) (ConversationResponse, error) {

	var result ConversationResponse
	url := "http://ellis.centralindia.cloudapp.azure.com/openaipass/ChatGPT/Conversation"
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(data))
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiI4NTM0M2NlMi04MzZkLTRjZGItYmI4MS0zNGNhMWIxOThiNWUiLCJ1c2VyaWQiOiIxMjM0IiwibmFtZSI6IuS5lOebm-mjniIsImVtYWlsIjoidGVzdCIsIm1hbmFnZXIiOiIiLCJpYXQiOiIyMDIzLzYvMjkgMDk6MjQ6NTEiLCJodHRwOi8vc2NoZW1hcy5taWNyb3NvZnQuY29tL3dzLzIwMDgvMDYvaWRlbnRpdHkvY2xhaW1zL3JvbGUiOiLnrqHnkIblkZgiLCJyb2xlIjoi566h55CG5ZGYIiwibmJmIjoxNjg4MDAxODkxLCJleHAiOjE2OTA1OTM4OTEsImlzcyI6IlNDTSIsImF1ZCI6IlNDTSJ9.KUN9M3f5xU2YgIS1zcEnRAe98CX865cNdq4wUGzra9s")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return result, err
	}

	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
		return result, err
	}
	return result, nil
}
