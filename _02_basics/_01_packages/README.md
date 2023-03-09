<div id="chart" class="mermaid">
graph TD;
    A[Start] --> B["Print My favorite number is"]
    B --> C[Generate random number]
    C --> D[Print random number]
</div>

# Packages

Every Go program is made up of packages.

Programs start running in package `main`.

This program is using the packages with import paths `"fmt"` and `"math/rand"`.

By convention, the package name is the same as the last element of the import path. For instance, the `"math/rand"` package comprises files that begin with the statement `package`rand`.


