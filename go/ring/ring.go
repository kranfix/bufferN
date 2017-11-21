package ringN

struct Ring {
  buf []byte
  off int
  end int
  N int
}

func New(n int) *RingN {
  c := &Circular{}
  c.buf = make([]byte,n)
  c.off = 0
  c.end = 0
  c.N = 0
}
