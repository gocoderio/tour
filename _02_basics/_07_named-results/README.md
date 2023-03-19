<div id="chart" class="mermaid">
graph TD;
    A[Start] --> B["split(sum int)"];
    B --> C["x = sum * 4 / 9"];
    B --> D["y = sum - x"];
    C --> E[x];
    D --> F[y];
    E --> G[Return x];
    F --> H[Return y];
</div>

# Named return values
Go's return values may be named. If so, they are treated as variables defined at the top of the function.

These names should be used to document the meaning of the return values.

A `return` statement without arguments returns the named return values. This is known as a "naked" return.

Naked return statements should be used only in short functions, as with the example shown here. They can harm readability in longer functions.

<a onclick="nextOpen()">next lesson</a>
