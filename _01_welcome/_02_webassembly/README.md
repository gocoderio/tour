<div id="chart" class="mermaid">
graph TD
    A[Start]:::lineNumber=1
    B[Import packages]:::lineNumber=3
    C["Print Welcome to the playground!"]:::lineNumber=9
    D[Get and print current time]:::lineNumber=11
    A --> B
    B --> C
    C --> D
</div>

# The Go Playground
This tour is built atop the web assembly, it runs inside your web browser so your code stays private.

There are limitations to the programs that can be run in a web browser:

- In the playground the time is accurate if your computer time is accurate.

- There are no limits on execution time or CPU and memory usage, except for your hardware limits. The program can access external network hosts.

- Networking and Cross Origin API calls have been tested, but limited to what browser security allows.

The is web assembly playground uses the latest stable release of Go.

Read `Inside the GoCoder Webassembly` to learn more.
<a onclick="nextOpen();confetti()">next lesson</a>

