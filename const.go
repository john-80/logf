package logf

type LEVEL int

const (
	PanicLevel LEVEL = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

type PREFIX string

const (
	FatalFrefix PREFIX = "[FATAL] "
	ErrorPrefix PREFIX = "[ERROR] "
	WarnPrefix  PREFIX = "[WARN] "
	InfoPrefix  PREFIX = "[INFO] "
	DebugPrefix PREFIX = "[DEBUG] "
)

type MODE int

const (
	DayMode MODE = iota
	WeekMode
	MonthMode
	SizeMode
)
