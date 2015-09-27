package main

import (
    "golang.org/x/tour/wc"
    "strings"
)

func WordCount(s string) map[string]int {
    ss := strings.Fields(s)
    m := make(map[string]int)

    for i := 0; i < len(ss); i++ {
        v, ok := m[ss[i]]
        if ok {
            m[ss[i]] = v + 1
            } else {
                m[ss[i]] = 1
            }
        }
    return map[string]int{"x": 1}
}

func main() {
    wc.Test(WordCount)
}
