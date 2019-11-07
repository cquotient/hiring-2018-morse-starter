package main

import (
    "bufio"
    "flag"
    "fmt"
    "os"
    "strconv"
    "strings"
)

const MAX_CONTEXT_WORD_LEN = 10
const MAX_CONTEXT_WORDS = 100
const MAX_MORSE_WORD_LEN = 80
const MAX_MORSE_CHAR_LEN = 6

const DELIMITER = "*"

var err error
var errs = make(chan error, 1)
var contextWordsCount int

var numLines int
var morseCode = make(map[string]string)
// This is set up as a map to easily check if a contextWord exists. Feel free to convert to a slice if you prefer
var contextWords = make(map[string]struct{})
var morseWords []string

func main() {
    flag.Parse()

    go WatchErrors()

    ProcessInput()

    Decode()
}

func WatchErrors() {
    for {
        err := <-errs
        if err != nil {
            log.Fatalf("Failed: %v", err)
        }
    }
}

func ProcessInput() {
    var ProcessPhase = map[int]func(string){
        0: NumLines,
        1: MorseCode,
        2: ContextWords,
        3: MorseWords,
    }
    var phase int
    var lines int

    scanner := bufio.NewScanner(os.Stdin)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())

        // Advance to next phase if delimiter
        if line == DELIMITER {
            phase = phase + 1
            continue
        }

        ProcessPhase[phase](line)

        lines = lines + 1
        if lines >= numLines {
            break
        }
        if lines == 1 {
            // Advance to next phase after reading numLines from line 1
            phase = phase + 1
        }
    }
    errs <- scanner.Err()
}

func NumLines(line string) {
    numLines, err = strconv.Atoi(line)
    errs <- err
}

func MorseCode(line string) {
    items := strings.Fields(line)
    if len(items) < 2 {
        errs <- fmt.Errorf("morse code table doesn't have two items")
    }
    letter := strings.TrimSpace(items[0])
    code := strings.TrimSpace(items[1])
    if len(items[1]) > MAX_MORSE_CHAR_LEN {
        code = items[1][0:MAX_MORSE_CHAR_LEN]
    }

    morseCode[letter] = code
}

func ContextWords(line string) {
    contextWordsCount = contextWordsCount + 1
    if contextWordsCount > MAX_CONTEXT_WORDS {
        return
    }
    if len(line) <= MAX_CONTEXT_WORD_LEN {
        contextWords[line] = struct{}{}
    }
}

func MorseWords(line string) {
    items := strings.Fields(line)
    for _, item := range items {
        morseWord := strings.TrimSpace(item)
        if len(morseWord) <= MAX_MORSE_WORD_LEN {
            morseWords = append(morseWords, morseWord)
        }
    }
}

func Decode() {
    // TODO: decode and print output to stdout
}