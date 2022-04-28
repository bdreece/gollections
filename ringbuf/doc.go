// MIT License
// Copyright (c) 2022 Brian Reece

/*
ringbuf provides the RingBuf data structure.

A RingBuf is a slice of contiguous data that
acts as a fixed-length queue. Writes exceeding
the capacity of the RingBuf will wrap around to
the front, and begin overwriting values.
*/
package ringbuf
