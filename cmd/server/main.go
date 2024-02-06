package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"runtime/debug"

	"github.com/makrorof/gorm-graphql-postgres-echo-auth/configs"
	"github.com/makrorof/gorm-graphql-postgres-echo-auth/internal/app/server"
	"github.com/makrorof/gorm-graphql-postgres-echo-auth/tools"
	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"

	nested "github.com/antonfisher/nested-logrus-formatter"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln(string(debug.Stack()))
		}
	}()

	if cfg, err := configs.ReadConf("config.yaml"); err != nil {
		tools.FailOnError(err, "Config file not read")
		return
	} else {
		setupLog(cfg)
		server.Run(cfg)
	}
}

func setupLog(cfg *configs.Config) {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&nested.Formatter{
		HideKeys:    true,
		FieldsOrder: []string{"component", "category"},
	})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	//logrus.SetOutput(os.Stdout)
	ljack := &lumberjack.Logger{
		Filename:   cfg.LogConf.Path + "/service.log",
		MaxSize:    500, // megabytes
		MaxBackups: 0,
		MaxAge:     30 * 6, //days 6 ay
		Compress:   true,   // disabled by default
	}

	mWriter := io.MultiWriter(os.Stderr, ljack)
	logrus.SetOutput(mWriter)

	// Only log the warning severity or above.
	level := logrus.Level(cfg.LogConf.Level)
	logrus.SetLevel(level)

	fmt.Println("Log level: " + level.String())
}
