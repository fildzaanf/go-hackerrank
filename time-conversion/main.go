package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strings"
    "strconv"
)

/*
 * Complete the 'timeConversion' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts STRING s as parameter.
 */

func timeConversion(s string) string {
    hh := s[:2]
    mm := s[3:5]
    ss := s[6:8]
    format := s[8:]
    
    if format == "AM" {
        if hh == "12" {
            return "00" + ":" + mm + ":" + ss
        } 

     	return hh + ":" + mm + ":" + ss
    } else if format == "PM" {
        if hh == "12" {
            return "12" + ":" + mm + ":" + ss
        }  
    }
    
    hourInt, _ := strconv.Atoi(hh)
    hourInt += 12
    return fmt.Sprintf("%d:%s:%s", hourInt, mm, ss)
    
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 16 * 1024 * 1024)

    stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 16 * 1024 * 1024)

    s := readLine(reader)

    result := timeConversion(s)

    fmt.Fprintf(writer, "%s\n", result)

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
