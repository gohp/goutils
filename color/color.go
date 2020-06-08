package color

import (
	"fmt"
)

/**
* @Author: Jam Wong
* @Date: 2020/6/8
 */

const (
	colorTextBlack   = 30
	colorTextRed     = 31
	colorTextGreen   = 32
	colorTextYellow  = 33
	colorTextBlue    = 34
	colorTextMagenta = 35
	colorTextCyan    = 36
	colorTextWhite   = 37

	colorBgBlack   = 40
	colorBgRed     = 41
	colorBgGreen   = 42
	colorBgYellow  = 43
	colorBgBlue    = 44
	colorBgMagenta = 45
	colorBgCyan    = 46
	colorBgWhite   = 47

	textDefault   = 0
	textBold      = 1
	textUnderline = 4
	textBlink     = 5
	textReverse   = 7
	textHidden    = 8
)

var (
	Default = &Color{0, 0, textDefault}
	Black   = &Color{colorTextBlack, 0, textDefault}
	Red     = &Color{colorTextRed, 0, textDefault}
	Green   = &Color{colorTextGreen, 0, textDefault}
	Yellow  = &Color{colorTextYellow, 0, textDefault}
	Blue    = &Color{colorTextBlue, 0, textDefault}
	Magenta = &Color{colorTextMagenta, 0, textDefault}
	Cyan    = &Color{colorTextCyan, 0, textDefault}
	White   = &Color{colorTextWhite, 0, textDefault}
)

type Color struct {
	Text uint8
	Bg   uint8
	Font uint8
}

func (c *Color) Bold() *Color {
	return &Color{c.Text, c.Bg, textBold}
}

func (c *Color) Underline() *Color {
	return &Color{c.Text, c.Bg, textUnderline}
}

func (c *Color) Blink() *Color {
	return &Color{c.Text, c.Bg, textBlink}
}

func (c *Color) Reverse() *Color {
	return &Color{c.Text, c.Bg, textReverse}
}

func (c *Color) Hidden() *Color {
	return &Color{c.Text, c.Bg, textHidden}
}

func (c *Color) BlackBg() *Color {
	return &Color{c.Text, colorBgBlack, c.Font}
}

func (c *Color) RedBg() *Color {
	return &Color{c.Text, colorBgRed, c.Font}
}

func (c *Color) GreenBg() *Color {
	return &Color{c.Text, colorBgGreen, c.Font}
}

func (c *Color) BlueBg() *Color {
	return &Color{c.Text, colorBgBlue, c.Font}
}

func (c *Color) YellowBg() *Color {
	return &Color{c.Text, colorBgYellow, c.Font}
}

func (c *Color) MagentaBg() *Color {
	return &Color{c.Text, colorBgMagenta, c.Font}
}

func (c *Color) CyanBg() *Color {
	return &Color{c.Text, colorBgCyan, c.Font}
}

func (c *Color) WhiteBg() *Color {
	return &Color{c.Text, colorBgWhite, c.Font}
}

func (c *Color) color(msg string) string {
	if c.Font == 0 {
		return fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m", c.Bg, c.Text, msg)
	}
	if c.Bg == 0 && c.Font != 0 {
		return fmt.Sprintf("\x1b[%d;%dm%s\x1b[0m", c.Text, c.Font, msg)
	}
	return fmt.Sprintf("\x1b[%d;%d;%dm%s\x1b[0m", c.Font, c.Bg, c.Text, msg)
}

func (c *Color) Println(s string) {
	fmt.Println(c.color(s))
}

func (c *Color) Printf(layout string, a ...interface{}) {
	fmt.Printf(c.color(fmt.Sprintf(layout, a...)))
}

func (c *Color) Sprintf(layout string, a ...interface{}) string {
	return c.color(fmt.Sprintf(layout, a...))
}
