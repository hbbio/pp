package pp

import (
	"fmt"
	"reflect"
)

const (
	// NoColor has no color
	NoColor uint16 = 1 << 15
)

const (
	// Foreground colors for ColorScheme.
	_ uint16 = iota | NoColor
	// Black color
	Black
	// Red color
	Red
	// Green color
	Green
	// Yellow color
	Yellow
	// Blue color
	Blue
	// Magenta color
	Magenta
	// Cyan color
	Cyan
	// White color
	White
	bitsForeground       = 0
	maskForegorund       = 0xf
	ansiForegroundOffset = 30 - 1
)

const (
	// Background colors for ColorScheme.
	_ uint16 = iota<<bitsBackground | NoColor
	// BackgroundBlack color
	BackgroundBlack
	// BackgroundRed color
	BackgroundRed
	// BackgroundGreen color
	BackgroundGreen
	// BackgroundYellow color
	BackgroundYellow
	// BackgroundBlue color
	BackgroundBlue
	// BackgroundMagenta color
	BackgroundMagenta
	// BackgroundCyan color
	BackgroundCyan
	// BackgroundWhite color
	BackgroundWhite
	bitsBackground       = 4
	maskBackground       = 0xf << bitsBackground
	ansiBackgroundOffset = 40 - 1
)

const (
	// Bold flag for ColorScheme.
	Bold     uint16 = 1<<bitsBold | NoColor
	bitsBold        = 8
	maskBold        = 1 << bitsBold
	ansiBold        = 1
)

// ColorScheme is used with SetColorScheme.
type ColorScheme struct {
	Bool            uint16
	Integer         uint16
	Float           uint16
	String          uint16
	StringQuotation uint16
	EscapedChar     uint16
	FieldName       uint16
	PointerAdress   uint16
	Nil             uint16
	Time            uint16
	StructName      uint16
	ObjectLength    uint16
}

var (
	// ColoringEnabled can enable/disable coloring
	ColoringEnabled = true

	defaultScheme = ColorScheme{
		Bool:            Cyan | Bold,
		Integer:         Blue | Bold,
		Float:           Magenta | Bold,
		String:          Red,
		StringQuotation: Red | Bold,
		EscapedChar:     Magenta | Bold,
		FieldName:       Yellow,
		PointerAdress:   Blue | Bold,
		Nil:             Cyan | Bold,
		Time:            Blue | Bold,
		StructName:      Green,
		ObjectLength:    Blue,
	}
)

func (cs *ColorScheme) fixColors() {
	typ := reflect.Indirect(reflect.ValueOf(cs))
	defaultType := reflect.ValueOf(defaultScheme)
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Uint() == 0 {
			field.SetUint(defaultType.Field(i).Uint())
		}
	}
}

func colorize(text string, color uint16) string {
	if !ColoringEnabled {
		return text
	}

	foreground := color & maskForegorund >> bitsForeground
	background := color & maskBackground >> bitsBackground
	bold := color & maskBold

	if foreground == 0 && background == 0 && bold == 0 {
		return text
	}

	modBold := ""
	modForeground := ""
	modBackground := ""

	if bold > 0 {
		modBold = "\033[1m"
	}
	if foreground > 0 {
		modForeground = fmt.Sprintf("\033[%dm", foreground+ansiForegroundOffset)
	}
	if background > 0 {
		modBackground = fmt.Sprintf("\033[%dm", background+ansiBackgroundOffset)
	}

	return fmt.Sprintf("%s%s%s%s\033[0m", modForeground, modBackground, modBold, text)
}
