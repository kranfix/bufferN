package ringN

import "io"

func (r *Ring) ReadFrom(rio io.Reader) (n int64, err error) {
  n := int64(0)
  off := r.off
  end := r.end
  L := r.Len()
  N := 0
  if end < off {
    N,err = rio.Read(c.buf[end:off])
    n = int64(N)
    end += N
  } else if L > 0 {
    N,err = rio.Read(c.buf[end:])
    n = int64(N)
    end += N

    if end < r.Cap {
      goto
    }

    if err != nil {
      return
    }

    N,err = rio.Read(c.buf[:off])
    n += int64(N)
    end = N
  }
  return
}
