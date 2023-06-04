package backend

import (
	"encoding/base64"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"log"
	"os"
	"real-time-forum/db"
	"strconv"
	"strings"

	"github.com/gorilla/websocket"
)

func ProfPic(pic string, id int, act string, count string, conn *websocket.Conn) {
	r := EData{}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(strings.Split(pic, "base64,")[1]))
	indx := strings.Index(string(pic), ",")
	typ := strings.TrimSuffix(pic[5:indx], ";base64")
	name := act + strconv.Itoa(id) + "_" + count + "." + strings.Split(typ, "/")[1]
	if act == "profile_" {
		err := db.UpdProfPic(name, id)
		if err != nil {
			r.Errrr = "Can't save image"
			conn.WriteJSON(r)
		}
	}
	switch typ {
	case "image/png":
		log.Println("png")
		img, _ := png.Decode(reader)
		saveToPNG(img, name)
	case "image/jpeg":
		log.Println("jpeg")
		img, _ := jpeg.Decode(reader)
		saveToJPEG(img, name)
	case "image/gif":
		log.Println("gif")
		img, _ := gif.Decode(reader)
		saveToGIF(img, name)
	}
}
func saveToGIF(img image.Image, name string) {
	f, err := os.Create("./Frontend/images/" + name)
	if err != nil {
		// err
	}
	defer f.Close()
	opt := gif.Options{
		NumColors: 256,
	}
	err = gif.Encode(f, img, &opt)
	if err != nil {
		// err
	}
}
func saveToJPEG(img image.Image, name string) {
	f, err := os.Create("./Frontend/images/" + name)
	if err != nil {
		// err
	}
	defer f.Close()
	opt := jpeg.Options{
		Quality: 100,
	}
	err = jpeg.Encode(f, img, &opt)
	if err != nil {
		// err
	}
}
func saveToPNG(img image.Image, name string) {
	f, err := os.Create("./Frontend/images/" + name)
	if err != nil {
		// err
	}
	defer f.Close()
	err = png.Encode(f, img)
	if err != nil {
		// err
	}
}
