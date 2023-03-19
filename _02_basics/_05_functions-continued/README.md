<div id="chart" class="mermaid">
graph TD;
    A[Start] --> B["Print add(42, 13)"]
    B --> C[Call add function]
    C --> D[Return value of add function]
</div>

# Functions continued
When two or more consecutive named function parameters share a type, you can omit the type from all but the last.

In this example, we shortened

	x int, y int

to

	x, y int

<a onclick="nextOpen()">next lesson</a>
