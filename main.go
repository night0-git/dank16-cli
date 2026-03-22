package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"reflect"
)

func main() {
	primary := flag.String("color", "", "Primary hex color to base the palette on (Required)")
	light := flag.Bool("light", false, "Generate a light mode palette (Default: Dark)")
	noDPS := flag.Bool("no-dps", false, "Disable Delta Phi Star contrast calculation (Default: Enabled)")
	bg := flag.String("bg", "", "Optional custom background hex color")
	asString := flag.Bool("string", false, "Output as a simple string list instead of JSON")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Example: %s --color \"#3873f5\" --light --string\n\n", os.Args[0])
		
		// Manually formatting to show double dashes in help text
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  --color string\n    \tPrimary hex color (Required)\n")
		fmt.Fprintf(os.Stderr, "  --light\n    \tGenerate a light mode palette instead\n")
		fmt.Fprintf(os.Stderr, "  --no-dps\n    \tDisable Delta Phi Star contrast\n")
		fmt.Fprintf(os.Stderr, "  --bg string\n    \tCustom background hex\n")
		fmt.Fprintf(os.Stderr, "  --string\n    \tOutput as string list instead of JSON\n")
	}

	flag.Parse()

	if *primary == "" {
		flag.Usage()
		os.Exit(1)
	}

	opts := PaletteOptions{
		IsLight:    *light,
		Background: *bg,
		UseDPS:     !*noDPS,
	}

	palette := GeneratePalette(*primary, opts)

	if *asString {
		v := reflect.ValueOf(palette)
		for i := 0; i < v.NumField(); i++ {
			colorField := v.Field(i).Interface().(ColorInfo)
			fmt.Printf("Color %d: %s\n", i, colorField.Hex)
		}
	} else {
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "  ")
		if err := encoder.Encode(palette); err != nil {
			fmt.Fprintf(os.Stderr, "Error encoding JSON: %v\n", err)
			os.Exit(1)
		}
	}
}
