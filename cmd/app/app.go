package app

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ffajarpratama/gommerce-api/config"
	"github.com/ffajarpratama/gommerce-api/internal/http/handler"
	"github.com/ffajarpratama/gommerce-api/internal/repository"
	"github.com/ffajarpratama/gommerce-api/internal/usecase"
	"github.com/ffajarpratama/gommerce-api/lib/mysql"
)

func Exec() error {
	cnf := config.New()

	mysqlClient, err := mysql.NewMySQLClient(cnf)
	if err != nil {
		return err
	}

	repo := repository.New(mysqlClient)
	uc := usecase.New(cnf, repo, mysqlClient)
	r := handler.NewHTTPRouter(cnf, uc)

	addr := flag.String("http", fmt.Sprintf(":%d", cnf.App.Port), "HTTP listen address")
	httpServer := &http.Server{
		Addr:              *addr,
		Handler:           r,
		ReadHeaderTimeout: 90 * time.Second,
	}

	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		<-sig

		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Printf("[graceful-shutdown-time-out] \n%v\n", err.Error())
			}
		}()

		defer cancel()

		log.Println("graceful shutdown.....")

		err = httpServer.Shutdown(shutdownCtx)
		if err != nil {
			log.Printf("[graceful-shutdown-error] \n%v\n", err.Error())
		}

		serverStopCtx()
	}()

	log.Printf("HTTP server running on %s\n", *addr)

	err = httpServer.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Printf("[http-server-failed] \n%v\n", err.Error())
		return err
	}

	<-serverCtx.Done()

	return nil
}
