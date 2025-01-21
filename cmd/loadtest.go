/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// loadtestCmd represents the loadtest command
var loadtestCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flag("url").Value.String()
		concurrency := cmd.Flag("concurrency").Value
		requests := cmd.Flag("requests").Value.String()
		fmt.Println("loadtest called")
		fmt.Println("URL:", url)
		fmt.Println("Concurrency:", concurrency)
		fmt.Println("Requests:", requests)
	},
}

func init() {
	rootCmd.AddCommand(loadtestCmd)
	loadtestCmd.Flags().StringP("url", "u", "", "URL para testar")
	loadtestCmd.MarkFlagRequired("url")
	loadtestCmd.Flags().IntP("concurrency", "c", 1, "Quantidade de chamadas simultâneas")
	loadtestCmd.Flags().IntP("requests", "r", 1, "Número total de requests")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loadtestCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loadtestCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
