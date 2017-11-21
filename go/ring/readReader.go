package ringN

import "io"

func (c *Ring) ReadReader(r io.Reader) (n int) {
  off := c.off
  end := end

  if c.N == len(c.buf) {
    return 0
  }

  if end < off {
    n,_ = r.Read(c.buf[end:off])
    c.end += n
  } else if end < len(c.buf) {
    n,_ = r.Read(c.buf[end:])
    c.end += n
  } else {
    n,_ = r.Read(c.buf[end:])
    c.end = n
  }
  return
}
