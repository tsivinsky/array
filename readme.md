# array methods

This is me practicing generics and testing in Go

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
