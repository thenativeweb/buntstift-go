# buntstift-go

buntstift-go makes the CLI colorful.

![buntstift](https://github.com/thenativeweb/buntstift-go/raw/master/images/logo.jpg "buntstift")

## Installation

    $ go get gopkg.in/thenativeweb/buntstift-go.v0

## Quick start

```go
b := buntstift.New()
```

To write messages to the console use the `Success` and `Error` methods to show that your application has succeeded or failed. If you want to provide additional information, use the `Info` and `Verbose` methods. In case of any warnings, use the `Warn` method.

```go
b.Info("Updating...")
b.Success("Done.");
```

## Printing blank lines

To print a blank line call the `NewLine` method.

```go
b.NewLine();
```

## Printing lines

To print a line call the `Line` method.

```go
b.Line();
```

## Using lists

To write a list to the console use the `List` method. Optionally, you may specify an indentation level. Setting the indentation level to `0` is equal to omitting it.

```go
b.List("foo");
b.List("bar", 0);
b.List("baz", 1);

//  ∙ foo
//  ∙ bar
//    ∙ baz
```

## Disabling colors

If you want to force disable colors set the NoColor option to true.

```go
var options = buntstift.Options{
  NoColor: true,
}

b := buntstift.New(options)
b.Info("No color")
```

## Disabling UTF characters

If your system does not support UTF characters, disable them by setting the NoUtf8 option to true.

```go
var options = buntstift.Options{
  NoUtf8:  true,
}

b := buntstift.New(options)
b.Info("Ascii")
```

## Waiting for long-running tasks

If your application performs a long-running task, you may use the `WaitFor` method to show a waiting indicator to the user.

```go
b.WaitFor(func() {
  // Do something
  // ...
});
```

## License

The MIT License (MIT)
Copyright (c) 2015 the native web.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
