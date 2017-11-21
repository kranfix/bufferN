package ringN

/*
Write writes the ring buffer with buf data
*/
func (r *Ring) Write(buf []byte) (n int) {
  for n = 0; r.N < len(r.buf) && n < len(buf); n++ {
    end := r.end
    if  end == len(r.buf) {
     end = 0
    }
    r.buf[end] = buf[n]
    end++

    r.end = end
    r.N++
  }
  return n
}
