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

// Note: the example only works with the code within the same release/branch.
package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	//
	// Uncomment to load all auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth"
	//
	// Or uncomment to load specific auth plugins
	// _ "k8s.io/client-go/plugin/pkg/client/auth/azure"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/gcp"
	// _ "k8s.io/client-go/plugin/pkg/client/auth/oidc"
)

func main() {
	// rest.InClusterConfig()
	config, err := clientcmd.BuildConfigFromFlags("", "/Users/liyang/.kube/config")
	// creates the in-cluster config
	// config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	// creates the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	namespace := "test"
	deployments, err := clientset.AppsV1().Deployments(namespace).List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("There are %d deployments in the cluster\n", len(deployments.Items))
	for _, deployment := range deployments.Items {
		fmt.Println(deployment.GetName())
	}

	var updateName string = "nginx-1"
	var sc int32 = 0
	updateDeployment, err := clientset.AppsV1().Deployments(namespace).Get(context.TODO(), updateName, metav1.GetOptions{})
	if err != nil {
		panic(err.Error())
	}
	updateDeployment.Spec.Replicas = &sc
	clientset.AppsV1().Deployments(namespace).Update(context.TODO(), updateDeployment, metav1.UpdateOptions{})

}
