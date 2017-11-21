package ringN

type Ring struct {
  buf []byte
  off int
  end int
  N int
}

func New(n int) *Ring {
  r := &Ring{}
  r.buf = make([]byte,n)
  r.off = 0
  r.end = 0
  r.N = 0
  return r
}

func (r *Ring) Len() int {
  return c.N
}

func (r *Ring) Cap() int {
  return len(c.buf)
}
