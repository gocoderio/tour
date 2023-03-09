<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Declare function pow];
B --> C["Calculate x^n"];
C --> D["Check if x^n is less than lim"];
D -- Yes --> E["Return x^n"];
D -- No --> F["Print v >= lim"];
F --> G[Return lim];
G --> H["Print pow(3, 2, 10) and pow(3, 3, 20)"];
</div>

# If and else

Variables declared inside an `if` short statement are also available inside any
of the `else` blocks.

(Both calls to `pow` return their results before the call to `fmt.Println`
in `main` begins.)
