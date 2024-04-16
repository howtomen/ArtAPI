package logger

import (
	"io"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/pkgerrors"
	"gopkg.in/natefinch/lumberjack.v2"
)
var once sync.Once
var log zerolog.Logger

func GetLogger() zerolog.Logger {
	once.Do(func () {
		zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
		zerolog.TimeFieldFormat = time.RFC3339Nano

		logLevel, err := strconv.Atoi(os.Getenv("LOG_LEVEL"))
		if err != nil {
			logLevel = int(zerolog.InfoLevel)
		}
		var output io.Writer = zerolog.ConsoleWriter{
			Out: os.Stdout,
			TimeFormat: time.RFC3339,
		}

		if os.Getenv("APP_ENV") != "dev" { //store prod logs 
			fileLogger := &lumberjack.Logger{
				Filename: "art-api.log",
				MaxSize: 5, //max MB size of log before it gets rotated
				MaxBackups: 10,
				MaxAge: 14,
				Compress: true,
			}
			output = zerolog.MultiLevelWriter(os.Stderr, fileLogger)
		}
		
		log = zerolog.New(output).
			Level(zerolog.Level(logLevel)).
			With().
			Timestamp().
			Logger()
	})

	return log
}