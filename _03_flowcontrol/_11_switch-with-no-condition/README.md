<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Get current time];
B --> C["Switch on time.Hour()"];
C -- Hour < 12 --> D["Print Good morning!"];
C -- Hour < 17 --> E["Print Good afternoon."];
C -- default --> F["Print Good evening."];
</div>

# Switch with no condition

Switch without a condition is the same as `switch`true`.

This construct can be a clean way to write long if-then-else chains.

