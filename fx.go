package fxlogger



import (
	"fmt"
	"strings"

	"go.uber.org/fx/fxevent"
)

var FxEventLogger = func(log *Logger) fxevent.Logger {
	return log
}

func (l *Logger) LogEvent(event fxevent.Event) {
	//fmt.Println("Event received...",event)
	//l.ToZerolog().Debug().Interface("event", event).Msg("event received")
	switch e := event.(type) {
	case *fxevent.OnStartExecuting:
		l.ToZerolog().Debug().Str("callee", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStart hook executing")
	case *fxevent.OnStartExecuted:
		if e.Err != nil {
			l.ToZerolog().Error().Err(e.Err).
				Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Caller().
				//Str("",e.Runtime.).
				Msg("OnStart hook failed")
		} else {
			l.ToZerolog().Debug().Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStart hook executed")
		}
	case *fxevent.OnStopExecuting:
		l.ToZerolog().Debug().Str("callee", e.FunctionName).
			Str("caller", e.CallerName).
			Msg("OnStop hook executing")
	case *fxevent.OnStopExecuted:
		if e.Err != nil {
			l.ToZerolog().Warn().Err(e.Err).
				Str("callee", e.FunctionName).
				Str("callee", e.CallerName).
				Msg("OnStop hook failed")
		} else {
			l.ToZerolog().Debug().Str("callee", e.FunctionName).
				Str("caller", e.CallerName).
				Str("runtime", e.Runtime.String()).
				Msg("OnStop hook executed")
		}
	case *fxevent.Supplied:
		if e.Err != nil {
			l.ToZerolog().Warn().Err(e.Err).Str("type", e.TypeName).Msg("supplied")
		} else {

			l.ToZerolog().Debug().Str("type", e.TypeName).Msg("supplied")
		}
	case *fxevent.Provided:
		for _, rtype := range e.OutputTypeNames {
			l.ToZerolog().Debug().Str("type", rtype).
				//Caller().
				Str("constructor", e.ConstructorName).
				Str("module", e.ModuleName).
				Msg("provided")
		}
		if e.Err != nil {
			l.ToZerolog().Error().Err(e.Err).Str("module", e.ModuleName).Msg("error encountered while applying options")
		}
	case *fxevent.Invoking:
	case *fxevent.Invoked:
		if e.Err != nil {
fmt.Println("Error in invoked...",e.Err)
			l.ToZerolog().Error().
				Err(e.Err).Str("stack", e.Trace).
				Str("function", e.FunctionName).Msg("invoke failed")
		} else {
			l.ToZerolog().Debug().Str("function", e.FunctionName).Msg("invoked")
		}
	case *fxevent.Stopping:
		l.ToZerolog().Debug().Str("signal", strings.ToUpper(e.Signal.String())).Msg("received signal")
	case *fxevent.Stopped:
		if e.Err != nil {
			l.ToZerolog().Error().Err(e.Err).Msg("stop failed")
		}
	case *fxevent.RollingBack:
		l.ToZerolog().Error().Err(e.StartErr).Msg("start failed, rolling back")
	case *fxevent.RolledBack:
		if e.Err != nil {
			l.ToZerolog().Error().Err(e.Err).Msg("rollback failed")
		}
	case *fxevent.Started:
		if e.Err != nil {
			l.ToZerolog().Error().Err(e.Err).Msg("start failed")
		} else {
			l.ToZerolog().Debug().Msg("started")
		}
	case *fxevent.LoggerInitialized:
		if e.Err != nil {
			l.ToZerolog().Debug().Err(e.Err).Msg("custom logger initialization failed")
		} else {
			l.ToZerolog().Debug().Str("function", e.ConstructorName).Msg("initialized custom fxevent.Logger")
		}
	}
}
