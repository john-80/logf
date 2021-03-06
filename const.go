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

const (
	FatalFrefix string = "[FATAL] "
	ErrorPrefix string = "[ERROR] "
	WarnPrefix  string = "[WARN] "
	InfoPrefix  string = "[INFO] "
	DebugPrefix string = "[DEBUG] "
)

type MODE int

const (
	DayMode MODE = iota
	WeekMode
	MonthMode
	SizeMode
)
