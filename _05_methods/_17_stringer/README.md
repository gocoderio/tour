* Stringers

One of the most ubiquitous interfaces is [[/pkg/fmt/#Stringer][`Stringer`]] defined by the [[/pkg/fmt/][`fmt`]] package.

	type Stringer interface {
		String() string
	}

A `Stringer` is a type that can describe itself as a string. The `fmt` package
(and many others) look for this interface to print values.
