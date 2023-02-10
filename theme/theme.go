package theme

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"image/color"
)

var msyh []byte

type MyTheme struct{}

var _ fyne.Theme = (*MyTheme)(nil)

func (m MyTheme) Font(s fyne.TextStyle) fyne.Resource {
	if s.Bold {
		return ResourceMsyhbdTtf
	}
	if s.Italic {
		return theme.DefaultTheme().Font(s)
	}
	if s.Monospace {
		return theme.DefaultTheme().Font(s)
	}
	return ResourceMsyhTtf
}
func (*MyTheme) Color(n fyne.ThemeColorName, v fyne.ThemeVariant) color.Color {
	return theme.DefaultTheme().Color(n, v)
}

func (*MyTheme) Icon(n fyne.ThemeIconName) fyne.Resource {
	return theme.DefaultTheme().Icon(n)
}

func (*MyTheme) Size(n fyne.ThemeSizeName) float32 {
	return theme.DefaultTheme().Size(n)
}
