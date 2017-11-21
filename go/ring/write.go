package ringN

func (r *Ring) Write(buf []byte) (n int) {
  for n = 0; r.N > 0 && n < len(buf); n++ {
    off := r.off
    buf[n] = r.buf[off]

    off++
    if off == len(r.buf){
      off = 0
    }

    r.off = off
    r.N--
  }
  return n
}
