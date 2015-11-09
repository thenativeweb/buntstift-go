# buntstift-go

buntstift-go makes the CLI colorful.

![buntstift](https://github.com/thenativeweb/buntstift-go/raw/master/images/logo.jpg "buntstift")

## Installation

```
  go get gopkg.in/thenativeweb/buntstift-go.v0
```

## Quick start

```
b := buntstift.New()

```
With options:

```

var options = buntstift.Options{
  NoColor: true,
  NoUtf8:  true,
}

b := buntstift.New(options)

b.Success("Hello World")

```
Methods:

```
func (b *Buntstift) WaitFor(worker func())
    WaitFor shows a spinner while worker is in progress

func (b *Buntstift) Error(text string)
    Error ...

func (b *Buntstift) Info(text string)
    Info ...

func (b *Buntstift) Line()
    Line ...

func (b *Buntstift) List(text string)
    List ...

func (b *Buntstift) ListIndent(level int, text string)
    ListIndent ...

func (b *Buntstift) Success(text string)
    Success ...

func (b *Buntstift) Warn(text string)
    Warn ...

```


## License

The MIT License (MIT)
Copyright (c) 2015 the native web.

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
