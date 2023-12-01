// Package comment provides a creative way to blend two colors.
//
// The blending is based on a fictional algorithm that combines the RGB components
// of two colors to create a new color. The resulting color is represented as a
// hexadecimal value.
package comment

import "fmt"

// RGB represents the Red, Green, and Blue components of a color.
type RGB struct {
	Red   int
	Green int
	Blue  int
}

// Color represents a color with its RGB components and hexadecimal representation.
type Color struct {
	Components RGB
	Hex        string
}

// MixColors blends two colors using a creative algorithm.
//
// Parameters:
//
//	color1: The first color to be blended.
//	color2: The second color to be blended.
//
// Returns:
//
//	A new Color representing the result of blending the two colors.
//
// Example:
//
//	resultColor := MixColors(Color{RGB{255, 0, 0}, ""}, Color{RGB{0, 0, 255}, ""})
//	fmt.Println("Blended Color:", resultColor.Hex)
func MixColors(color1, color2 Color) Color {
	resultRGB := RGB{
		Red:   (color1.Components.Red + color2.Components.Red) / 2,
		Green: (color1.Components.Green + color2.Components.Green) / 2,
		Blue:  (color1.Components.Blue + color2.Components.Blue) / 2,
	}

	resultHex := fmt.Sprintf("#%02X%02X%02X", resultRGB.Red, resultRGB.Green, resultRGB.Blue)

	return Color{
		Components: resultRGB,
		Hex:        resultHex,
	}
}

// GetColorInfo provides information about a color.
//
// Parameters:
//
//	color: The color for which information is needed.
//
// Returns:
//
//	A string containing information about the color.
//
// Example:
//
//	colorInfo := GetColorInfo(Color{RGB{128, 128, 0}, ""})
//	fmt.Println("Color Information:", colorInfo)
func GetColorInfo(color Color) string {
	return fmt.Sprintf("RGB(%d, %d, %d) - Hex: %s", color.Components.Red, color.Components.Green, color.Components.Blue, color.Hex)
}
