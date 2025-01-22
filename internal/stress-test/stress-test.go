package stresstest

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func Exec(url string, totalRequests, concurrency int) {
	fmt.Println("Iniciando teste de carga...")
	fmt.Printf("URL: %s\n", url)
	fmt.Printf("Total de Requests: %d\n", totalRequests)
	fmt.Printf("Chamadas Simultâneas: %d\n", concurrency)

	start := time.Now()
	var wg sync.WaitGroup
	requestsPerWorker := totalRequests / concurrency

	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < requestsPerWorker; j++ {
				resp, err := http.Get(url)
				if err != nil {
					fmt.Printf("Erro ao fazer a request: %s\n", err)
					continue
				}
				resp.Body.Close()
				fmt.Printf("Status: %s\n", resp.Status)
			}
		}()
	}
	wg.Wait()
	elapsed := time.Since(start)
	fmt.Printf("Teste de carga concluído em %s\n", elapsed)
}
