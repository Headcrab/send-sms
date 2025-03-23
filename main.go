package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

type SMSRequest struct {
	Recipient string `json:"recipient"`
	Message   string `json:"message"`
}

func main() {
	// Определяем параметры командной строки
	recipient := flag.String("recipient", "", "Номер телефона получателя в формате +XXXXXXXXXXX")
	message := flag.String("message", "", "Текст сообщения")
	serverURL := flag.String("url", "http://192.168.1.1/api/sendsms", "URL API сервера для отправки SMS")
	username := flag.String("user", "root", "Имя пользователя для авторизации")
	password := flag.String("pass", "root", "Пароль для авторизации")

	flag.Parse()

	// Проверяем обязательные параметры
	if *recipient == "" || *message == "" {
		fmt.Println("Ошибка: параметры recipient и message обязательны")
		flag.Usage()
		os.Exit(1)
	}

	// Формируем запрос
	smsReq := SMSRequest{
		Recipient: *recipient,
		Message:   *message,
	}

	// Преобразуем запрос в JSON
	jsonData, err := json.Marshal(smsReq)
	if err != nil {
		fmt.Printf("Ошибка при формировании JSON: %v\n", err)
		os.Exit(1)
	}

	// Создаем HTTP клиент
	client := &http.Client{}

	// Формируем URL
	url := *serverURL

	// Создаем запрос
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Ошибка при создании запроса: %v\n", err)
		os.Exit(1)
	}

	// Добавляем заголовки
	req.Header.Set("Content-Type", "application/json")

	// Добавляем Basic Auth
	req.SetBasicAuth(*username, *password)

	// Отправляем запрос
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Ошибка при отправке запроса: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Читаем ответ
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Ошибка при чтении ответа: %v\n", err)
		os.Exit(1)
	}

	// Проверяем статус ответа
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Ошибка при отправке SMS. Код: %d, Ответ: %s\n", resp.StatusCode, string(body))
		os.Exit(1)
	}

	fmt.Println("SMS успешно отправлено!")
	fmt.Printf("Ответ сервера: %s\n", string(body))
}
