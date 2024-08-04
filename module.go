package fxlogger

import (
	"io"
	"os"
	 //config "gotemplate/config"
	 config "github.com/templatedop/fxconfig"
	"github.com/templatedop/fxlogger/fxloggertest"

	"github.com/rs/zerolog"
	"go.uber.org/fx"
)

var FxLoggerModule = fx.Module("logger",
	fx.Provide(
		NewDefaultLoggerFactory,
		NewFxLogger,
		fx.Annotate(
			fxloggertest.GetTestLogBufferInstance,
			fx.ResultTags(`name:"test-log-buffer"`),
		),
	),
)

type FxLoggerParam struct {
	fx.In
	Factory LoggerFactory
	Config  config.Econfig
	//Config  *fxconfig.EConfig
}

func NewFxLogger(p FxLoggerParam) (*Logger, error) {
	// level
	level := FetchLogLevel(p.Config.LogLevel)
	if p.Config.LogLevel=="debug" {
		level = zerolog.DebugLevel
	}	

	// output writer
	var outputWriter io.Writer
	if p.Config.AppEnv == "test" {
		outputWriter = fxloggertest.GetTestLogBufferInstance()
	} else {
		outputWriter = os.Stdout
	}

	// logger
	return p.Factory.Create(
		WithName(p.Config.AppName),		
		WithLevel(level),
		WithOutputWriter(outputWriter),
	)
}
