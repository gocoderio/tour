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
This tour is built in web assembly (In Go!), and runs inside your web browser.
Go compiled to WASM runs your go program that you are writing now. 


- In the playground the time is accurate if your computer time is accurate.

- There are no limits on memory usage, except for your hardware limits. The program can access external network hosts.

- Networking and Cross Origin API calls work with PRO accounts.

The web assembly playground uses the latest stable release of Go.

<div class="nextLink"><a onclick="nextOpen()">next lesson</a></div>
