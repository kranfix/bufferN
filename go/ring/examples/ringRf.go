package main

import (
  "os"
  "fmt"
  "time"
  "github.com/kranfix/bufferN/go/ring"
)

func main() {
  r := ringN.New(5)
  f,_ := os.Open("hello.txt")
  defer f.Close()

  N1 := int64(0)
  N2 := int64(0)
  stop := false
  go func(){
    for {
      n,err := r.ReadFrom(f)
      N1 += n
      //fmt.Printf("Numeros %d %d\n",n,N1)
      if err != nil {
        stop = true
        return
      }
      time.Sleep(10 * time.Millisecond)
    }
  }()

  s := make([]byte,3)
  for {
    n := r.Read(s)
    N2 += int64(n)
    fmt.Printf("%s",s[:n])
    if N2 >= N1 && stop {
      break
    }
    time.Sleep(10 * time.Millisecond)
  }
  fmt.Println()
}
