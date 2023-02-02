Packages, variables, and functions.
Learn the basic components of any Go program.

The Go Authors
https://golang.org

* Packages

Every Go program is made up of packages.

Programs start running in package `main`.

This program is using the packages with import paths `"fmt"` and `"math/rand"`.

By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement `package`rand`.

#appengine: *Note:* The environment in which these programs are executed is
#appengine: deterministic, so each time you run the example program
#appengine: `rand.Intn` will return the same number.
#appengine:
#appengine: (To see a different number, seed the number generator; see [[/pkg/math/rand/#Seed][`rand.Seed`]].
#appengine: Time is constant in the playground, so you will need to use something else as the seed.)

