// Найдите веб-сайт, который содержит большое количество дан­ных.
// Исследуйте работу кеширования путем двукратного запуска fetchall и сравнения времени запросов.
// Получаете ли вы каждый раз одно и то же содержимое?
// Измените fetchall так, чтобы вывод осуществлялся в файл и чтобы затем можно было его изучить.
package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	for range 2 {
		run()
	}
}

func run() {
	start := time.Now()
	ch := make(chan string)

	out, err := os.OpenFile("output.txt", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("%v", err)
		return
	}
	defer out.Close()

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Запуск go-подпрограммы
	}

	for range os.Args[1:] {
		line := fmt.Sprintf("%s\n", <-ch)
		out.WriteString(line)
	}

	line := fmt.Sprintf("%.2fs elapsed\n\n", time.Since(start).Seconds())
	out.WriteString(line)
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // Отправка в канал ch
		return
	}
	defer resp.Body.Close()

	hasher := sha256.New()
	nbytes, err := io.Copy(hasher, resp.Body)
	if err != nil {
		ch <- fmt.Sprintf("while reading: %s: %v", url, err)
		return
	}
	bodyHash := fmt.Sprintf("%X", hasher.Sum(nil))
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s %s", secs, nbytes, url[:50], bodyHash)
}
