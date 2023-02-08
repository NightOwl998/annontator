package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

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
	value := SendRequestAPI()
	for _, node := range nodes.Items {
		//we have one link so one bandwidth quality we can use in both cases, whereasœœ

		annotations := node.GetAnnotations()
		if annotations == nil {
			annotations = make(map[string]string)
		}

		annotations["Connexion_Quality"] = value

		node.SetAnnotations(annotations)

		_, err := clientset.CoreV1().Nodes().Update(context.TODO(), &node, metav1.UpdateOptions{})

		if err != nil {
			fmt.Printf("Error updating node %s: %v\n", node.Name, err)
		} else {
			fmt.Printf("Successfully annotated node %s with connexion value, yayy\n", node.Name)
		}
	}
}
func SendRequestAPI() string {
	resp, err := http.Get("http://localhost:9090/wifi/isEnabled")
	if err != nil {
		log.Fatalln(err)
	}
	//We Read the response body on the line below.
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	//Convert the body to type string
	sb := string(body)
	if sb == "true" {
		resp, err := http.Get("http://localhost:9090/wifi/quality")
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		//Convert the body to type string
		sb := string(body)
		return sb
	} else {
		resp, err := http.Get("http://localhost:9090/lte/quality")
		if err != nil {
			log.Fatalln(err)
		}
		//We Read the response body on the line below.
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		//Convert the body to type string
		sb := string(body)
		return sb

	}

	//log.Printf(sb)

}
