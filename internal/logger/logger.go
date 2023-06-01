package logger

import (
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog"
)

func New() *zerolog.Logger {
	dateString := time.Now().Format("20060102")

	file, err := os.OpenFile(
		fmt.Sprintf("logs/%s.log", dateString),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644)
	if err != nil {
		panic(fmt.Errorf("failed to open log file: %w", err))
	}

	multiLogger := zerolog.MultiLevelWriter(os.Stdout, file)
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	logger := zerolog.New(multiLogger).With().Timestamp().Logger()
	return &logger
}
