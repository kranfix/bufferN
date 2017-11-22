package ringN

/*
Read reads the ring buffer and stores it on buf data
*/
func (r *Ring) Read(buf []byte) (n int,err error) {
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
  return
}
