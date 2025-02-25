package main

import (
 "fmt"
 "image"
 "image/color"
 "image/jpeg"
 "log"
 "os"
 "path/filepath"
 "sort"
 "strconv"
 "strings"
 _"image/png"
 _"golang.org/x/image/webp"
 "github.com/nfnt/resize"
)

const (
 inputFolder  = "img"
 outputFolder = "dataset"
 imageSize    = 728
 contrastFactor = 1.0
)

func main() {
 err := processImages()
 if err != nil {
  log.Fatal(err)
 }
 fmt.Println("Обработка завершена.")
}

func processImages() error {
	files, err := os.ReadDir(inputFolder)
	if err != nil {
		return err
	}

	nextIndex, err := getNextFileIndex(outputFolder)
	if err != nil {
		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		filePath := filepath.Join(inputFolder, file.Name())
		newFileName := fmt.Sprintf("img_%d.jpg", nextIndex) 
		outputPath := filepath.Join(outputFolder, newFileName)

		err := resizeAndSave(filePath, outputPath)
		if err != nil {
			fmt.Printf("Ошибка обработки %s: %v\n", file.Name(), err)
			continue
		}

		fmt.Printf("Обработано: %s -> %s\n", filePath, newFileName)

		err = os.Remove(filePath)
		if err != nil {
			fmt.Printf("Ошибка удаления %s: %v\n", file.Name(), err)
		}

		nextIndex++
	}

	return nil
}


func resizeAndSave(inputPath, outputPath string) error {
 file, err := os.Open(inputPath)
 if err != nil {
  return err
 }
 defer file.Close()

 img, _, err := image.Decode(file)
 if err != nil {
  return err
 }

 resizedImg := resize.Resize(imageSize, imageSize, img, resize.Lanczos3)
 contrastImg := increaseContrast(resizedImg, contrastFactor)

 outFile, err := os.Create(outputPath)
 if err != nil {
  return err
 }
 defer outFile.Close()

 return jpeg.Encode(outFile, contrastImg, nil)
}

func increaseContrast(img image.Image, factor float64) image.Image {
 bounds := img.Bounds()
 contrastImg := image.NewRGBA(bounds)

 for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
  for x := bounds.Min.X; x < bounds.Max.X; x++ {
   oldColor := img.At(x, y)
   r, g, b, a := oldColor.RGBA()

   r = uint32(float64(r>>8) * factor)
   g = uint32(float64(g>>8) * factor)
   b = uint32(float64(b>>8) * factor)

   if r > 255 {
    r = 255
   }
   if g > 255 {
    g = 255
   }
   if b > 255 {
    b = 255
   }

   newColor := color.RGBA{uint8(r), uint8(g), uint8(b), uint8(a >> 8)}
   contrastImg.Set(x, y, newColor)
  }
 }

 return contrastImg
}

func getNextFileIndex(folder string) (int, error) {
 files, err := os.ReadDir(folder)
 if err != nil {
  return 1, err
 }

 var indices []int
 for _, file := range files {
  name := file.Name()
  if strings.HasPrefix(name, "luminal_") && strings.HasSuffix(name, ".jpg") {
   parts := strings.Split(strings.TrimSuffix(name, ".jpg"), "_")
   if len(parts) < 2 {
    continue
   }
   index, err := strconv.Atoi(parts[1])
   if err == nil {
    indices = append(indices, index)
   }
  }
 }

 sort.Ints(indices)
 if len(indices) == 0 {
  return 1, nil
 }

 return indices[len(indices)-1] + 1, nil
}