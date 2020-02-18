/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Used to add numbers",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		floatFlagStatus, _ := cmd.Flags().GetBool("float")
		if floatFlagStatus {
			addFloat(args)
		} else {
			addInt(args)
		}
	},
}

func addFloat(args []string) {
	result := 0.0
	for _, num := range args {
		numInt, err := strconv.ParseFloat(num, 64)
		if err != nil {
			fmt.Println(err)
		}
		result = result + numInt
	}
	fmt.Printf("Addition of numbers %s \n%f\n:", args, result)
}

func addInt(args []string) {
	var result int64
	for _, num := range args {
		numInt, err := strconv.ParseInt(num, 10, 64)
		fmt.Println(numInt)
		if err != nil {
			fmt.Println(err)
		}
		result = result + numInt
	}
	fmt.Printf("Addition of numbers %s \n%d\n:", args, result)
}

func init() {
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	addCmd.Flags().BoolP("float", "f", false, "Add floating numbers ")
}
