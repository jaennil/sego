package main

import (
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "log"
)

var logger *zap.SugaredLogger

func init() {
    c := zap.NewDevelopmentConfig()
    c.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder

    l, err := c.Build()
    if err != nil {
        log.Fatal(err)
    }

    logger = l.Sugar()
}
