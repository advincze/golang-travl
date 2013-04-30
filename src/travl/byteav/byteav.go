package byteav

type ByteAv struct {
}

func New() *ByteAv {
	return &ByteAv{}
}

func (b *ByteAv) Len() int {
	return 0
}
