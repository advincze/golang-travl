package av

type TimeResolution int

const (
	sec      TimeResolution = 1
	Minute   TimeResolution = sec * 60
	Minute5  TimeResolution = Minute * 5
	Minute15 TimeResolution = Minute * 15
	Hour     TimeResolution = Minute * 60
	Day      TimeResolution = Hour * 24
)

func (tr TimeResolution) String() string {
	switch tr {
	case sec:
		return "sec"
	case Minute:
		return "min"
	case Minute5:
		return "5 min"
	case Minute15:
		return "15 min"
	case Hour:
		return "hour"
	case Day:
		return "day"
	}
	panic("no other options")
}
