package stresstest

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	str "stress-test/internal/report"
)

func Exec(url string, totalRequests, concurrency int) {
	fmt.Println("Iniciando teste de carga...")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Total de Requests: %d\n", totalRequests)
	fmt.Printf("Chamadas Simultâneas: %d\n", concurrency)

	// Inicializa o relatório
	report := str.TReport{
		TotalRequests:      totalRequests,
		SuccessfulRequests: 0,
		StatusDistribution: make(map[string]int),
	}

	start := time.Now()
	var wg sync.WaitGroup

	// Mutex para proteger o acesso ao relatório
	var mutex sync.Mutex

	// Canal para distribuir as requisições entre os workers
	requests := make(chan int, totalRequests)
	for i := 0; i < totalRequests; i++ {
		requests <- i
	}
	close(requests)

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for range requests {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Erro ao fazer a request: %s\n", err)
					continue
				}
				resp.Body.Close()

				// Atualiza o relatório
				mutex.Lock()
				status := resp.Status
				report.StatusDistribution[status]++
				if status == "200 OK" {
					report.SuccessfulRequests++
				}
				mutex.Unlock()

				// fmt.Printf("Status: %s\n", status)
			}
		}()
	}

	wg.Wait()
	elapsed := time.Since(start)
	report.TotalTime = elapsed

	// Exibe o relatório
	str.GenerateReport(&report)
}
