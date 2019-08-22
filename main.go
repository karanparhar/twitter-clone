package main

import (
	"context"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/twitter-clone/config"
	"github.com/twitter-clone/controller"

	//"github.com/twitter-clone/middleware"
	tweetsrepository "github.com/twitter-clone/tweets/repository"
	usecase "github.com/twitter-clone/usecase/users"
	userrepo "github.com/twitter-clone/user-management/repository"
)

func main() {

	conn := config.GetSession(config.Conf)

	repo := userrepo.NewUsersRepo(conn)
	tweetsrepo := tweetsrepository.NewTweetsRepo(conn)

	router := mux.NewRouter()

	Context, _ := context.WithTimeout(context.Background(), time.Duration(5*time.Second))

	u := usecase.NewtransactionUsecase(Context, repo, tweetsrepo)

	r := controller.NewHandler(router, u)

	log.Fatal(http.ListenAndServe(":8090", handlers.LoggingHandler(os.Stdout, r)))

}
