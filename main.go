package main

import (
	"context"
	"fmt"
	"log"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

func main() {
	//kubeconfig := flag.String("kubeconfig", "/home/fadia/.kube/config", "location to your config file")
	//config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	//config, err := clientcmd.BuildConfigFromFlags("", "/home//fadia/.kube/config")
	//config, err := clientcmd.BuildConfigFromFlags("", "./config")
	//if err != nil {
	//panic(err)
	//}
	config, err := rest.InClusterConfig()
	if err != nil {
		log.Fatal(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	nodes, err := clientset.CoreV1().Nodes().List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, node := range nodes.Items {
		annotations := node.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}

		annotations["new_value"] = "it_worked"

		node.SetAnnotations(annotations)

		_, err := clientset.CoreV1().Nodes().Update(context.TODO(), &node, metav1.UpdateOptions{})

		if err != nil {
			fmt.Printf("Error updating node %s: %v\n", node.Name, err)
		} else {
			fmt.Printf("Successfully annotated node %s with key=value, yayy\n", node.Name)
		}
	}
}
