#appengine: * The Go Playground
#appengine:
#appengine: This tour is built atop the [[https://play.golang.org/][Go Playground]], a
#appengine: web service that runs on [[/][golang.org]]'s servers.
#appengine:
#appengine: The service receives a Go program, compiles, links, and runs the program inside
#appengine: a sandbox, then returns the output.
#appengine:
#appengine: There are limitations to the programs that can be run in the playground:
#appengine:
#appengine: - In the playground the time begins at 2009-11-10 23:00:00 UTC (determining the significance of this date is an exercise for the reader). This makes it easier to cache programs by giving them deterministic output.
#appengine:
#appengine: - There are also limits on execution time and on CPU and memory usage, and the program cannot access external network hosts.
#appengine:
#appengine: The playground uses the latest stable release of Go.
#appengine:
#appengine: Read "[[/blog/playground][Inside the Go Playground]]" to learn more.
#appengine:
#appengine: .play welcome/sandbox.go
