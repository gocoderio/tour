<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Declare function sqrt];
B --> C[Check if x is negative];
C -- Yes --> D["Call sqrt with -x and append i"];
C -- No --> E[Return square root of x];
E --> F["Print sqrt(2)"];
E --> G["Print sqrt(-4)"];
</div>

# If

Go's `if` statements are like its `for` loops; the expression need not be
surrounded by parentheses `(`)` but the braces `{`}` are required.

<a onclick="nextOpen()">next lesson</a>
