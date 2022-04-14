package logger

type (
	Logger interface {
		Info(field LogField)
		Warn(field LogField)
		Error(field LogField)
	}

	LogField struct {
		Section  string
		Function string
		Params   any
		Message  string
	}
)
