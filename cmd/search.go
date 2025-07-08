/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "search an entry",
	Long:  `search an entry from the phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		if name == "" {
			fmt.Println("Please provide a name using --name flag.")
			return
		}

		index, rec := search(name)
		if index == -1 {
			fmt.Printf("No entry found for '%s'.\n", name)
		} else {
			fmt.Printf("Found: %s - %s\n", rec.Name, rec.Phone)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	searchCmd.Flags().StringP("name", "n", "", "Name of the contact to search")
}
