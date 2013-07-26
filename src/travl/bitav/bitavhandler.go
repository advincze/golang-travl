package bitav

type BitAvHandler interface {
	FindOrNewBitAv(id string) BitAv
	SaveBitAv(id string, ba BitAv)
	FindBitAv(id string) BitAv
	DeleteBitAv(id string)
}

type BitAvHandlerMem struct {
	bitavs map[string]BitAv
}

func NewBitAvHandlerMem() *BitAvHandlerMem {
	return &BitAvHandlerMem{
		bitavs: make(map[string]BitAv),
	}
}

func (bah BitAvHandlerMem) SaveBitAv(id string, ba BitAv) {
	bah.bitavs[id] = ba
}

func (bah BitAvHandlerMem) FindBitAv(id string) BitAv {
	return bah.bitavs[id]
}

func (bah BitAvHandlerMem) FindOrNewBitAv(id string) BitAv {
	bitav := bah.FindBitAv(id)
	if bitav == nil {
		bitav = NewSegmentBitAvDefault(id)
		bah.SaveBitAv(id, bitav)
	}
	return bitav
}

func (bah BitAvHandlerMem) DeleteBitAv(id string) {
	delete(bah.bitavs, id)
}
