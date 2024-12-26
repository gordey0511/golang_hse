package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type DecodeRequest struct {
	Input string `json:"inputString"`
}

type DecodeResponse struct {
	Output string `json:"outputString"`
}

const apiVersion = "v1.0.0"
const serverPort = 1111

func main() {
	router := http.NewServeMux()

	router.HandleFunc("/version", versionHandler)
	router.HandleFunc("/decode", decodeHandler)
	router.HandleFunc("/hard-op", hardOperationHandler)

	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", serverPort),
		Handler: router,
	}

	stopSignal := make(chan os.Signal, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Ошибка сервера: %v", err)
		}
	}()

	log.Printf("Сервер запущен на порте %d", serverPort)

	<-stopSignal
	log.Println("Сервер упал...")

	shutdownContext, shutdownCancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownCancel()

	if err := server.Shutdown(shutdownContext); err != nil {
		log.Fatalf("Forced shutdown: %v", err)
	}

	log.Println("Сервер остановлен")
}

func versionHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, apiVersion)
}

func decodeHandler(w http.ResponseWriter, r *http.Request) {
	var requestData DecodeRequest
	if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
		http.Error(w, "Неверный формат запроса", http.StatusBadRequest)
		return
	}

	decodedData, err := base64.StdEncoding.DecodeString(requestData.Input)
	if err != nil {
		http.Error(w, "Неверные данные в формате base64", http.StatusBadRequest)
		return
	}

	responseData := DecodeResponse{Output: string(decodedData)}
	if err := json.NewEncoder(w).Encode(responseData); err != nil {
		http.Error(w, "Ошибка кодирования запроса", http.StatusInternalServerError)
		return
	}
}

func hardOperationHandler(w http.ResponseWriter, r *http.Request) {
	delay := time.Duration(10+rand.Intn(11)) * time.Second
	time.Sleep(delay)

	if rand.Intn(2) == 0 {
		http.Error(w, "Внутренняя ошибка сервера", http.StatusInternalServerError)
	} else {
		fmt.Fprintln(w, "Успешно")
	}
}
