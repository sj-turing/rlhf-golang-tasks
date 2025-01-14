package main

import (
    "bytes"
    "fmt"
    "math/rand"
    "strings"
    "testing"
)

const numStrings = 200
const stringLength = 50

var inputStrings = generateStrings()

func generateStrings() []string {
    result := make([]string, numStrings)
    for i := range result {
        result[i] = getString(rand.Intn(stringLength))
    }
    return result
}

func getString(length int) string {
    const alphanum = "abcdefghijklmnopqrstuvwxyz0123456789"
    b := make([]byte, length)
    for i := range b {
        b[i] = alphanum[rand.Intn(len(alphanum))]
    }
    return string(b)
}

func concatBySprintf(strings []string) string {
    result := ""
    for _, s := range strings {
        result = fmt.Sprintf("%s%s", result, s)
    }
    return result
}

func concatByStringsJoin(strings []string) string {
    return strings.Join(strings, "")
}

func concatByBytesBuffer(strings []string) string {
    var buf bytes.Buffer
    for _, s := range strings {
        buf.WriteString(s)
    }
    return buf.String()
}

func concatByPlusOperator(strings []string) string {
    result := ""
    for _, s := range strings {
        result += s
    }
    return result
}

func concatByBytesAppend(strings []string) string {
    var result []byte
    for _, s := range strings {
        result = append(result, []byte(s)...)
    }
    return string(result)
}

func concatByBytesBuilder(strings []string) string {
    var b strings.Builder
    for _, s := range strings {
        b.WriteString(s)
    }
    return b.String()
}

func BenchmarkConcatBySprintf(b *testing.B) {
    for i := 0; i < b.N; i++ {
        concatBySprintf(inputStrings)
    }
}

func BenchmarkConcatByStringsJoin(b *testing.B) {
    for i := 0; i < b.N; i++ {
        concatByStringsJoin(inputStrings)
    }
}

func BenchmarkConcatByBytesBuffer(b *testing.B) {
    for i := 0; i < b.N; i++ {
        concatByBytesBuffer(inputStrings)
    }
}

func BenchmarkConcatByPlusOperator(b *testing.B) {
    for i := 0; i < b.N; i++ {
        concatByPlusOperator(inputStrings)
    }
}

func BenchmarkConcatByBytesAppend(b *testing.B) {
    for i := 0; i < b.N; i++ {
        concatByBytesAppend(inputStrings)
