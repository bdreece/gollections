# gollections

[![Golang](https://img.shields.io/badge/-Go-00ADD8?logo=go&logoColor=white&style=for-the-badge)](https://go.dev/)
[![MIT License](https://img.shields.io/github/license/bdreece/gollections?style=for-the-badge)](https://github.com/bdreece/gollections/blob/main/LICENSE.md)
[![CI Status](https://img.shields.io/github/actions/workflow/status/bdreece/gollections/go.yml?style=for-the-badge)](https://github.com/bdreece/gollections/actions/workflows/go.yml)

A collection of collections!

## Table of Contents

- [Overview](#overview)
  - [Usage](#usage)
  - [Go Type Interop](#go-type-interop)
- [Packages](#packages)
  - [`pkg/collection`](#collection)
  - [`pkg/dll`](#doubly-linked-list)
  - [`pkg/hashmap`](#hash-map)
  - [`pkg/iterator`](#iterator)
  - [`pkg/list`](#list)
  - [`pkg/queue`](#queue)
  - [`pkg/slice`](#slice)
  - [`pkg/sll`](#singly-linked-list)
  - [`pkg/stack`](#stack)

## Overview

[gollections](#gollections) is a collection of generic data structures, written with Go 1.18.
The purpose of this package is to provide a flexible system of generic collections and lazily-
executed iterators for use in a variety of applications

### Usage

Each of the collections resides in its own module, located under the [`pkg/`](pkg/) directory.
Currently, only the following concrete data structures are implemented:

- [`pkg/dll`](#dll): Doubly-linked list
- [`pkg/hashmap`](#hashmap): A basic hash map
- [`pkg/slice`](#slice): Go slice wrapper
- [`pkg/sll`](#sll): Singly-linked list

The remaining packages in this repository provide interfaces dedicated to more specific data structures (e.g. [`pkg/stack`](#stack), [`pkg/queue`](#queue), etc.), and iterators (see [`pkg/iterator`](#iterator))

### Go Type Interop

Go slices are used as the underlying type for the [`pkg/slice`](#slice) collection, and can be
coerced as follows:

```go
package main

import "github.com/bdreece/gollections/pkg/slice"

func main() {
    s := slice.From([]int{1, 2, 3})

    /*
        s.Get(2)
        v := s.First()
        ...
    */

}
```

## Packages

### Collection

The [`pkg/collection`](pkg/collection/) module provides the base interface for all
collection interfaces in this repository.

A `collection.Collection[TItem]` provides the following methods:

- `Concat(iterator.IntoIterator[TItem]) Collection[TItem]`
- `Collect(iterator.Iterator[TItem]) Collection[TItem]`
- `Append(TItem) Collection[TItem]`
- `Count() int`

The `collection.Collection[TItem]` interface also implements the [`iterator.IntoIterator[TItem]`](#iterator) interface.

### Doubly-Linked List

The [`pkg/dll`](pkg/dll/) module provides the `dll.DLL[TItem]` interface, which is implemented as a doubly-linked list of nodes.

A `dll.DLL[TItem]` implements the following interfaces:

- [`stack.Stack[TItem]`](#stack)
- [`queue.Queue[TItem]`](#queue)
- [`collection.Collection[TItem]`](#collection)
- [`iterator.IntoIterator[TItem]`](#iterator)

### Hash Map

The [`pkg/hashmap`](pkg/hashmap/) module provides the `hashmap.HashMap[TKey, TValue]` interface, which is implemented as a hash map data structure.

A `hashmap.HashMap[TKey, TValue]` implements the following interfaces:

- `collection.Collection[hashmap.Pair[TKey, TValue]]`
- `iterator.IntoIterator[hashmap.Pair[TKey, TValue]]`

A `hashmap.HashMap[TKey, TValue]` also provides the following methods:

- `Get(TKey) (*TValue, error)`
- `Set(TKey, TValue) error`
- `Remove(TKey) (*TValue, error)`

### Iterator

The [`pkg/iterator`](pkg/iterator/) module provides iterator interfaces and related functions
for lazily-iterating over collections.

The following functions can be used to operate over an `iterator.Iterator[TItem]`:

- `Chain[TItem](iterator.Iterator[TItem], iterator.Iterator[TItem]) iterator.Iterator[TItem]`
- `Empty[TItem]() iterator.Iterator[TItem]`
- `Enumerate[TItem](iterator.Iterator[TItem], EnumerateFunc[TItem])`
  - `type EnumerateFunc[TItem] func(TItem, int)`
- `Filter[TItem](iterator.Iterator[TItem], FilterFunc[TItem])`
  - `type FilterFunc[TItem] func(TItem) bool`
- `Find[TItem](iterator.Iterator[TItem], FindFunc[TItem])`
  - `type FindFunc[TItem] func(TItem) bool`
- `FlatMap[TInput, TOutput](iterator.Iterator[TInput], FlatMapFunc[TInput, TOutput])`
  - `type FlatMapFunc[TInput, TOutput] func(TInput) iterator.IntoIterator[TOutput]`
- `ForEach[TItem](iterator.Iterator, ForEachFunc[TItem])`
  - `type ForEachFunc[TItem] func(TItem)`
- `Map[TInput, TOutput](iterator.Iterator[TInput], MapFunc[TInput, TOutput])`
  - `type MapFunc[TInput, TOutput] func(TInput) TOutput`
- `Reduce[TItem, TAggregate](iterator.Iterator[TItem], ReduceFunc[TItem, TAggregate], TAggregate)`
  - `type ReduceFunc[TItem, TAggregate] func(TAggregate, TItem) TAggregate`
- `Take[TItem](iterator.Iterator[TItem], int) iterator.Iterator[TItem]`

The [`pkg/iterator`](pkg/iterator/) module also provides the `iterator.IntoIterator[TItem]`
interface, which provides the following method:

- `Iter[TItem]() iterator.Iterator[TItem]`

### List

The [`pkg/list`](pkg/list) module provides the `list.List[TItem]` interface, which represents a basic enumerable list.

A `list.List[TItem]` implements the following interfaces:

- [`slice.Slice[TItem]`](#slice)
- [`collection.Collection[TItem]`](#collection)
- [`iterator.IntoIterator[TItem]`](#iterator)

### Queue

The [`pkg/queue`](pkg/queue) module provides the `queue.Queue[TItem]` interface, which represents a FIFO buffer.

A `queue.Queue[TItem]` provides the following methods:

- `Enqueue(TItem)`
- `Dequeue() *TItem`

A `queue.Queue[TItem]` also implements the following interfaces:

- [`collection.Collection[TItem]`](#collection)
- [`iterator.IntoIterator[TItem]`](#iterator)

### Slice

The [`pkg/slice`](pkg/slice) module provides the `slice.Slice[TItem]` interface, which is implemented as a thin wrapper around the primitive Go slice.

A `slice.Slice[TItem]` provides the following methods:

- `First() *TItem`
- `Last() *TItem`
- `Add(TItem)`
- `Get(int) (*TItem, error)`
- `Set(int, TItem) error`
- `Insert(int, TItem) error`
- `Remove(int) (*TItem, error)`

A `slice.Slice[TItem]` also implements the following interfaces:

- [`collection.Collection[TItem]`](#collection)
- [`iterator.IntoIterator[TItem]`](#iterator)

### Singly-Linked List

The [`pkg/sll`](pkg/sll) module provides the `sll.SLL[TItem]` interface, which is implemented as a singly-linked list of nodes.

A `sll.SLL[TItem]` implements the following interfaces:

- [`stack.Stack[TItem]`](#stack)
- [`queue.Queue[TItem]`](#queue)
- [`collection.Collection[TItem]`](#collection)
- [`iterator.IntoIterator[TItem]`](#iterator)

### Stack

The [`pkg/stack`](pkg/stack) modules provides the `stack.Stack[TItem]` interface, which represents a FILO buffer

A `stack.Stack[TItem]` provides the following methods:

- `Push(TItem)`
- `Pop() *TItem`
- `Peek() *TItem`

A `stack.Stack[TItem]` also implements the following interfaces:

- [`collection.Collection[TItem]`](#collection)
- [`iterator.IntoIterator[TItem]`](#iterator)
