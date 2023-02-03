* Switch evaluation order

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

.play flowcontrol/switch-evaluation-order.go
