package main

import (
	"fmt"
	"image"
	"image/draw"
	_ "image/gif"
	"image/jpeg"
	_ "image/jpeg"
	"image/png"
	_ "image/png"
	"io/ioutil"
	"log"
	"os"
	"regexp"
)

func main() {
	nowPath, _ := os.Getwd()
	log.Println("获取当前目录", nowPath)

	// todo: 读取当前目录下 png 或者jpg的图片
	fileinf, err := ioutil.ReadDir("./")
	if err != nil {
		fmt.Println("读取目录失败，错误:", err)
		return
	}

	num := 0

	for _, file := range fileinf {
		if file.IsDir() {
			log.Println("是一个文件夹 跳过")
		} else {
			log.Println("是一个文件", file.Name())
			if match, _ := regexp.MatchString(".png", file.Name()); match {
				num++
				splitImg(file.Name(), num)
			}
			if match, _ := regexp.MatchString(".jpg", file.Name()); match {
				num++
				splitImg(file.Name(), num)
			}
		}
	}
}

func splitImg(filePath string, num int) {
	file, err := os.Open(filePath)

	if err != nil {
		panic(err)
	}

	defer file.Close()
	img, err := png.Decode(file)
	if err != nil {
		panic(err)
	}
	dir := fmt.Sprintf("try%d", num)

	os.Mkdir(dir, 777)

	newImg1 := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx()/2, img.Bounds().Dy()/2))
	newImg2 := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx()/2, img.Bounds().Dy()/2))
	newImg3 := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx()/2, img.Bounds().Dy()/2))
	newImg4 := image.NewRGBA(image.Rect(0, 0, img.Bounds().Dx()/2, img.Bounds().Dy()/2))

	draw.Draw(newImg1, newImg1.Bounds(), img, image.Point{
		X: img.Bounds().Dx() / 2,
		Y: img.Bounds().Dy() / 2,
	}, draw.Src)
	draw.Draw(newImg2, newImg2.Bounds(), img, img.Bounds().Min, draw.Src)
	draw.Draw(newImg3, newImg3.Bounds(), img, image.Point{
		X: img.Bounds().Dx() / 2,
		Y: 0,
	}, draw.Src)
	draw.Draw(newImg4, newImg4.Bounds(), img, image.Point{
		X: 0,
		Y: img.Bounds().Dy() / 2,
	}, draw.Src)

	outputFile1, err := os.Create(dir + "/img1.jpg")
	if err != nil {
		panic(err)
	}
	defer outputFile1.Close()
	jpeg.Encode(outputFile1, newImg1, nil)

	outputFile2, err := os.Create(dir + "/img2.jpg")
	if err != nil {
		panic(err)
	}
	defer outputFile2.Close()
	jpeg.Encode(outputFile2, newImg2, nil)

	outputFile3, err := os.Create(dir + "/img3.jpg")
	if err != nil {
		panic(err)
	}
	defer outputFile3.Close()
	jpeg.Encode(outputFile3, newImg3, nil)

	outputFile4, err := os.Create(dir + "/img4.jpg")
	if err != nil {
		panic(err)
	}
	defer outputFile4.Close()
	jpeg.Encode(outputFile4, newImg4, nil)

}
