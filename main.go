package main

import (
	"embed"
	"fmt"
	"github.com/AlchemillaHQ/gotham-stack/db"
	"github.com/AlchemillaHQ/gotham-stack/internal/config"
	"github.com/AlchemillaHQ/gotham-stack/internal/handlers"
	"github.com/AlchemillaHQ/gotham-stack/internal/i18n"
	"github.com/AlchemillaHQ/gotham-stack/internal/logger"
	"github.com/AlchemillaHQ/gotham-stack/internal/middleware"
	"github.com/AlchemillaHQ/gotham-stack/internal/services"
	"go.uber.org/zap"
	"html/template"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlchemillaHQ/gotham-stack/cmd"
)

//go:embed templates/*
var templatesFS embed.FS

//go:embed public/*
var public embed.FS
var publicFS, _ = fs.Sub(public, "public")

func main() {
	cmd.AsciiArt()
	cfgPath := cmd.ParseFlags()
	cfg := config.ParseConfig(cfgPath)

	logger.InitializeLogger(cfg.Environment, cfg.LogLevel, cfg.DataPath)
	d := db.SetupDatabase(cfg.Environment, cfg.DatabaseURL, cfg.DataPath)
	db.SetupBasicData(d, cfg.Admins)
	middleware.InitializeSessionStore(cfg.SessionSecret)

	serviceRegistry := services.NewServiceRegistry(d)

	mux := http.NewServeMux()

	templ := template.New("").Funcs(template.FuncMap{
		"translate": func(page, key, lang string) string {
			return i18n.GetTranslation(page, key, lang)
		},
	})

	templ, err := templ.ParseFS(templatesFS, "templates/*.html", "templates/partials/*.html", "templates/auth/*.html")
	if err != nil {
		log.Fatal("Failed to parse templates", zap.Error(err))
	}

	mux.Handle("GET /{$}", middleware.EnsureAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { handlers.IndexHandler(w, r, templ) })))
	mux.Handle("GET /static/", http.StripPrefix("/static/", http.FileServer(http.FS(publicFS))))
	mux.Handle("/auth/{type}", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.AuthHandler(w, r, templ, serviceRegistry.AuthService)
	}))

	mux.Handle("/task/", middleware.EnsureAuthenticated(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handlers.TaskHandler(w, r, serviceRegistry.TaskService)
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
