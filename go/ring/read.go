package ringN

func (c *Ring) Read(buf []byte) (n int) {
  for n = 0; c.N < len(c.buf) && n < len(buf); n++ {
    end := c.end
    if  end == len(c.buf) {
     end = 0
    }
    c.buf[end] = buf[n]
    end++

    c.end = end
    c.N++
  }
  return n
}
