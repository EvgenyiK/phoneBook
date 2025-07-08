/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func insert(record Record) {
	data = append(data, record)
}

// insertCmd represents the insert command
var insertCmd = &cobra.Command{
	Use:   "insert",
	Short: "insert an entry",
	Long:  `insert an entry from the phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		phone, _ := cmd.Flags().GetString("phone")

		if name == "" || phone == "" {
			fmt.Println("Both --name and --phone flags are required.")
			return
		}

		newRecord := Record{
			Name:  name,
			Phone: phone,
		}

		insert(newRecord)
		fmt.Printf("Inserted: %s - %s\n", newRecord.Name, newRecord.Phone)
	},
}

func init() {
	rootCmd.AddCommand(insertCmd)

	insertCmd.Flags().StringP("name", "n", "", "Name of the contact")
	insertCmd.Flags().StringP("phone", "p", "", "Phone number of the contact")
}
