package main

import (
	"log"
	"net/http"
)

func (app *Config) SendMail(w http.ResponseWriter, r *http.Request) {

	type MailMessage struct {
		From    string `json:"from"`
		To      string `json:"to"`
		Subject string `json:"subject"`
		Message string `json:"message"`
	}

	var requestPayload MailMessage

	err := app.readJSON(w, r, &requestPayload)
	log.Printf("err: %v\n", err)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	msg := Message{
		From:    requestPayload.From,
		To:      requestPayload.To,
		Subject: requestPayload.Subject,
		Data:    requestPayload.Message,
	}

	err = app.Mailer.SendSMTPMessage(msg)
	log.Printf("err: %v\n", err)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println("Mail sent")

	payload := jsonResponse{
		Error:   false,
		Message: "Mail sent",
	}

	_ = app.writeJSON(w, http.StatusAccepted, payload)
}
