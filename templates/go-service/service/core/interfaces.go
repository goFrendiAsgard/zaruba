package core

// Comp interface
type Comp interface {
	Setup()
}

// App interface
type App interface {
	Liveness() bool
	Readiness() bool
	SetLiveness(liveness bool)
	SetReadiness(readiness bool)
	Setup(componentSetups []Comp)
	Run()
}
