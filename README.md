# go-jsontime

A slim `time.Time` wrapper that handles pointer to time and formats the encoded time as [`RFC3339/ISO8601`](https://github.com/golang/go/blob/master/src/time/format.go#L82), which is easy to format and read from.

## Usage

Either use `jsontime.JSONTime` as a pointer or the object itself, then just use it with the default `json` package.

```go
type testJSON struct {
	Test JSONTime `json:"test"`
}

type testPtrJSON struct {
	Test *JSONTime `json:"test"`
}

```

For a more detailed usage example please have a look at the [tests](json_time_test.go).

## License

```
MIT License

Copyright (c) 2019 Alexander Pinnecke
```
