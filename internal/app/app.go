package app

import (
	"context"
	"go.uber.org/zap"
	"sego/config"
	"sego/internal/domain"
	"sego/internal/repository"
	"sego/internal/ui"
	"sego/pkg/logger"
)

func Run(config *config.Config) {
	log := logger.NewZapLogger()
	if zapLogger, ok := log.(*logger.ZapLogger); ok {
		// no need to handle Sync error https://github.com/uber-go/zap/issues/328
		defer func() {
			_ = zapLogger.Sync()
		}()
	}
	log.Info("initialized logger")

	log.Debug("app config: ", zap.Any("config", config))

	log.Info("connecting to database")
	db, err := repository.New()
	if err != nil {
		log.Fatal("failed to connect to database: ", err)
	}
	defer func() {
		log.Info("closing database connection")
		err := db.Close(context.Background())
		if err == nil {
			log.Info("database connection closed")
		} else {
			log.Error("failed to close database connection: ", err)
		}
	}()
	log.Info("connected to database")

	r := repository.NewRepository(db)
	d := domain.NewDomain(r)
	u := ui.NewUI(d)
	u.Start()
}
