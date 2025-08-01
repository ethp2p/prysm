// Package x contains interceptors that are invoked at different points in the
// runtime flow of Prysm. Interceptors are represented as public global
// variables that Afferent probes can override to inject themselves.
// The correct usage is to initialize these variables _before_ the Prysm
// instance is started, or they'll have no effect. The default logic for all
// interceptors is no-op, i.e. Prysm will run as it would upstream.
package x

func Identity[T any](o T) T {
	return o
}
