<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Print Go runs on];
B --> C[Switch on runtime.GOOS];
C -- darwin --> D[Print OS X];
C -- linux --> E[Print Linux];
C -- js --> F["Print js !!!"];
C -- default --> G[Print os];
</div>

# Switch

A `switch` statement is a shorter way to write a sequence of `if`-`else` statements.
It runs the first case whose value is equal to the condition expression.

Go's switch is like the one in C, C++, Java, JavaScript, and PHP,
except that Go only runs the selected case, not all the cases that follow.
In effect, the `break` statement that is needed at the end of each case in those
languages is provided automatically in Go.
Another important difference is that Go's switch cases need not
be constants, and the values involved need not be integers.

<a onclick="nextOpen()">next lesson</a>
