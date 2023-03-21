<div id="chart" class="mermaid" style="float: right; width: 50%; margin-left: 10px; margin-bottom: 10px;">
flowchart TD
    A(Print message)
    B(Get current weekday)
    C(Compare weekday to Saturday)
    D("Print Today.")
    E("Print Tomorrow.")
    F("Print In two days.")
    G("Print Too far away.")
    A --> B --> C
    C -- today + 0 --> D
    C -- today + 1 --> E
    C -- today + 2 --> F
    C -- default --> G
</div>

# Switch evaluation order
Switch cases evaluate cases from top to bottom, stopping when a case succeeds.

(For example,

	switch i {
	case 0:
	case f():
	}

does not call `f` if `i==0`.)

#appengine: *Note:* Time in the Go playground always appears to start at
#appengine: 2009-11-10 23:00:00 UTC, a value whose significance is left as an
#appengine: exercise for the reader.

<div class="nextLink"><a onclick="nextOpen()">next lesson</a></div>
