package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"

	"github.com/dblogcorp/dbloggy/config"
	"github.com/dblogcorp/dbloggy/internal/utils"
	"github.com/dblogcorp/dbloggy/internal/utils/log"
	"github.com/dblogcorp/dbloggy/pkg/sso"
	"github.com/dblogcorp/dbloggy/version"
)

var (
	cfgFile string
	timeout = 20

	rootCmd    *cobra.Command
	versionCmd *cobra.Command
)

func NewCommand() *cobra.Command {
	rootCmd := &cobra.Command{
		Use: "run",
		Short: "Start the application.",
		Run: func(cmd *cobra.Command, args []string) {
			// Init config from cfg_file
			config.InitConfig("", &sso.EnvConfig{}, &sso.EnvConfig{})

			r := utils.NewGinEngine()
			// init gin router

			// Register graceful shutdown
			srv := &http.Server{
				Addr: ":8080",
				Handler: r,
			}
			go func(srv *http.Server, timeout int64) {
				registerGracefulShutdown(srv, timeout)
			}(srv, 20)

			// Start
			log.Infof("Server serving on: %s", ":8080")
			if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
				panic(err)
			}
		},
	}
	versionCmd = &cobra.Command{
		Use: "version",
		Short: "Print the version number.",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(version.String())
		},
	}

	rootCmd.AddCommand(versionCmd)
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "f", "./config.yaml", "config file path")
	rootCmd.PersistentFlags().IntVar(&timeout, "graceful-timeout", 20, "timeout for graceful shutdown of server")
	return rootCmd
}

func registerGracefulShutdown(srv *http.Server, timeout int64) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block util a signal is received
	<-quit
	log.Warnf("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeout)*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server shutdown failed, err: %s.", err.Error())
	}
	log.Warnf("Server exited.")
}
