package logger

import "go.uber.org/zap"

func New(env string) (*zap.Logger, error) {
	if env == "prod" || env == "release" {
		return zap.NewProduction()
	}
	cfg := zap.NewDevelopmentConfig()
	return cfg.Build()
}
