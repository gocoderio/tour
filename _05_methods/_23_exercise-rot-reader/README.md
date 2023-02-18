* Exercise: rot13Reader

A common pattern is an [[/pkg/io/#Reader][io.Reader]] that wraps another `io.Reader`, modifying the stream in some way.

For example, the [[/pkg/compress/gzip/#NewReader][gzip.NewReader]] function takes an `io.Reader` (a stream of compressed data) and returns a `*gzip.Reader` that also implements `io.Reader` (a stream of the decompressed data).

Implement a `rot13Reader` that implements `io.Reader` and reads from an `io.Reader`, modifying the stream by applying the [[https://en.wikipedia.org/wiki/ROT13][rot13]] substitution cipher to all alphabetical characters.

The `rot13Reader` type is provided for you.
Make it an `io.Reader` by implementing its `Read` method.
