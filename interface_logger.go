package go_interface_logger

type InterfaceLogger interface {
	Close()

	Debug(args ...interface{})
	DebugF(format string, args ...interface{})
	Info(args ...interface{})
	InfoF(format string, args ...interface{})
	Warn(args ...interface{})
	WarnF(format string, args ...interface{})
	Error(args ...interface{})
	ErrorF(format string, args ...interface{})
	ErrorWithStack(args ...interface{})
	ErrorWithStackF(format string, args ...interface{})
}

