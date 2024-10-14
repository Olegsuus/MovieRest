package app

import (
	"MovieRest/internal/config"
	handlers "MovieRest/internal/handlers/movies"
	"MovieRest/internal/handlers/routers"
	services "MovieRest/internal/services/movies"
	grpc "MovieRest/pkg/grpc/movies"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg         *config.Config
	logger      *slog.Logger
	httpServer  *gin.Engine
	movieClient *grpc.MovieClient
}

func NewApp(cfg *config.Config) (*App, error) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	movieClient, err := grpc.NewMovieClient(cfg.GRPC.Address)
	if err != nil {
		logger.Error("Ошибка при создании gRPC клиента", slog.Any("error", err))
		return nil, fmt.Errorf("не удалось создать gRPC клиент: %w", err)
	}

	movieRepository := grpc.NewMovieRepository(movieClient)

	movieService := services.NewMoviesService(movieRepository, logger)

	movieHandler := handlers.NewMovieHandlers(movieService)

	router := gin.Default()

	routers.SetupRouter(router, movieHandler)

	return &App{
		cfg:         cfg,
		logger:      logger,
		httpServer:  router,
		movieClient: movieClient,
	}, nil
}

func (a *App) Run() error {
	port := a.cfg.App.Server.Port
	addr := fmt.Sprintf(":%d", port)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		if err := a.httpServer.Run(addr); err != nil {
			a.logger.Error("Ошибка при запуске HTTP-сервера", slog.Any("error", err))
			os.Exit(1)
		}
	}()

	a.logger.Info("HTTP-сервер запущен", slog.String("address", addr))

	<-stop

	a.logger.Info("Получен сигнал завершения, останавливаем приложение")

	if err := a.movieClient.Close(); err != nil {
		a.logger.Error("Ошибка при закрытии gRPC клиента", slog.Any("error", err))
	}

	return nil
}
