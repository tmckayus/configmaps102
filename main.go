/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

/* Copyright 2017 Trevor McKay */

// Note: the example only works with the code within the same release/branch.
package main

import (
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"path/filepath"
	"io/ioutil"
	"time"
)

func authStuff() *kubernetes.Clientset {
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	return clientset
}

func writeConfigMap(clientset *kubernetes.Clientset, name, directory, namespace string) {
	cmap, err := clientset.CoreV1().ConfigMaps(namespace).Get(name, metav1.GetOptions{})
	if err == nil {
		fmt.Printf("Reading configmap %s\n", name)
		for k, v := range cmap.Data {
			path := filepath.Join(directory, k)
			fmt.Printf("Writing %s to %s\n", v, path)
			file, _ := os.Create(path)
			file.WriteString(v)
		}
		return
        }
        // else {
        //              panic(err.Error())
        // }
	fmt.Printf("Sorry, didn't find configmap %s\n", name)
}

func main() {

	clientset := authStuff()

	name := os.Getenv("configname")
	directory := filepath.Join("/tmp", name)
	bytes, err := ioutil.ReadFile("/var/run/secrets/kubernetes.io/serviceaccount/namespace")
	if err != nil {
		panic(err.Error())
	}
	namespace := string(bytes)


	for {
		os.RemoveAll(directory)
		os.MkdirAll(directory, 0755)
		writeConfigMap(clientset, name, directory, namespace)
		fmt.Println("Sleeping...")
		time.Sleep(5 * time.Second)
	}
}
