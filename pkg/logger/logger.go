package logger

import (
	"flag"
	"github.com/rs/zerolog"
	"os"
)

var (
	Log zerolog.Logger
)

func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	// set location of log file
	var logpath = "monitor.log"

	flag.Parse()
	var file, err = os.Create(logpath)

	if err != nil {
		panic(err)
	}

	Log = zerolog.New(file)
	Log.Info().Msg("LogFile : " + logpath)
}
