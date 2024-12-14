package main

import (
	"flag"
	"log"

	"github.com/Amirsport/web-10/internal/count/api"
	"github.com/Amirsport/web-10/internal/count/config"
	"github.com/Amirsport/web-10/internal/count/provider"
	"github.com/Amirsport/web-10/internal/count/usecase"

	_ "github.com/lib/pq"
)

func main() {
	configPath := flag.String("config-path", "../../configs/count_example.yaml", "путь к файлу конфигурации")
	flag.Parse()

	cfg, err := config.LoadConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}

	prv := provider.NewProvider(cfg.DB.Host, cfg.DB.Port, cfg.DB.User, cfg.DB.Password, cfg.DB.DBname)
	use := usecase.NewUsecase(prv)
	srv := api.NewServer(cfg.IP, cfg.Port, use)

	log.Printf("Сервер запущен на %s\n", srv.Address)
	srv.Run()
}
