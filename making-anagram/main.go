package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
)

/*
 * Complete the 'makingAnagrams' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. STRING s1
 *  2. STRING s2
 */

func makingAnagrams(s1 string, s2 string) int32 {
    frequency := make([]int, 26)
    
    for _, ch := range s1 {
        frequency[ch - 'a']++
    }
    
    for _, ch := range s2 {
        frequency[ch - 'a']--
    }
    
    del := 0
    for _, count := range frequency {
        if count < 0 {
            del -= count
        } else {
            del += count
        }
    }
    
    return int32(del)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s1 := readLine(reader)

    s2 := readLine(reader)

    result := makingAnagrams(s1, s2)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
