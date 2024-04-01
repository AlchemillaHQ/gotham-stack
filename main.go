package main

import (
	"fmt"
	"github.com/AlchemillaHQ/gotham-stack/db"
	"github.com/AlchemillaHQ/gotham-stack/internal/config"
	"github.com/AlchemillaHQ/gotham-stack/internal/handlers"
	"github.com/AlchemillaHQ/gotham-stack/internal/logger"
	"github.com/AlchemillaHQ/gotham-stack/internal/middleware"
	"github.com/AlchemillaHQ/gotham-stack/internal/services"
	"go.uber.org/zap"
	"html/template"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlchemillaHQ/gotham-stack/cmd"
)

func main() {
	cmd.AsciiArt()
	cfgPath := cmd.ParseFlags()
	cfg := config.ParseConfig(cfgPath)

	fmt.Println(cfg.Environment)

	logger.InitializeLogger(cfg.Environment, cfg.LogLevel)
	d := db.SetupDatabase(cfg.Environment, cfg.DatabaseURL)
	db.SetupBasicData(d, cfg.Admins)
	middleware.InitializeSessionStore(cfg.SessionSecret)

	serviceRegistry := services.NewServiceRegistry(d)

	mux := http.NewServeMux()

	fsStatic := http.FileServer(http.Dir("./public"))
	templ, err := template.New("").ParseGlob("./templates/*.html")

	_, err = templ.ParseGlob("./templates/partials/*.html")
	_, err = templ.ParseGlob("./templates/auth/*.html")

	if err != nil {
		logger.Fatal("Failed to parse templates", zap.Error(err))
	}

	mux.Handle("GET /{$}", middleware.EnsureAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handlers.IndexHandler(w, r, templ) })))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fsStatic))
	mux.Handle("/auth/{type}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AuthHandler(w, r, templ, serviceRegistry.AuthService)
	}))

	mux.Handle("/counter/{type}", middleware.EnsureAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.CounterHandler(w, r, serviceRegistry.CounterService)
	})))

	go func() {
		logger.Info("Starting server", zap.Int("port", cfg.Port))
		if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.Port), mux); err != nil {
			logger.Fatal("Failed to start server", zap.Error(err))
		}
	}()

	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)

	<-stopChan
	logger.Info("Shutting down server...")
}
