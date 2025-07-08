/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/spf13/cobra"
)

func PrettyPrintJSONstream(data interface{}) (string, error) {
	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(jsonBytes), nil
}

type PhoneBook []Record

func (p PhoneBook) Len() int {
	return len(p)
}

func (p PhoneBook) Less(i, j int) bool {
	return p[i].Name < p[j].Name // сортировка по имени
}

func (p PhoneBook) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func list() {
	sort.Sort(PhoneBook(data))
	text, err := PrettyPrintJSONstream(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
	fmt.Printf("%d records in total.\n", len(data))
}

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list an entry",
	Long:  `list an entry from the phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
