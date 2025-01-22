/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	stresstest "stress-test/internal/stress-test"

	"github.com/spf13/cobra"
)

// loadtestCmd represents the loadtest command
var loadtestCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "Realiza um teste de carga em um serviço web",
	Long:  `Realiza um teste de carga em um serviço web, enviando um número específico de requisições com um nível de concorrência definido.`,

	Run: func(cmd *cobra.Command, args []string) {
		url := cmd.Flag("url").Value.String()

		// usando o Flags().GetInt()
		totalRequests, _ := cmd.Flags().GetInt("requests")

		// o valor da flag concurrency é um ponteiro para um inteiro, passado por referência
		//concurrency := cmd.Flag("concurrency").Value

		stresstest.Exec(url, totalRequests, concurrency)

	},
}

var concurrency int

func init() {
	rootCmd.AddCommand(loadtestCmd)

	loadtestCmd.Flags().StringP("url", "u", "", "URL para testar")
	loadtestCmd.MarkFlagRequired("url")

	loadtestCmd.Flags().IntP("requests", "r", 1, "Número total de requests")
	loadtestCmd.Flags().IntVarP(&concurrency, "concurrency", "c", 1, "Número de requests concorrentes")
}
