package config

import (
	"os"
)

// Config はアプリケーション設定情報を管理する構造体例
type Config struct {
	AppName   string
	OutputDir string
}

// LoadConfig は環境変数やファイルから設定をロードする（簡易版）
func LoadConfig() *Config {
	appName := os.Getenv("APP_NAME")
	if appName == "" {
		appName = "MyApp"
	}
	outputDir := os.Getenv("OUTPUT_DIR")
	if outputDir == "" {
		outputDir = "./output"
	}
	return &Config{
		AppName:   appName,
		OutputDir: outputDir,
	}
}
