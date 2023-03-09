<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Infinite loop];
B --> C[Sleep for 1 second];
C --> B;
</div>

# Forever

If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
