package ringN

func (c *RingN) Write(buf []byte) (n int) {
  for n = 0; c.N > 0 && n < len(buf); n++ {
    off := c.off
    buf[k] = c.buf[c.off]

    off++
    if off == len(c.buf){
      off = 0
    }

    c.off = off
    c.N--
  }
  return n
}
