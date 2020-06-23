package imgcode

import (
	"image/jpeg"
	"log"
	"os"

	"github.com/derpl-del/gopro2/envcode/logcode"
	"github.com/nfnt/resize"
)

//ResizeImg for Image
func ResizeImg(input string) {
	filetittle := "data_img/upload_" + input + "_1.png"
	file, err := os.Open(filetittle)
	if err != nil {
		logcode.LogE(err)
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		log.Fatal(err)
	}
	file.Close()
	m := resize.Resize(0, 620, img, resize.Lanczos3)
	out, err := os.Create(filetittle)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	// write new image to file
	jpeg.Encode(out, m, nil)
}
