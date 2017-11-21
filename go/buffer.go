package Buffer

struct Buffer{
  buf []byte
}

func New(n int) *Buffer {
  b := &Buffer{}
  b.SetSlice(make([]byte,n))
}

func (b *buffer) SetSlice(s []byte) {
  b.buf = s
}
