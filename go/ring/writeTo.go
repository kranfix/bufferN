package ringN

import "io"

func (r *Ring) WriteTo(w io.Writer) (n int64, err error) {
  N := r.Len()
  s := make([]byte,n)
  N,err = r.Read(s)
  if err != nil {
    return
  }

  N,err = w.Write(s)
  if err != nil {
    return
  }

  n := int64(N)
  return
}
