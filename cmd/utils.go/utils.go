// fjwt v1.0.0
// Author: Alp Keskin
// Github: github.com/alpkeskin
// Linkedin: linkedin.com/in/alpkeskin
package utils

import (
	"bufio"
	"os"
	"sync/atomic"
	"time"

	"github.com/briandowns/spinner"
)

const Version = "v1.0.0"

var (
	Wordlist *string
	Threads  *int
	Counter  uint64 = 0
	Total    uint   = 0
	Result   string = ""
	Spinner         = spinner.New(spinner.CharSets[26], 100*time.Millisecond)
)

func ReadFileToStringList(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		Total++
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func PrintResult(secret string, duration time.Duration) {
	println("\r=====================================")
	println("\x1b[32m[FOUND]\x1b[0m", "Secret:", secret)
	println("Attempts:", atomic.LoadUint64(&Counter))
	println("Elapsed:", duration.Round(time.Second).String())
	println("=====================================")
}
