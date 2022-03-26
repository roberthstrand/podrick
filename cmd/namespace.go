/*
Copyright © 2022 Roberth Strand (rob@podrick.dev)

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"

	"github.com/roberthstrand/podrick/pkg/manifests"
)

// namespaceCmd represents the namespace command
var namespaceCmd = &cobra.Command{
	Use:   "namespace",
	Short: "Generate namespace manifest",
	Long:  `Generates a manifest for deploying namespaces.`,
	Run: func(cmd *cobra.Command, args []string) {
		// If the user specified a name through the flags, use that.
		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}

		// If the user didn't set a name through flags,
		// ask the user what name they want to use.
		if name == "not_set" {
			fmt.Println("Name:")
			fmt.Scanln(&name)
		}

		// todo: add support for labels
		// todo: add support for annotations

		manifest := manifests.CreateNamespace(name)

		fmt.Println(manifest)

		d, err := yaml.Marshal(&manifest)
		if err != nil {
			log.Fatal(err)
		}

		filename := fmt.Sprintf("namespace-%s.yaml", name)
		writeErr := ioutil.WriteFile(filename, d, 0664)
		if writeErr != nil {
			log.Fatal(err)
		}

		fmt.Println("Namespace manifest generated...")
	},
}

func init() {
	generateCmd.AddCommand(namespaceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// namespaceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// namespaceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	namespaceCmd.Flags().StringP("name", "n", "not_set", "The name of the namespace")
	namespaceCmd.Flags().StringP("labels", "l", "", "The labels of the namespace")
}
