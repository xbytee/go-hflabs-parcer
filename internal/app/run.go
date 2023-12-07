package app

import (
	"GoHFLabsParcer/config"
	"GoHFLabsParcer/internal/usecase"
	"time"

	"github.com/sirupsen/logrus"
)

func Run() {
	config, client, err := config.ReadConfig()
	if err != nil {
		logrus.Fatalf("error parce config: %v", err)
	}

	// Если не использовать cron...
	for {
		logrus.Info("Start Work")
		if err := usecase.CheckTable(config, client); err != nil {
			logrus.Errorf("error during table generation, please check config file or update your service key: %v", err)
		} else {
			logrus.Info("Success Update")
		}

		timer1 := time.NewTimer(time.Duration(config.TimeToRepeat) * time.Minute)
		<-timer1.C
	}
}
