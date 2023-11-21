package minifier

import (
	"bytes"
	"github.com/gofiber/fiber/v2"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
	"github.com/tdewolff/minify/v2/html"
	"github.com/tdewolff/minify/v2/js"
	"regexp"
)

func New(config ...Config) fiber.Handler {
	var (
		cfg Config = configDefault(config...)
	)

	return func(c *fiber.Ctx) error {
		var (
			err      error
			origBody []byte
			m        *minify.M
		)

		if cfg.Next != nil && cfg.Next(c) {
			return c.Next()
		}

		if err = c.Next(); err != nil {
			return err
		}

		if c.Response().StatusCode() != fiber.StatusOK {
			return nil
		}

		origBody = c.Response().Body()
		c.Response().ResetBody()

		m = minify.New()
		if cfg.minifyHTML {
			m.AddFunc("text/html", html.Minify)
		}
		if cfg.minifyCSS {
			m.AddFunc("text/css", css.Minify)
		}
		if cfg.minifyJS {
			m.AddFuncRegexp(regexp.MustCompile("^(application|text)/(x-)?(java|ecma)script$"), js.Minify)
		}
		if err = m.Minify("text/html", c.Response().BodyWriter(), bytes.NewReader(origBody)); err != nil {
			// just in case minifying does not work for any reason, we fail in a gentle way
			// by writing the original (un-minified) body
			c.Response().BodyWriter().Write(origBody)
		}

		return nil
	}
}
