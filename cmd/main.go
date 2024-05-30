package main

import (
	"diplomServer/internal/app"
	"diplomServer/internal/config"
	"diplomServer/internal/lib/logger"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	cfg, err := config.LoadConfigFromYaml()
	if err != nil {
		log.Fatal("Error read config: ", err.Error())
	}

	lg, err := logger.NewLogger()
	if err != nil {
		log.Fatal("Can't initialize logger: ", err.Error())
	}

	application := app.New(lg, cfg)

	go application.GRPSrv.MustRun()

	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	signal := <-stop

	lg.Info(fmt.Sprintf("application stopped signal: %s", signal.String()))

	application.GRPSrv.Stop()

	lg.Info("application stop")

}
