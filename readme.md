# array methods

This is me practicing generics and testing in Go

```bash
go get -u github.com/tsivinsky/array
```

## array.Each

```go
array.Each[T any](slice []T, func (item T, i int))
```

## array.Map

```go
array.Map[T any, M any](slice []T, func(item T, i int) M) []M
```

## array.Reduce

```go
array.Reduce[T any, M any](slice []T, func(accumulator M, currentValue T) M, initialValue M) M
```

## array.Some

```go
array.Some[T any](slice []T, func(item T, i int) bool) bool
```

## array.Every

```go
array.Every[T any](slice []T, func(item T, i int) bool) bool
```

## array.Find

```go
array.Find[T any](slice []T, func (item T, i int) bool) T
```

## array.FindIndex

```go
array.FindIndex[T any](slice []T, func (item T, i int) bool) int
```

## array.Filter

```go
array.Filter[T any](slice []T, func (item T, i int) bool) []T
```

## array.Includes

```go
array.Includes[T comparable](slice []T, item T) bool
```

## array.At

Try passing negative number as second argument

```go
array.At[T any](s []T, i int) T
```
