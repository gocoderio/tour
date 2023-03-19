<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Declare variable sum];
B --> C[While loop];
C --> D[Add sum to itself];
D --> E[Check if sum is less than 1000];
E -- Yes --> C;
E -- No --> F[Print sum];
</div>

# For is Go's "while"

At that point you can drop the semicolons: C's `while` is spelled `for` in Go.

<a onclick="nextOpen()">next lesson</a>
