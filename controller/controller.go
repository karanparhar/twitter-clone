package controller

import (
	"encoding/json"
	"net/http"

	"github.com/mitchellh/mapstructure"

	"github.com/gorilla/context"

	"github.com/gorilla/mux"
	"github.com/twitter-clone/middleware"
	"github.com/twitter-clone/models"
	u "github.com/twitter-clone/usecase"
)

type transactionhandler struct {
	user u.Usecase
}

func NewHandler(r *mux.Router, uc u.Usecase) *mux.Router {

	router := r.NewRoute().Subrouter()
	tw := router.NewRoute().Subrouter()
	tw.Use(middleware.ValidateMiddleware)

	t := &transactionhandler{
		uc,
	}
	router.HandleFunc("/signup", t.SignUp).Methods("POST", "OPTIONS")
	router.HandleFunc("/login", t.Login).Methods("POST", "OPTIONS")
	tw.HandleFunc("/newtweet", t.NewTweet).Methods("POST", "OPTIONS")
	tw.HandleFunc("/followuser", t.FollowPeoples).Methods("POST", "OPTIONS")
	tw.HandleFunc("/getfollowerstweets", t.GetFollowersTweets).Methods("GET", "OPTIONS")
	//this route will be used for instancegroup healthcheck
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("This is a catch-all route")
	})

	return router

}

func (t *transactionhandler) SignUp(w http.ResponseWriter, req *http.Request) {
	var user models.Profile
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err := t.user.SignUp(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")

}
func (t *transactionhandler) Login(w http.ResponseWriter, req *http.Request) {
	var user models.Profile
	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&user); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	token, err := t.user.Login(user)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)

}

func (t *transactionhandler) NewTweet(w http.ResponseWriter, req *http.Request) {
	userdata := context.Get(req, "decoded")

	var username models.Profile
	err := mapstructure.Decode(userdata, &username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var tweet models.Tweet

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&tweet); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	tweet.Username = username.Username

	err = t.user.InsertTweets(tweet)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")

}
func (t *transactionhandler) FollowPeoples(w http.ResponseWriter, req *http.Request) {
	userdata := context.Get(req, "decoded")

	var username models.Profile
	err := mapstructure.Decode(userdata, &username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	var followerid models.Follower

	defer req.Body.Close()
	if err := json.NewDecoder(req.Body).Decode(&followerid); err != nil {

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	err = t.user.FollowUser(username, followerid.ID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return

	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("success")

}

func (t *transactionhandler) GetFollowersTweets(w http.ResponseWriter, req *http.Request) {
	userdata := context.Get(req, "decoded")

	var username models.Profile
	err := mapstructure.Decode(userdata, &username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	res, err := t.user.GetFollowersTweets(username)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)

}
