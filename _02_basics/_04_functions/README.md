<div id="chart" class="mermaid">
graph TD;
    A[Start] --> B["Print add(42, 13)"]
    B --> C[Call add function]
    C --> D[Return value of add function]
</div>

# Functions
A function can take zero or more arguments.

In this example, `add` takes two parameters of type `int`.

Notice that the type comes _after_ the variable name.

(For more about why types look the way they do, see the [[https://blog.golang.org/gos-declaration-syntax][article on Go's declaration syntax]].)

<a onclick="nextOpen()">next lesson</a>
