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
 * Complete the 'minimumNumber' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER n
 *  2. STRING password
 */

func minimumNumber(n int32, password string) int32 {
    number := "0123456789"
    lower := "abcdefghijklmnopqrstuvwxyz"
    upper := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
    special := "!@#$%^&*()-+"
    
    var hasNum, hasLower, hasUpper, hasSpecial bool
    var missing int32
    
    for _, ch := range password {
        
        if strings.ContainsRune(number, ch){
            hasNum = true
        } else if strings.ContainsRune(lower, ch){
            hasLower = true
        } else if strings.ContainsRune(upper, ch){
            hasUpper = true
        } else if strings.ContainsRune(special, ch){
            hasSpecial = true
        }
    }
    
    if !hasNum {
        missing++
    }
    
    if !hasLower {
        missing++
    }
    
    if !hasUpper {
        missing++
    }
    
    if !hasSpecial {
        missing++
    }
    
    if missing > 6 - n {
        return missing
    }
    
    return 6 - n
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    nTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
    checkError(err)
    n := int32(nTemp)

    password := readLine(reader)

    answer := minimumNumber(n, password)

    fmt.Fprintf(writer, "%d\n", answer)

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
