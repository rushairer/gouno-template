package gouno

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rushairer/gouno"
	"github.com/spf13/cobra"
	"{{.ModulePath}}/config"
	"{{.ModulePath}}/middleware"
	"{{.ModulePath}}/router"
)

var webCmd = &cobra.Command{
	Use: "web",
	Run: startWebServer,
}

func init() {
	webCmd.Flags().StringP("config", "c", "./config/config.yaml", "config file path")
	webCmd.Flags().StringP("address", "a", "0.0.0.0", "address to listen on")
	webCmd.Flags().StringP("port", "p", "8080", "port to listen on")
	webCmd.Flags().BoolP("debug", "d", false, "debug mode")
}

func startWebServer(cmd *cobra.Command, args []string) {
	log.Printf("starting web server...")

	err := gouno.InitConfig(cmd.Flag("config").Value.String())
	if err != nil {
		log.Fatalf("init config failed, err: %v", err)
	}

	config.GlobalConfig.WebServerConfig.Address = cmd.Flag("address").Value.String()
	config.GlobalConfig.WebServerConfig.Port = cmd.Flag("port").Value.String()
	config.GlobalConfig.WebServerConfig.Debug, _ = cmd.Flags().GetBool("debug")

	if config.GlobalConfig.WebServerConfig.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	engine := gin.New()
	engine.Use(
		middleware.RecoveryMiddleware(),
		middleware.TimeoutMiddleware(),
	)
	router.RegisterWebRouter(engine)

	httpServer := &http.Server{
		Addr:              fmt.Sprintf("%s:%s", config.GlobalConfig.WebServerConfig.Address, config.GlobalConfig.WebServerConfig.Port),
		IdleTimeout:       config.GlobalConfig.WebServerConfig.IdleTimeout,
		WriteTimeout:      config.GlobalConfig.WebServerConfig.WriteTimeout,
		ReadTimeout:       config.GlobalConfig.WebServerConfig.ReadTimeout,
		ReadHeaderTimeout: config.GlobalConfig.WebServerConfig.ReadHeaderTimeout,
		Handler:           engine,
	}

	log.Printf("listening on %s", httpServer.Addr)
	log.Printf("press Ctrl+C to exit")

	go func() {
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Printf("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("server forced to shutdown: %v", err)
	}

	// Close

	log.Printf("server exiting")
}
