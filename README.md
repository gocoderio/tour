# Hello, from web assembly
```mermaid
graph TD
    A[main]:::lineNumber=5 --> B[Inception]:::lineNumber=7
    B --> C[Level 1: Hello from inside Inception]:::lineNumber=11
    B --> D[level2]:::lineNumber=12
    D --> E[Level 2: Hello from inside Level 2]:::lineNumber=16
    D --> F[level3]:::lineNumber=17
    F --> G[Level 3: Hello from inside Level 3]:::lineNumber=21
```
<p style="padding-left:20px;padding-right:20px;font-size:small">
Web assembly is a browser technology that lets you create and run applications in your own web browser. This is up to 10x faster than cloud based programming (sending your code to the cloud and then waiting for the results).
</p>
<br>
<p style="padding-left:10px;padding-right:10px;">
&nbsp;&nbsp;&nbsp;&nbsp;<span>Run the complete <b>'Tour of Go'</b> in your web browser. You can turn off your internet and see that it still works by clicking the <a onclick="highlightAndClick('#runButton')">run</a> button, then click to the next lesson to start with a simple 'hello world' program.<span>
</p>
<a href="https://go.dev/tour/welcome/1" target="_blank">Based on the original 'Tour of Go', by the Go Authors</a>
<div class="nextLink"><a onclick="nextOpen()">next lesson</a></div>
