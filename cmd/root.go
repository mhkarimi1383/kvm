/*
Copyright © 2022 Muhammed Hussein Karimi <info@karimi.dev>
*/
package cmd

import (
	"errors"
	"fmt"
	"kvm/helper"
	"os"

	"github.com/ktr0731/go-fuzzyfinder"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "kvm",
	Short: "Easily switch between kubectl version",
	Long: `Managing kubectl versions can help you when managing multiple clusters
with different versions, can make you fully compatible with your clusters`,
	Run: func(cmd *cobra.Command, args []string) {
		tags, err := helper.GetVersions()
		if err != nil {
			panic(err)
		}
		idx, err := fuzzyfinder.Find(tags,
			func(i int) string {
				return tags[i]
			}, fuzzyfinder.WithPreviewWindow(func(i, w, h int) string {
				if i == -1 {
					return ""
				}
				return fmt.Sprintf("Change log and more info:\n  https://github.com/kubernetes/kubernetes/releases/tag/%s", tags[i])
			}))
		if err != nil {
			panic(err)
		}
		version := tags[idx]
		homePath, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		fmt.Printf("selected: %v\n", version)
		basePath := homePath + "/.kvm/"
		binPath := homePath + "/.kvm/bin/"
		if _, err := os.Stat(basePath + "kubectl-" + version); errors.Is(err, os.ErrNotExist) {
			path, err := helper.DownloadKubectlBinary(version)
			if err != nil {
				panic(err)
			}
			err = helper.MoveFile(path, basePath+"kubectl-"+version)
			if err != nil {
				panic(err)
			}
		} else if err != nil {
			panic(err)
		}
		os.Remove(binPath + "kubectl")
		err = os.Symlink(basePath+"kubectl-"+version, binPath+"kubectl")
		if err != nil {
			panic(err)
		}
		err = os.Chmod(basePath+"kubectl-"+version, 0775)
		if err != nil {
			panic(err)
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kvm.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
