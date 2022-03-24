// Mandelbrot emits a PNG image of the Mandelbrot fractal.
// ch3-03
// https://books.studygolang.com/gopl-zh/ch3/ch3-03.html
package mandelbrot

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"os"
	"path/filepath"
	"time"

	"github.com/xiaobo9/gopl/utils"
)

func Mandelbrot() {
	fileName := fmt.Sprintf("mandelbrot_%v.png", time.Now().Unix())
	f, err := utils.CreateTempFile(fileName)
	log.Println(filepath.Abs(f.Name()))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		// xmin, ymin, xmax, ymax = -4, -4, +4, +4
		width, height = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin
			z := complex(x, y)
			// Image point (px, py) represents complex value z.
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(f, img) // NOTE: ignoring errors
	f.Close()
}

func mandelbrot(z complex128) color.Color {
	const iterations = 240
	const contrast = 30

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return color.Gray{255 - contrast*n}
		}
	}
	// return color.Black
	return color.RGBA{0x00, 0x00, 0xBB, 0xff}
}

// 练习 3.5： 实现一个彩色的 Mandelbrot 图像，使用 image.NewRGBA 创建图像，使用color.RGBA或color.YCbCr生成颜色。

// 练习 3.6： 升采样技术可以降低每个像素对计算颜色值和平均值的影响。简单的方法是将每个像素分成四个子像素，实现它。

// 练习 3.7： 另一个生成分形图像的方式是使用牛顿法来求解一个复数方程，例如$z^4-1=0$。每个起点到四个根的迭代次数对应阴影的灰度。方程根对应的点用颜色表示。

// 练习 3.8： 通过提高精度来生成更多级别的分形。使用四种不同精度类型的数字实现相同的分形：complex64、complex128、big.Float和big.Rat。（后面两种类型在math/big包声明。Float是有指定限精度的浮点数；Rat是无限精度的有理数。）它们间的性能和内存使用对比如何？当渲染图可见时缩放的级别是多少？

// 练习 3.9： 编写一个web服务器，用于给客户端生成分形的图像。运行客户端通过HTTP参数指定x、y和zoom参数。
