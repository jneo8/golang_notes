// Fetchall fetch URLS in parallel and report their times and sizes.

package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
)


func main() {
    start := time.Now()
    fmt.Printf("start: %s \n", start)

    ch := make (chan string)
    for _, url := range os.Args[1:] {
        go fetch(url, ch)  // start a goroutine
    }
    for range os.Args[1:] {
        fmt.Println(<-ch)  // receive from channel ch
    }
    fmt.Printf("%.2fs elapesed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string ) {
    start := time.Now()

    resp, err := http.Get(url)
    if err != nil{
        ch <- fmt.Sprint(err)  // send to channel ch
        return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }

    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("Seconds: %.2fs bytes: %7d url: %s", secs, nbytes, url)
}
