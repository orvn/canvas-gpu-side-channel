package main

import (
	"encoding/base64"
	"encoding/json"
	"image/png"
	"net/http"
	"os"
	"strings"

	"github.com/n7olkachev/imgdiff"
)

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	json.NewDecoder(r.Body).Decode(&data)
	imageDataUrl := data["imageDataUrl"]
	imageData, _ := base64.StdEncoding.DecodeString(imageDataUrl[strings.IndexByte(imageDataUrl, ',')+1:])
	os.WriteFile("image.png", imageData, 0644)
	// Previous image is saved as previous.png
	diffImage("previous.png", "image.png")
}

func diffImage(image1Path, image2Path string) {
	file1, _ := os.Open(image1Path)
	defer file1.Close()
	file2, _ := os.Open(image2Path)
	defer file2.Close()
	img1, _ := png.Decode(file1)
	img2, _ := png.Decode(file2)
	diff, _ := imgdiff.Diff(img1, img2)
	diffImg := imgdiff.DrawDiff(img1, img2, diff)
	outFile, _ := os.Create("diff.png")
	png.Encode(outFile, diffImg)
	outFile.Close()
}

func main() {
	http.HandleFunc("/upload", uploadHandler)
	http.ListenAndServe(":8090", nil)
}
