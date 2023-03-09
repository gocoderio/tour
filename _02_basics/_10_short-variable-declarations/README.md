<div id="chart" class="mermaid">
graph TD;
    A[Start] --> B[Declare i and j as integers];
    B --> C[Declare k as integer and assign value 3];
    C --> D[Declare c, python, and java as booleans and string];
    D --> E[Assign values to c, python, and java];
    E --> F[Print i, j, k, c, python, and java];
</div>

# Short variable declarations
Inside a function, the `:=` short assignment statement can be used in place of a `var` declaration with implicit type.

Outside a function, every statement begins with a keyword (`var`, `func`, and so on) and so the `:=` construct is not available.
