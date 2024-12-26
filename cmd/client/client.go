package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const baseURL = "http://localhost:1111/"
const timeout = 15 * time.Second

type DecodeInput struct {
	EncodedString string `json:"inputString"`
}

type DecodeOutput struct {
	DecodedString string `json:"outputString"`
}

func main() {
	versionResponse, err := fetchVersion()
	if err != nil {
		log.Fatalf("Ошибка извлечения версии: %v", err)
	}
	fmt.Printf("Version: %s\n", versionResponse)

	decodedString, err := decodeBase64("aGVsbG8gd29ybGQ=")
	if err != nil {
		log.Fatalf("Ошибка декодирования base64: %v", err)
	}
	fmt.Printf("Decoded string: %s\n", decodedString)

	err = performHardOperationWithTimeout()
	if err != nil {
		fmt.Printf("Превышено время операция или ошибка: %v\n", err)
	}
}

func fetchVersion() (string, error) {
	res, err := http.Get(baseURL + "version")
	if err != nil {
		return "", fmt.Errorf("Ошибка извелечения версия: %w", err)
	}
	defer res.Body.Close()

	versionData, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("Ошибка чтения версии запроса: %w", err)
	}

	return string(versionData), nil
}

func decodeBase64(encoded string) (string, error) {
	decodeRequest := DecodeInput{EncodedString: encoded}
	reqBody, _ := json.Marshal(decodeRequest)

	res, err := http.Post(baseURL+"decode", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return "", fmt.Errorf("ошибка оптравки в /decode: %w", err)
	}
	defer res.Body.Close()

	var decodeResponse DecodeOutput
	if err := json.NewDecoder(res.Body).Decode(&decodeResponse); err != nil {
		return "", fmt.Errorf("Ошибка декадирования ответа: %w", err)
	}

	return decodeResponse.DecodedString, nil
}

func performHardOperationWithTimeout() error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, baseURL+"hard-op", nil)
	if err != nil {
		return fmt.Errorf("Ошибка создания запроса: %w", err)
	}

	client := &http.Client{}
	res, err := client.Do(req)

	if err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("Запрос к /hard-op превышен по времени")
		}
		return fmt.Errorf("Ошибка преобразоваия /hard-op: %w", err)
	}
	defer res.Body.Close()

	responseBody, _ := io.ReadAll(res.Body)
	fmt.Printf("Hard-op запрос: %s, статус кода: %d\n", responseBody, res.StatusCode)
	return nil
}
