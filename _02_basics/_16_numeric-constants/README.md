<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Declare constants Big and Small];
B --> C[Declare function needInt];
C --> D[Declare function needFloat];
D --> E["Print needInt(Small)"];
E --> F["Print needFloat(Small)"];
F --> G["Print needFloat(Big)"];
</div>

* Numeric Constants
Numeric constants are high-precision _values_.

An untyped constant takes the type needed by its context.

Try printing `needInt(Big)` too.

(An `int` can store at maximum a 64-bit integer, and sometimes less.)

<a onclick="nextOpen();confetti()">next</a>
