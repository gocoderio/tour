<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Declare function pow];
B --> C["Calculate x^n"];
C --> D["Check if x^n is less than lim"];
D -- Yes --> E["Return x^n"];
D -- No --> F[Return lim];
F --> G["Print pow(3, 2, 10) and pow(3, 3, 20)"];
</div>

# If with a short statement

Like `for`, the `if` statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in scope until the end of the `if`.

(Try using `v` in the last `return` statement.)

<a onclick="nextOpen()">next lesson</a>
