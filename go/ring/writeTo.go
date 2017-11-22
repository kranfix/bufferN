package ringN

import "io"

func (r *Ring) WriteTo(w io.Writer) (n int64, err error) {
  n = int64(0)
  off := r.off
  end := r.end
  L := r.Len()
  N := 0

  if off >= end && L > 0 {
    N,err = w.Write(r.buf[off:])
    n = int64(N)
    off += N

    r.N -= N
    //fmt.Printf("from 1: end=%d off=%d\n",end,off)
    if off == r.Cap() && end != off {
      off = 0
    }
    r.off = off

    if err != nil {
      return
    }
  }


  if off == r.Cap() && end == off && r.N > 0 {
    off = 0
  }
  N,err = w.Write(r.buf[off:end])
  n += int64(N)
  r.N -= N
  r.off = off + N
  return
}
