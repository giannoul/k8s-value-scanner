package cmd

import (
	"github.com/spf13/cobra"
	"github.com/giannoul/k8s-value-scanner/internal/scanner"
	"fmt"
	"os"
)

var (
	scanCmd = &cobra.Command{
		Use:   "scan",
		Short: "scan will search for the given value",
		Long:  ``,
		Run: func (cmd *cobra.Command, args []string) {
			if len(args) > 0 {
				scanner.Scan(kubeconfigPath, args[0])
			} else {
				fmt.Fprintln(os.Stderr, "No value was given for search, exiting.")
				return
			}
			
		},
	}
)

