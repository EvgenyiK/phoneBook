/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
)

type Record struct {
	Name  string
	Phone string
}

var data []Record

// Serialize сериализует переданные данные в JSON и записывает их в writer
func Serialize(data interface{}, writer io.Writer) error {
	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "  ") // для читаемого формата, по желанию
	return encoder.Encode(data)
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = Serialize(&data, f)
	if err != nil {
		return err
	}
	return nil
}

// search ищет запись по имени
func search(name string) (int, Record) {
	for i, rec := range data {
		if rec.Name == name {
			return i, rec
		}
	}
	return -1, Record{}
}

// deleteEntry удаляет запись по имени
func deleteEntry(name string) error {
	index, _ := search(name)
	if index == -1 {
		return fmt.Errorf("Запись с именем '%s' не найдена", name)
	}
	data = append(data[:index], data[index+1:]...)
	return nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phoneBook",
	Short: "Phone Book Application",
	Long:  `A simple phone book application to manage contacts.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
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

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
