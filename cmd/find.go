/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"
	"fmt"
	"strings"
	"todos/todo"
	"text/tabwriter"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var searchPhrase string

// findCmd represents the find command
var findCmd = &cobra.Command{
	Use:   "find",
	Short: "Search for a task in todos",
	Long: `Search for a task in todos`,
	Run: findRun,
}

func findRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	fmt.Println(items)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 1, ' ', 0)
	for _, item := range items {
		if strings.Contains(item.Text, searchPhrase) {
			fmt.Fprintln(w, item.Label()+"\t"+item.PrettyDone()+"\t"+item.PrettyP()+"\t"+item.Text+"\t")
		}
	}
	w.Flush()
}

func init() {
	rootCmd.AddCommand(findCmd)

	// Here you will define your flags and configuration settings.
	findCmd.Flags().StringVar(&searchPhrase, "search", "", "search phrase to look in the todos")

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// findCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// findCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
