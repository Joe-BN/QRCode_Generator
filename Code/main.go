package main

import (
	"fmt"
	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

func main() {
	qrc, err := qrcode.New(" <link/ text>")
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}
	
	w, err := standard.New("<imagename>.png") // relative/absolute path
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	}
	
	// save file
	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}