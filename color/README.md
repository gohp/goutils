# Color

```go
func main() {
	color.Black.Println("Black")
	color.Red.Println("Red")
	color.Green.Println("Green")
	color.Blue.Println("Blue")
	color.Yellow.Println("Yellow")
	color.Magenta.Println("Magenta")
	color.Cyan.Println("Cyan")
	color.White.Println("White")

	color.Black.RedBg().Println("RedBg")
	color.Black.BlueBg().Println("BlueBg")
	color.Black.MagentaBg().Println("MagentaBg")
	color.Black.YellowBg().Println("YellowBg")
	color.Black.CyanBg().Println("CyanBg")
	color.White.BlackBg().Println("BlackBg")
	color.Black.WhiteBg().Println("WhiteBg")
	color.Black.GreenBg().Println("GreenBg")

	color.Black.RedBg().Println("Default")
	color.Black.RedBg().Bold().Println("Bold")
	color.Black.RedBg().Underline().Println("Underline")
	color.Black.RedBg().Blink().Println("Blink")
	color.Black.RedBg().Reverse().Println("Reverse")
	color.Black.RedBg().Hidden().Println("Hidden")
}
```