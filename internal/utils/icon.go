package utils

import (
	"bytes"
	"image"
	"image/png"
	"log"
	"os"

	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

func GetIcon(iconPath string) []byte {
	iconData, err := os.ReadFile(iconPath)
	if err != nil {
		log.Printf("Error reading icon file: %v", err)
		return nil
	}

	// Convert SVG to PNG
	icon, err := oksvg.ReadIconStream(bytes.NewReader(iconData))
	if err != nil {
		log.Printf("Error parsing SVG: %v", err)
		return nil
	}

	w, h := int(icon.ViewBox.W), int(icon.ViewBox.H)
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	icon.Draw(rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, img, img.Bounds())), 1)

	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		log.Printf("Error encoding PNG: %v", err)
		return nil
	}

	return buf.Bytes()
}
