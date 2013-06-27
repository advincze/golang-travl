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
		return "5min"
	case Minute15:
		return "15min"
	case Hour:
		return "hour"
	case Day:
		return "day"
	}
	panic("no other options")
}

func ParseTimeResolution(s string) TimeResolution {
	switch s {
	case "sec":
		return sec
	case "min":
		return Minute
	case "5min":
		return Minute5
	case "15min":
		return Minute15
	case "hour":
		return Hour
	case "day":
		return Day
	}
	panic("no other options")
}
