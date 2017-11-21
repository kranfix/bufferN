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
