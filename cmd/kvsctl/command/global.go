package command

type GlobalFlags struct {
	Host      string
	Port      int
	DebugMode bool
}

func GetLogLevel(debugMode bool) string {
	if debugMode {
		return "debug"
	}
	return "info"
}
