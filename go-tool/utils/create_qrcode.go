package utils

import (
	"errors"
	"fmt"
	"github.com/nfnt/resize"
	"github.com/skip2/go-qrcode"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
)

var err error


func CreateAvatar(avatar ,resu ,url string) (image.Image, error) {
	var (
		bgImg      image.Image
		offset     image.Point
		avatarFile *os.File
		avatarImg  image.Image
	)

	bgImg, err = createQrCode(url)

	if err != nil {
		fmt.Println("创建二维码失败:", err)
		return nil, errors.New("创建二维码失败")
	}
	avatarFile, err = os.Open(avatar)
	avatarImg, err = jpeg.Decode(avatarFile)
	avatarImg = ImageResize(avatarImg, 40, 40)
	b := bgImg.Bounds()

	// 设置为居中
	offset = image.Pt((b.Max.X-avatarImg.Bounds().Max.X)/2, (b.Max.Y-avatarImg.Bounds().Max.Y)/2)

	m := image.NewRGBA(b)

	draw.Draw(m, b, bgImg, image.Point{X: 0, Y: 0,}, draw.Src)

	draw.Draw(m, avatarImg.Bounds().Add(offset), avatarImg, image.Point{X: 0, Y: 0}, draw.Over)

	f,err := os.Create(resu)
	defer f.Close()
	err = png.Encode(f,m)
	return m, err
}

func createQrCode(content string) (img image.Image, err error) {
	var qrCode *qrcode.QRCode

	qrCode, err = qrcode.New(content, qrcode.Highest)

	if err != nil {
		return nil, errors.New("创建二维码失败")
	}
	qrCode.DisableBorder = true

	img = qrCode.Image(150)

	return img, nil
}

func ImageResize(src image.Image, w, h int) image.Image {
	return resize.Resize(uint(w), uint(h), src, resize.Lanczos3)
}