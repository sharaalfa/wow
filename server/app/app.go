package app

import (
	"log"
	"net"
	"time"
	"wow/pkg/config"
	"wow/pkg/httpclient"
	"wow/server/internal/services/handler"
	"wow/server/internal/services/quote"
)

type App struct {
	listener net.Listener
	handler  *handler.ConnectionHandler
}

func NewApp() (*App, error) {
	config.InitConfig()
	cfg := config.GetConfig()
	listener, err := net.Listen("tcp", cfg.Server.Host+":"+cfg.Server.Port)
	if err != nil {
		return nil, err
	}

	httpClient := &httpclient.RealHTTPClient{}
	quoteService := quote.NewService(httpClient, cfg.Quotes.EntryPoint, cfg.Quotes.MaxRetries)
	connectionHandler := handler.NewConnectionHandler(cfg.Difficulty, time.Duration(cfg.Timeout)*time.Second, quoteService)
	return &App{listener: listener, handler: connectionHandler}, nil
}

func (a *App) Run() {
	log.Println("Server started on", a.listener.Addr())

	for {
		conn, err := a.listener.Accept()
		if err != nil {
			log.Println("Failed to accept connection: ", err.Error())
			continue
		}
		go a.handler.HandleConnection(conn)
	}
}

func (a *App) Close() error {
	return a.listener.Close()
}
