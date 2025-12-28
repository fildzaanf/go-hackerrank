package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

/*
 * Complete the 'steadyGene' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts STRING gene as parameter.
 */

func steadyGene(gene string) int32 {
    n := len(gene)
    target := n / 4

    count := make(map[byte]int)
    for i := 0; i < n; i++ {
        count[gene[i]]++
    }

    ok := true
    for _, c := range []byte{'A', 'C', 'G', 'T'} {
        if count[c] > target {
            ok = false
            break
        }
    }
    if ok {
        return 0
    }

    minLen := n
    left := 0

    for right := 0; right < n; right++ {
        count[gene[right]]--

        for left <= right &&
            count['A'] <= target &&
            count['C'] <= target &&
            count['G'] <= target &&
            count['T'] <= target {

            if right-left+1 < minLen {
                minLen = right - left + 1
            }

            count[gene[left]]++
            left++
        }
    }

    return int32(minLen)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)
    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16*1024*1024)

    _, err = strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)

    gene := readLine(reader)

    result := steadyGene(gene)
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
