// Copyright (c) 2012-2016 The Revel Framework Authors, All rights reserved.
// Revel Framework source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package revel

// Filter type definition for Revel's filter.
type Filter func(c *Controller, filterChain []Filter)

// Filters is the default set of global filters.
// It may be set by the application on initialization.
var Filters = []Filter{
	PanicFilter,             // Recover from panics and display an error page instead.
	RouterFilter,            // Use the routing table to select the right Action.
	FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
	ParamsFilter,            // Parse parameters into Controller.Params.
	SessionFilter,           // Restore and write the session cookie.
	FlashFilter,             // Restore and write the flash cookie.
	ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
	I18nFilter,              // Resolve the requested language.
	InterceptorFilter,       // Run interceptors around the action.
	CompressFilter,          // Compress the result.
	BeforeAfterFilter,
	ActionInvoker, // Invoke the action.
}

// NilFilter and NilChain are helpful in writing filter tests.
var (
	NilFilter = func(_ *Controller, _ []Filter) {}
	NilChain  = []Filter{NilFilter}
)
