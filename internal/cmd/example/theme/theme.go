package theme

import (
	_ "embed"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
)

//go:embed PingFang.ttf
var font []byte

var myfont = &fyne.StaticResource{
	StaticName:    "PingFang.ttf",
	StaticContent: font,
}

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

func (m *MyTheme) Font(_ fyne.TextStyle) fyne.Resource {
	return myfont
	//return theme.DefaultTheme().Font(s)
}

func (m *MyTheme) Size(n fyne.ThemeSizeName) float32 {
	if n == theme.SizeNameInlineIcon {
		return 48
	}
	return theme.DefaultTheme().Size(n)
}

func (m *MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	v = theme.VariantDark
	return theme.DefaultTheme().Color(n, v)
}

func (m *MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}
