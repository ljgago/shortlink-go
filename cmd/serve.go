package cmd

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/spf13/cobra"

	"github.com/ljgago/shortlink-go/pkg/log"
	"github.com/ljgago/shortlink-go/ui"
)

// serveCmd represents the server command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Starts the application server",
	Long: `serve is a command that starts the main application server,
handling HTTP requests, API endpoints, and background tasks.`,
	Run: func(cmd *cobra.Command, args []string) {
		serveRun(cmd, args)
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.Flags().StringP("port", "p", "8080", "port to run the server on")
}

func serveRun(cmd *cobra.Command, _ []string) {
	port, err := cmd.Flags().GetString("port")
	if err != nil {
		log.Errorf("Error getting port flag: %s\n", err)
		return
	}

	// define all routes
	r := chi.NewRouter()
	r.Use(middleware.Recoverer)
	ui.RegisterRoutes(r)

	// Run the server
	log.Info("Server start at :" + port)
	server := &http.Server{
		Addr:        ":" + port,
		ReadTimeout: 1 * time.Second,
		Handler:     r,
	}
	if err := server.ListenAndServe(); err != nil {
		log.Errorf("Server failed: %s\n", err)
	}
}
