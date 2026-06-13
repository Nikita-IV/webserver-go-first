package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Описываем структуру игрока
type Player struct {
	Name      string   `json:"name"`      // Эти теги нужны, чтобы Go красиво
	Coins     int      `json:"coins"`     // превращал данные в формат JSON
	Inventory []string `json:"inventory"` // который понимают все сайты в мире
}

func main() {
	// Создаем нашего героя
	hero := Player{
		Name:      "Воин_Никитос",
		Coins:     1340,
		Inventory: []string{"Меч", "Щит", "Зелье", "Кольцо", "Лук"},
	}

	fmt.Println("🚀 Веб-сервер запускается на порту :8080...")
	fmt.Println("👉 Открой браузер и перейди по ссылке: http://localhost:8080/profile")

	// Говорим серверу: "Если кто-то перейдет по адресу /profile — запусти функцию ниже"
	http.HandleFunc("/profile", func(w http.ResponseWriter, r *http.Request) {
		// Устанавливаем заголовки, чтобы браузер понял, что мы передаем данные, а не просто текст
		w.Header().Set("Content-Type", "application/json")

		// Кодируем структуру нашего героя в формат JSON и отправляем прямо в браузер
		json.NewEncoder(w).Encode(hero)
	})

	// Запускаем бесконечное прослушивание порта 8080
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
