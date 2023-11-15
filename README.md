# gofiber-minifier

## Introduction
Minify is a HTML5, CSS, and JavaScript minifier implemented as a middleware for [Fiber](https://gofiber.io/). 

Minification is the process of removing all unnecessary characters from code without altering its functionality. This includes the removal of whitespace and comments, which eventually results in more compact file sizes.

## Usage
```go
import "github.com/THREATINT/gofiber-minifier/minifier"
```

```go
app.Use(minifier.New())
```

## Credits
This Fiber middleware is a thin wrapper for [Minify](https://github.com/tdewolff/minify) by Taco de Wollf. 

## License
Release under the MIT License.

## QA
[![Codacy Badge](https://app.codacy.com/project/badge/Grade/52f0fd025b014597a177808959534f2f)](https://app.codacy.com/gh/THREATINT/gofiber-minifier/dashboard?utm_source=gh&utm_medium=referral&utm_content=&utm_campaign=Badge_grade)
