package tenkiGetter

type Config struct {
	RunMode Mode
}

type Mode int

const (
	day Mode = iota + 1
	week
)

func NewConfig(mode Mode) *Config {
	return &Config{RunMode: mode}
}

func (m Mode) String() string {
	switch m {
	case day:
		return "day"
	case week:
		return "week"
	default:
		return "unknown"
	}
}
