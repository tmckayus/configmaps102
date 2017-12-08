# configmaps102
Source examples from KubeCon Austin 2017 ConfigMaps102 talk

This example is based on the examples in https://github.com/kubernetes/client-go
Make sure to look through that repository for context on client-go

You can build like this although you're on your own for downloading the deps :)

$ go build -o app .

$ docker build -t testapp-loop .

You can launch like this. Note that testapp.yaml is set up for referencing the image
in the local docker daemon on OpenShift Origin or minikube. If you want to pull the
image from docker hub or another registry, modify accordingly.

$ kubectl create -f testapp.yaml

Also note that if things don't seem to be working after you've created a configmap
named *bob* you can add/uncomment panic(err.Error()) lines in main.go to see what's happening.

Enjoy!
