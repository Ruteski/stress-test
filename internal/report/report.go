package report

import (
	"fmt"
	"time"
)

type TReport struct {
	TotalRequests      int
	SuccessfulRequests int
	StatusDistribution map[string]int
	TotalTime          time.Duration
}

// Função para gerar o relatório
func GenerateReport(report *TReport) {
	fmt.Println("\nRelatório do Teste de Carga:")
	fmt.Println("----------------------------")
	fmt.Printf("Tempo Total Gasto: %s\n", report.TotalTime)
	fmt.Printf("Total de Requests Realizados: %d\n", report.TotalRequests)
	fmt.Printf("Requests com Status 200: %d\n", report.SuccessfulRequests)

	fmt.Println("\nDistribuição de Status HTTP:")
	for status, count := range report.StatusDistribution {
		if status != "200 OK" {
			fmt.Printf("%s: %d\n", status, count)
		}
	}
}
