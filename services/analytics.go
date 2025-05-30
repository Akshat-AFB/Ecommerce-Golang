package services

import (
	"backend-go/types"
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

const lambdaURL = "https://tvyhkxovv5qp7rwrgjh42wq6he0lpzni.lambda-url.ap-south-1.on.aws/"


func EmitLoginEvent(userID uint) {
	payload := types.LoginEvent{
		UserID:    userID,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
	}
	log.Println("Payload", payload)
	body, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error marshalling login event:", err)
		return
	}

	resp, err := http.Post(lambdaURL, "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println("Error sending login event:", err)
		return
	}
	defer resp.Body.Close()
	log.Println("Error ",err)
	log.Println("StatusCode ", resp.StatusCode)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Login event failed with status:", resp.Status)
	}
}
