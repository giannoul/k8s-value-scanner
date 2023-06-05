package scanner

import (
	"context"
	"fmt"
	"github.com/giannoul/k8s-value-scanner/internal/auth"
	"k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	kubernetesAppsV1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	"github.com/olekukonko/tablewriter"
	"log"
	"sync"
	"strings"
	"os"
	"github.com/briandowns/spinner"
	"time"
)

var k8sClient = kubernetes.Clientset{}

func initializeKubeconfig(kubeconfig string){
	// https://pkg.go.dev/k8s.io/client-go/kubernetes#NewForConfig
	config := auth.GetKubeconfig(kubeconfig)
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		log.Fatal(err)
	}
	k8sClient = *clientset
}

func Scan(kubeconfig string, needle string) {
	initializeKubeconfig(kubeconfig)
	s := spinner.New(spinner.CharSets[35], 200*time.Millisecond)
	s.Start() 
	allItems := getItems() 
	s.Stop()

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Kind", "Name", "Namespace", "Key", "Value"})
	table.SetRowLine(true)
	table.SetRowSeparator("-")
	var found = false
	for _,item := range allItems {
		if strings.Contains(item.value, needle) {
			found = true
			table.Append([]string{item.kind, item.name, item.namespace, item.key, item.value})
		} 
	}
	if found {
		table.Render()
	}else{
		fmt.Println("No variable found with the given value.")
	}
}

func getItems() []Item {
	res := []Item{}
	namespaces := getNameSpaces()
	for _, n := range namespaces {
		res = append(res, getDeploymentsStatefulsetsDaemonsetsItems(n)...)
		res = append(res, getPodsItems(n)...)
		res = append(res, getConfigmapItems(n)...)
		res = append(res, getSecretItems(n)...)
	}
	//fmt.Printf("%v", res)
	return res
}

func getNameSpaces() []string {
	var l = []string{}
	nsList, err := k8sClient.CoreV1().Namespaces().List(context.Background(), v1.ListOptions{})
	if err != nil {
		log.Fatal(err)
	}

	for _, n := range nsList.Items {
		l = append(l, n.Name)
	}

	return l
}

func getPodsItems(namespace string) []Item {
	res := []Item{}
	list, err := k8sClient.CoreV1().Pods(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, l := range list.Items {
		temp := []Item{}
		obj, err := k8sClient.CoreV1().Pods(namespace).Get(context.TODO(), l.ObjectMeta.Name, v1.GetOptions{})
		if err != nil {
			panic(err)
		}
		getContainerEnvVars(obj.Spec.Containers, &temp, "Pod", l.ObjectMeta.Name, namespace)
		res = append(res, temp...)
		temp = []Item{}
	}
	return res
}

func getConfigmapItems(namespace string) []Item {
	res := []Item{}
	list, err := k8sClient.CoreV1().ConfigMaps(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, l := range list.Items {
		temp := []Item{}
		obj, err := k8sClient.CoreV1().ConfigMaps(namespace).Get(context.TODO(), l.ObjectMeta.Name, v1.GetOptions{})
		if err != nil {
			panic(err)
		}
		getConfigmapContents(obj.Data, &temp, "Configmap", l.ObjectMeta.Name, namespace)
		res = append(res, temp...)
	}
	return res
}

func getSecretItems(namespace string) []Item {
	res := []Item{}
	list, err := k8sClient.CoreV1().Secrets(namespace).List(context.TODO(), v1.ListOptions{})
	if err != nil {
		panic(err)
	}
	for _, l := range list.Items {
		temp := []Item{}
		obj, err := k8sClient.CoreV1().Secrets(namespace).Get(context.TODO(), l.ObjectMeta.Name, v1.GetOptions{})
		if err != nil {
			panic(err)
		}
		getSecretContents(obj.Data, &temp, "Secret", l.ObjectMeta.Name, namespace)
		res = append(res, temp...)
	}
	return res
}

func getDeploymentsStatefulsetsDaemonsetsItems(namespace string) []Item {
	res := []Item{}
	type Temp struct {
		kind   string
		getter interface{}
		items  []Item
	}

	workloadKinds := []Temp{
		{
			kind:   "Deployment",
			getter: k8sClient.AppsV1().Deployments(namespace),
			items:  []Item{},
		},
		{
			kind:   "Statefulset",
			getter: k8sClient.AppsV1().StatefulSets(namespace),
			items:  []Item{},
		},
		{
			kind:   "Daemonset",
			getter: k8sClient.AppsV1().DaemonSets(namespace),
			items:  []Item{},
		},
	}

	var wg sync.WaitGroup
	for i, _ := range workloadKinds {
		wg.Add(1)
		go func(w *Temp) {
			defer wg.Done()
			// https://go.dev/doc/effective_go#interface_conversions
			// https://go.dev/doc/effective_go#type_switch
			switch cli := w.getter.(type) {
			case kubernetesAppsV1.DeploymentInterface:
				list, err := cli.List(context.TODO(), v1.ListOptions{})
				if err != nil {
					panic(err)
				}
				for _, l := range list.Items {
					obj, err := cli.Get(context.TODO(), l.ObjectMeta.Name, v1.GetOptions{})
					if err != nil {
						panic(err)
					}
					getContainerEnvVars(obj.Spec.Template.Spec.Containers, &w.items, w.kind, l.ObjectMeta.Name, namespace)
				}
			case kubernetesAppsV1.StatefulSetInterface:
				list, err := cli.List(context.TODO(), v1.ListOptions{})
				if err != nil {
					panic(err)
				}
				for _, l := range list.Items {
					obj, err := cli.Get(context.TODO(), l.ObjectMeta.Name, v1.GetOptions{})
					if err != nil {
						panic(err)
					}
					getContainerEnvVars(obj.Spec.Template.Spec.Containers, &w.items, w.kind, l.ObjectMeta.Name, namespace)
				}
			case kubernetesAppsV1.DaemonSetInterface:
				list, err := cli.List(context.TODO(), v1.ListOptions{})
				if err != nil {
					panic(err)
				}
				for _, l := range list.Items {
					obj, err := cli.Get(context.TODO(), l.ObjectMeta.Name, v1.GetOptions{})
					if err != nil {
						panic(err)
					}
					getContainerEnvVars(obj.Spec.Template.Spec.Containers, &w.items, w.kind, l.ObjectMeta.Name, namespace)
				}
			default:
				panic("No https://pkg.go.dev/k8s.io/client-go@v0.27.2/kubernetes/typed/apps/v1#AppsV1Interface interface found!")
			}
		}(&workloadKinds[i])
	}
	wg.Wait()
	for _, w := range workloadKinds {
		if len(w.items) > 0 {
			res = append(res, w.items...)
		}
	}
	return res
}
