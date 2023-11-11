# gofiber-minifier

## Introduction
Minify is a HTML5 minifier implemented as a middleware for [Fiber](https://gofiber.io/). 

Minification is the process of removing all unnecessary characters from source code 
without altering its functionality. This includes the removal of whitespace and comments. 
Minification of HTML5 code results in compact file size.

## Usage
```
import "github.com/THREATINT/gofiber-minifier/minifier"
```

```
app.Use(minifier.New())
```

## Credits
This Fiber middleware is a thin wrapper for [Minify](https://github.com/tdewolff/minify) 
by Taco de Wollf. 

## License
Release under the MIT License. (see LICENSE)

## QA
