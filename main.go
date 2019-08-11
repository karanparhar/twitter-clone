package main

import (
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/ethereum_project/config"
	"github.com/ethereum_project/controller"
	"github.com/ethereum_project/repository"
	u "github.com/ethereum_project/usecase"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	conn := repository.GetSession(config.Conf)

	repo := repository.NewBlocksRepository(conn)

	router := mux.NewRouter()

	ethereum_client, err := u.NewClient(u.Trurl)

	if err != nil {

		log.Fatalf("Not able to connect ethereum client", err)

	}

	timeoutContext := time.Duration(30 * time.Second)

	usecase := u.NewtransactionUsecase(repo, timeoutContext, ethereum_client)

	r := controller.NewHandler(router, usecase)

	log.Fatal(http.ListenAndServe(":8090", handlers.LoggingHandler(os.Stdout, r)))

}
