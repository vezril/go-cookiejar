package trit

type Trit struct {
  // Doing v0 and v1 instead of array to not fiddle around with endianess
  // LSB
  v0 bool
  // MSB
  v1 bool
  err bool
  over bool
}

func New() *Trit {
  t := new(Trit)

  return t
}

func (t *Trit) Set(v interface{}) {
  conv, ok := v.(int)
  if ok != nil || conv >= 4  {
    panic(ok)
  }

  t.v0 = bool(conv % 2)
  t.v1 = bool(conv / 2)
}
