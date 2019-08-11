# ethereum_project

This is sample ethereum_project

## Quick start

### Prerequisites
- [go](https://golang.org/dl/) version v1.10+

### Steps to run

```
$ mkdir $GOPATH/src/github.com/
$ cd $GOPATH/src/github.com/
$ git clone https://karanjit@bitbucket.org/karanjit/ethereum_project.git
$ cd ethereum_project
$ make docker
```

## check status
```
$ curl -X GET http://localhost:8090/
```
## get blocks
```
$ curl -X GET http://localhost:8090/fetchblocks
```
## get transactions
```
$ curl -X GET http://localhost:8090/gettransaction
```



### Note it will build rest service and crd and deploy in local kubernetes cluster 

### if role binding issue came kubectl create clusterrolebinding serviceaccounts-cluster-admin --clusterrole=cluster-admin --group=system:serviceaccounts
 
 