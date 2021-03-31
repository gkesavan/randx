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
	"fmt"

	"github.com/spf13/cobra"
)

var (
	countFlagNumbers int
	rangeFlagNumbers []string
	countFlag        int
	rangeFlag        []string
)

// numbersCmd represents the numbers command
var numbersCmd = &cobra.Command{
	Use:   "numbers",
	Short: "Returns random numbers",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("numbers called")
		fmt.Println("--count:", countFlagNumbers)
		fmt.Println("--range:", rangeFlagNumbers)
		fmt.Println("--verbose:", verbose)
	},
}

func init() {
	rootCmd.AddCommand(numbersCmd)

	// Here you will define your flags and configuration settings.
	numbersCmd.Flags().IntVarP(
		&countFlag, "count", "c", 0,
		"A count of random numbers",
	)
	numbersCmd.MarkFlagRequired("count")

	numbersCmd.Flags().StringSliceVarP(
		&rangeFlag, "range", "r", []string{"1:100"},
		"Range of numbers. Optional",
	)
	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// numbersCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// numbersCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
