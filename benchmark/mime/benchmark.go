package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/akashrtd/mime/pkg/mime"
)

const (
	iterations = 5
	testURL    = "https://example.com"
)

type Results struct {
	Startup    []int64 `json:"startup"`
	Navigation []int64 `json:"navigation"`
	Extraction []int64 `json:"extraction"`
	Screenshot []int64 `json:"screenshot"`
	Total      []int64 `json:"total"`
}

type Output struct {
	Tool       string             `json:"tool"`
	Iterations int                `json:"iterations"`
	Averages   map[string]float64 `json:"averages"`
	Raw        Results            `json:"raw"`
}

func main() {
	fmt.Println("=== MIME Benchmark ===")

	results := Results{
		Startup:    make([]int64, 0, iterations),
		Navigation: make([]int64, 0, iterations),
		Extraction: make([]int64, 0, iterations),
		Screenshot: make([]int64, 0, iterations),
		Total:      make([]int64, 0, iterations),
	}

	ctx := context.Background()

	for i := 0; i < iterations; i++ {
		fmt.Printf("Run %d/%d\n", i+1, iterations)

		runStart := time.Now()

		// Startup time
		startupStart := time.Now()
		m, err := mime.New(ctx)
		if err != nil {
			fmt.Printf("Error creating MIME: %v\n", err)
			return
		}
		results.Startup = append(results.Startup, time.Since(startupStart).Milliseconds())

		// Navigation time
		navStart := time.Now()
		if err := m.Navigate(testURL); err != nil {
			fmt.Printf("Error navigating: %v\n", err)
			m.Close()
			return
		}
		// Wait for page load
		time.Sleep(100 * time.Millisecond)
		results.Navigation = append(results.Navigation, time.Since(navStart).Milliseconds())

		// Extraction time
		extractStart := time.Now()
		_, _ = m.Extract("h1")
		_, _ = m.HTML()
		results.Extraction = append(results.Extraction, time.Since(extractStart).Milliseconds())

		// Screenshot time
		screenshotStart := time.Now()
		_, _ = m.Screenshot()
		results.Screenshot = append(results.Screenshot, time.Since(screenshotStart).Milliseconds())

		m.Close()

		results.Total = append(results.Total, time.Since(runStart).Milliseconds())
	}

	// Calculate averages
	fmt.Printf("\n=== Results (average of %d runs) ===\n", iterations)
	fmt.Printf("Startup:     %d ms\n", avg(results.Startup))
	fmt.Printf("Navigation:  %d ms\n", avg(results.Navigation))
	fmt.Printf("Extraction:  %d ms\n", avg(results.Extraction))
	fmt.Printf("Screenshot:  %d ms\n", avg(results.Screenshot))
	fmt.Printf("Total:       %d ms\n", avg(results.Total))

	// Output JSON
	output := Output{
		Tool:       "mime",
		Iterations: iterations,
		Averages: map[string]float64{
			"startup":    float64(avg(results.Startup)),
			"navigation": float64(avg(results.Navigation)),
			"extraction": float64(avg(results.Extraction)),
			"screenshot": float64(avg(results.Screenshot)),
			"total":      float64(avg(results.Total)),
		},
		Raw: results,
	}

	fmt.Println("\nJSON:")
	jsonData, _ := json.Marshal(output)
	fmt.Println(string(jsonData))
}

func avg(arr []int64) int64 {
	if len(arr) == 0 {
		return 0
	}
	var sum int64
	for _, v := range arr {
		sum += v
	}
	return sum / int64(len(arr))
}
