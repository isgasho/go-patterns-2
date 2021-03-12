package main

// This pattern is based on the OOP example where class A depends on class B. We add an instance of B as a parameter to the
// constructor so we can use the interface of B to provide a test fake.
// Thanks to @mhmxs for the suggestion.

// region Factories

// NewA creates a new copy of A. It expects B as a dependency.
func NewA(b B) A {
	return &a{
		b: b,
	}
}

func NewB() B {
	return &b{}
}

// endregion

// region Interfaces
type B interface {
	DoSomething()
}

type A interface {
	DoSomethingElse()
}
// endregion


// region Implementation
type b struct {

}

func (b b) DoSomething() {
	println("Hello world!")
}

type a struct {
	// The b variable will contain the necessary dependency.
	b B
}

func (a *a) DoSomethingElse() {
	a.b.DoSomething()
}
// endregion

func main() {
	b := NewB()
	a := NewA(b)
	a.DoSomethingElse()
}