package main

import (
	"flag"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/serialt/sugar/v2"
)

func init() {
	flag.BoolVar(&appVersion, "v", false, "Display build and version messages")
	flag.StringVar(&ConfigFile, "c", "config.yaml", "Config file")
	flag.Parse()

	err := sugar.LoadConfig(ConfigFile, &config)
	if err != nil {
		config = new(Config)
	}
	slog.SetDefault(sugar.New(
		sugar.WithLevel(config.Log.Level),
		sugar.WithFile(config.Log.File),
	),
	)

}
func main() {
	if appVersion {
		fmt.Printf("APPVersion: %v  BuildTime: %v  GitCommit: %v\n",
			APPVersion,
			BuildTime,
			GitCommit)
		return
	}

	service()

	// 进程持续运行
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	s := <-c
	slog.Info("Aborting...", "signal", s)
	os.Exit(2)
}
