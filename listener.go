package logger

type LogListener func(level Level, params ...interface{})

var listeners []LogListener

func AddLogListener(listener LogListener) {
	listeners = append(listeners, listener)
}
