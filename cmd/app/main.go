package main

import (
	"github.com/Olegsuus/MoviesRest/internal/app"
	"github.com/Olegsuus/MoviesRest/internal/config"
	"log"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Ошибка загрузки конфигурации: %v", err)
	}

	App, err := app.NewApp(cfg)
	if err != nil {
		log.Fatalf("Ошибка инициализации приложения: %v", err)
	}

	if err = App.Run(); err != nil {
		log.Fatalf("Ошибка при работе приложения: %v", err)
	}
}
