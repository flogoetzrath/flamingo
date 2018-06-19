package requestlogger

import (
	"flamingo.me/flamingo/framework/dingo"
	"flamingo.me/flamingo/framework/router"
)

type (
	// Module for core/requestlogger
	Module struct{}

	// LogFilter to filter web requests
	LogFilter struct{}
)

// Configure DI
func (m *Module) Configure(injector *dingo.Injector) {
	injector.BindMulti((*router.Filter)(nil)).To(logger{})
}
