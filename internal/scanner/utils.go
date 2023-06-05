package scanner

import (
	apiCorev1 "k8s.io/api/core/v1"
)

func getContainerEnvVars(containers []apiCorev1.Container, items *[]Item, kind string, name string, namespace string) {
	for _, c := range containers {
		// https://pkg.go.dev/k8s.io/api/core/v1#Container
		// https://pkg.go.dev/k8s.io/api/core/v1#EnvVar
		for _, p := range c.Env {
			*items = append(*items, Item{kind: kind, name: name, namespace: namespace, key: p.Name, value: p.Value})
		}
	}
}

func getConfigmapContents(m map[string]string, items *[]Item, kind string, name string, namespace string) {
	for key, value := range m {
		*items = append(*items, Item{kind: kind, name: name, namespace: namespace, key: key, value: value})
	}
}

func getSecretContents(m map[string][]byte, items *[]Item, kind string, name string, namespace string) {
	for key, value := range m {
		*items = append(*items, Item{kind: kind, name: name, namespace: namespace, key: key, value: string(value)})
	}
}