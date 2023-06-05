package auth

import (
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	"log"
)



// https://pkg.go.dev/k8s.io/client-go@v0.27.2/tools/clientcmd#BuildConfigFromFlags
func GetKubeconfig(kubeconfig string) *rest.Config {
	config, err := clientcmd.BuildConfigFromFlags("", kubeconfig)
	if err != nil {
		log.Fatal(err)
	}
	return config
}
