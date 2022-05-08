# gollections

[![MIT License](https://img.shields.io/github/license/bdreece/gollections)](https://github.com/bdreece/gollections/blob/main/LICENSE.md)
[![GitHub Releases](https://img.shields.io/github/v/release/bdreece/gollections?include_prereleases)](https://github.com/bdreece/gollections/releases)
[![Build Status](https://img.shields.io/github/workflow/status/bdreece/gollections/Go)](https://github.com/bdreece/gollections/actions/workflows/go.yml)

# (WIP) WORK IN PROGRESS!

Common data structures and interfaces, written in Go

## Table of Contents

- [Overview](#overview)
    - [Common Interfaces](#common-interfaces)
    - [Package Hierarchy](#package-hierarchy)
- [Subpackages](#subpackages)
    - [vector](#vector)
    - [ringbuf](#ringbuf)
    - [list](#list)
    - [iterator](#iterator)
    - [maps](#maps)
    - [trees](#trees)
    - [graphs](#graphs)
    - [errors](#errors)
- [API Reference](https://pkg.go.dev/github.com/bdreece/gollections)
- [Future Plans](#future-plans)

## Overview

`gollections` provides common data structures and interfaces for Go 1.18. The
most common interfaces exist in the root `gollections` package, and more specific interfaces will be describes in subpackages.

### Common Interfaces

Common interfaces provided by `gollections` include:

- Array : Array-like structures (i.e. indexable collection)
- Stack : Stack operations (e.g. push, pop, peek)
- Queue : Queue operations (e.g. enqueue, dequeue)
- Deque : Double-ended queue operations (e.g. push back/front, pop back/front)

The root `gollections` package also provides the Collect interface, which is implemented by all collections.

### Package Hierarchy

The root `gollections` package only contains the interface that different concrete implementations adhere to. In order to instantiate a data structure, one must be selected from the subpackages under `github.com/bdreece/gollections/...`. Some subpackages (e.g. maps, trees, etc.) contain more specific interfaces as well as further nested subpackages. These follow the same structure as the root packages.

## Subpackages

### vector

The `vector` package provides the vector data structure: a dynamic array-like data structure stored in contiguous memory. Vector objects implement the Array, Queue, Deque, Stack, and Collect interfaces.

### ringbuf

The `ringbuf` package provides the ring buffer data structure: a queue-like data structure stored in fixed-size contiquous memory. Ring buffer objects implement the Queue and Collect interfaces.

### list

The `list` package provides the linked-list data structure: an array-like data structure with links to previous and next elements at each element. List objects implement the Array, Queue, Deque, Stack, and Collect interfaces

### iterator

The `iterator` package provides a number of interfaces related to iteration over collections. The primary interface, Iterator, may be utilized by iterator transformation functions, such as `ForEach`, `Filter`, `Map`, `Enumerate`, etc. Collections may be used to create Iterators, by implementing the Iterable interface (provides `IntoIterator()` and `FromIterator()`).

### maps

The `maps` package provides the generic Map interface, as well as some concrete implementors of this interface. Maps extend the Array and Collect interfaces with more specific functionality.

### trees

The `trees` package provides the generic Tree, Node, and Leaf interfaces, as well as some concrete implementors of this interface. Trees are only guaranteed to implement the Collect interface, but concrete trees may implement Array or Map as well.

### graphs

The `graphs` package provides the generic Graph, Vertex, and Edge interfaces, as well as some concrete implementors of this interface. Graphs are only guaranteed to implement the Collect interface, but concrete graphs may implement other interfaces as well.

### errors

The `errors` package provides a number of common data structure operation errors, such as "Index out of bounds", or "Empty".

## Future Plans

I need to write more tests, and expand the different interfaces and data structures.