package cmd

import (
	"github.com/spf13/cobra"
	"os"
	"log"
	"fmt"
)

var kubeconfigPath string

func getDefaultKubeconfig() string {
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s/%s", home, ".kube/config")
}

func init() {
	rootCmd.PersistentFlags().StringVar(&kubeconfigPath, "kubeconfig", getDefaultKubeconfig(), "The path to your kube config file")
	rootCmd.AddCommand(scanCmd)
}

var (
	rootCmd = &cobra.Command{
		Use:           "k8s-value-scanner",
		Short:         "k8s-value-scanner command-line tool",
		Long:          `k8s-value-scanner will search your entire cluster for the value you will give as argument`,
		SilenceErrors: true,
		SilenceUsage:  true,
	}
)

func Execute() error {
	return rootCmd.Execute()
}