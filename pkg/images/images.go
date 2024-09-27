package images

import (
	"errors"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"path/filepath"
	"strings"
)

type Thumb struct {
	src image.Image
	ext string
}

func NewThumb(reader io.Reader, filename string) (*Thumb, error) {
	ext := strings.ToLower(filepath.Ext(filename))
	if len(ext) == 0 {
		return nil, errors.New("unknown image ext")
	}

	// 短声明适用于赋值顺便定义类型，var适用于只定义类型；
	var img image.Image
	var err error

	switch ext[1:] {
	case "jpg":
		img, err = jpeg.Decode(reader)
	case "jpeg":
		img, err = jpeg.Decode(reader)
	case "gif":
		img, err = gif.Decode(reader)
	case "png":
		img, err = png.Decode(reader)
	default:
		return nil, errors.New("unknown image ext")
	}
	if err != nil {
		return nil, err
	}
	return &Thumb{
		src: img,
		ext: ext[1:],
	}, nil
}

func (tb *Thumb) GetSize() (int, int) {
	bounds := tb.src.Bounds()
	return bounds.Max.X, bounds.Max.Y
}

// 简化该方法；
// 一个struct可以有对应的修改属性，获取属性等读写操作，以及保存文件等其他操作；
func (*Thumb) SaveFile(path string) error {
	// ...
	return nil
}
