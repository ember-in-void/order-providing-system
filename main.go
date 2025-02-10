package main

import (
	"log"
	"net/http"

	"frappuccino/internal/service/usecase"
	pkgDB "frappuccino/internal/storage"
	"frappuccino/internal/storage/repo"
	"frappuccino/internal/transport"
	"frappuccino/internal/transport/handler"
	"frappuccino/pkg/logger"
)

func main() {
	// initialize logger
	logg, err := logger.NewCustomLogger()
	if err != nil {
		log.Fatalln("failed to init logger")
	}
	logg.Info("Logger initialized successfully")

	// initialize database
	db := pkgDB.SetupDataBase(logg)
	defer db.Close()

	// initialize repository, usecase and handler
	repository := repo.NewRepository(db)
	usecase := usecase.NewUsecase(repository)
	handlers := handler.NewHttpHandler(usecase, logg)

	// initialize router
	router := transport.SetupRouter(handlers)

	// start server
	log.Println("Starting server on port 8080")
	logg.Info("Starting server on port 8080")
	err = http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatalln(err)
		logg.Error(err)
	}
}
