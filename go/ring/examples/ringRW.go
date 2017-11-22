package main

import (
  "fmt"
  "github.com/kranfix/bufferN/go/ring"
)

func main(){
  b := [...]byte{11,12,13,14,15,16}
  s1 := b[:]
  s2 := make([]byte,5)

  r := ringN.New(6)

  n1,_ := r.Write(s1)
  n2,_ := r.Read(s2)
  fmt.Println("First try")
  fmt.Printf("s1 (%d): % x\n",n1,s1)
  fmt.Printf("s2 (%d): % x\n",n2,s2)

  s1 = b[1:5]
  n1,_ = r.Write(s1)
  n2,_ = r.Read(s2[0:3])
  fmt.Println("\nSecond read")
  fmt.Printf("s1 (%d): % x\n",n1,s1)
  fmt.Printf("s2 (%d): % x\n",n2,s2)

  n2,_ = r.Read(s2[0:3])
  fmt.Println("\nThird read")
  fmt.Printf("s1 (%d): % x\n",n1,s1)
  fmt.Printf("s2 (%d): % x\n",n2,s2)
}
