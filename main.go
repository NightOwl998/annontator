package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func main() {
	config, err := clientcmd.BuildConfigFromFlags("", "/home/fadia/.kube/config")
	if err != nil {
		panic(err)
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

		annotations["key"] = "value"

		node.SetAnnotations(annotations)

		_, err := clientset.CoreV1().Nodes().Update(context.TODO(), &node, metav1.UpdateOptions{})

		if err != nil {
			fmt.Printf("Error updating node %s: %v\n", node.Name, err)
		} else {
			fmt.Printf("Successfully annotated node %s with key=value\n", node.Name)
		}
	}
}
