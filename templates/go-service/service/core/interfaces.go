package core

// SetupComponent function to setup component
type SetupComponent func()

// App interface
type App interface {
	Liveness() bool
	Readiness() bool
	SetLiveness(liveness bool)
	SetReadiness(readiness bool)
	Setup(componentSetups []SetupComponent)
	Run()
}
