# twitter-clone

This is sample twitter-clone

## Quick start

### Prerequisites
- [go](https://golang.org/dl/) version v1.10+

### Steps to run

```
$ mkdir $GOPATH/src/github.com/
$ cd $GOPATH/src/github.com/
$ git clone https://karanjit@bitbucket.org/karanjit/twitter-clone.git
$ cd twitter-clone
$ make docker
$ make deploy
```

## check status
```
$ curl -X GET $(minikube service minitwitter-testing --url)/

```
## signup new user
```
$ curl -X POST $(minikube service minitwitter-testing --url)/signup -d '{ "username":"test", "email":"test@gmail.com", "password":"test0099" }'
```

## login api
```
$ curl -X POST $(minikube service minitwitter-testing --url)/login -d '{"username":"test","password":"test0099"}'
```

## new tweet
```
$ curl -H "Authorization: $token" -X POST $(minikube service minitwitter-testing --url)/newtweet -d { "username":"test", "text":"This is my first tweet" }
```

## get followers tweets
```
$ curl -H "Authorization: $token" -X GET $(minikube service minitwitter-testing --url)/getfollowerstweets
```
## follow peoples
```
$ curl -H "Authorization: $token" -X POST $(minikube service minitwitter-testing --url)/followuser -d { "id":"test1"}
```





