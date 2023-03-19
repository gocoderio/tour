<div id="chart" class="mermaid">
graph TD;
A[Start] --> B[Print counting];
B --> C[For loop];
C --> D[Defer Print i];
D --> C;
C --> E[Print done];
</div>

# Stacking defers

Deferred function calls are pushed onto a stack. When a function returns, its
deferred calls are executed in last-in-first-out order.

To learn more about defer statements read this
[[/blog/defer-panic-and-recover][blog post]].

--------------------

* Congratulations!

You finished this lesson!

You can go back to the list of [[/tour/list][modules]] to find what to learn next, or continue with the [[javascript:click('.next-page')][next lesson]].

<a onclick="nextOpen()">next lesson</a>
