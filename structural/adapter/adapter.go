package adapter

// Adapter: Converts OldLogger to NewLogger interface
type LoggerAdapter struct {
	oldLogger *OldLogger
}

func NewLoggerAdapter(oldLogger *OldLogger) NewLogger {
	return &LoggerAdapter{oldLogger: oldLogger}
}

func (a *LoggerAdapter) Log(level, message string) {
	adaptedMessage := "[" + level + "] " + message
	a.oldLogger.WriteLog(adaptedMessage)
}
