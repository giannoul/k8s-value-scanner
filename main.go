package main

import (
	"fmt"
	"github.com/giannoul/k8s-value-scanner/cmd"
)

func main(){
	err := cmd.Execute()
	if err != nil && err.Error() != "" {
		fmt.Println(err)
	}
}