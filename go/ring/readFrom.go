package ringN

import "io"
//import "fmt"

func (r *Ring) ReadFrom(rio io.Reader) (n int64, err error) {
  n = int64(0)
  off := r.off
  end := r.end
  L := r.Len()
  N := 0

  if end == r.Cap() {
    end = 0
  }

  if off < end || L == 0 {
    N,err = rio.Read(r.buf[end:])
    n = int64(N)
    end += N
    r.N += N
    r.end = end
    //fmt.Printf("from 1: end=%d off=%d\n",end,off)
    if end < r.Cap() || err != nil {
      return
    }
  }

  //fmt.Printf("from 2: end=%d off=%d\n",end,off)
  if end == r.Cap() {
    end = 0
  }

  N,err = rio.Read(r.buf[end:off])
  n += int64(N)
  r.N += N
  r.end = end + N
  return
}
