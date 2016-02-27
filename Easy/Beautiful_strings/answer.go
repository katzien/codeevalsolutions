package main

import (
  "fmt"
  "regexp"
  "strings"
  "sort"
  "bufio"
  "os"
  "log"
)

var (
	reg = regexp.MustCompile("[^a-zA-Z]+")
)

type Pair struct {
  Letter string
  Count int
}

type PairList []Pair

func main() {
    file, err := os.Open(os.Args[1])
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(countScore(scanner.Text()))
    }
}

func countScore(line string) (int) {
  sanitizedLine := prepLine(line)

  totals := map[string]int {}

  for _, l := range sanitizedLine {
    letter := string(l)

    if totals[letter] != 0 {
        totals[letter]++
    } else {
        totals[letter] = 1
    }
  }

  ordered := rankByLetterCount(totals)

  total := 0
  nextScore := 26

  for _, f := range ordered {
    total = total + (f.Count * nextScore)
    nextScore--
  }

  return total
}

// removes non-letter characters and lowercases the input string
func prepLine(line string) (string) {
  return strings.ToLower(reg.ReplaceAllString(line, ""))
}

func rankByLetterCount(letterFrequencies map[string]int) PairList {
  pl := make(PairList, len(letterFrequencies))
  i := 0
  for k, v := range letterFrequencies {
    pl[i] = Pair{k, v}
    i++
  }
  sort.Sort(sort.Reverse(pl))
  return pl
}

func (p PairList) Len() int { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Count < p[j].Count }
func (p PairList) Swap(i, j int){ p[i], p[j] = p[j], p[i] }
