package main

import (
  "os"
  "fmt"
  "time"
  "github.com/kranfix/bufferN/go/ring"
)

func main() {
  r := ringN.New(5)
  s := []byte("Hello! this is a text to test the WriteTo and ReadFrom methods\n")
  f,_ := os.Create("hello.txt")
  defer f.Close()

  N1 := 0
  N2 := int64(0)
  stop := false

  go func(){
    for {
      n,_ := r.Write(s[N1:])
      N1 += n
      fmt.Printf("Write: n=%d N1=%d\n",n,N1)
      if N1 == len(s) {
        stop = true
        return
      }
      time.Sleep(10 * time.Millisecond)
    }
  }()

  for {
    n,_ := r.WriteTo(f)
    N2 += n
    //fmt.Printf("Numeros %d %d\n",n,N1)
    if N2 >= int64(N1)  && stop {
      break
    }
    time.Sleep(10 * time.Millisecond)
  }
}
