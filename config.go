package tenkiGetter

type Config struct {
	RunMode    Mode
	OfficeCode string
}

type Mode int

const (
	Day Mode = iota + 1
	Week
)

func NewConfig(mode Mode) *Config {
	return &Config{RunMode: mode}
}

func (m Mode) String() string {
	switch m {
	case Day:
		return "day"
	case Week:
		return "week"
	default:
		return "unknown"
	}
}

func (m Mode) GetMode() string {
	switch m {
	case Day:
		return "overview_forecast"
	case Week:
		return "overview_week"
	default:
		return "unknown"
	}
}
