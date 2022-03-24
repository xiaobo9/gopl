// Surface computes an SVG rendering of a 3-D surface function.
// ch3-02.html 浮点数
// https://books.studygolang.com/gopl-zh/ch3/ch3-02.html
// TODO
package surface

import (
	"fmt"
	"io"
	"log"
	"math"
	"path/filepath"
	"time"

	"github.com/xiaobo9/gopl/utils"
)

const (
	// width, height = 600, 320            // canvas size in pixels
	width, height = 600, 320            // canvas size in pixels
	cells         = 50                  // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func SurfaceSvg() {
	fileName := fmt.Sprintf("surface%v.svg", time.Now().Unix())
	file, err := utils.CreateTempFile(fileName)
	if err != nil {
		log.Println("open file failed, ", err)
	}
	defer file.Close()

	WriteSurfaceSvg(file)
	absPath, _ := filepath.Abs(fileName)
	log.Println("out svg to file: ", absPath)
}

func WriteSurfaceSvg(writer io.StringWriter) {
	header := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	writer.WriteString(header)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)
			content := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			writer.WriteString(content)
		}
	}
	writer.WriteString("</svg>")
}

func corner(i, j int) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := computeHeight(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func computeHeight(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// 练习 3.1： 如果f函数返回的是无限制的float64值，那么SVG文件可能输出无效的多边形元素（虽然许多SVG渲染器会妥善处理这类问题）。修改程序跳过无效的多边形。

// 练习 3.2： 试验math包中其他函数的渲染图形。你是否能输出一个egg box、moguls或a saddle图案?

// 练习 3.3： 根据高度给每个多边形上色，那样峰值部将是红色（#ff0000），谷部将是蓝色（#0000ff）。

// 练习 3.4： 参考1.7节Lissajous例子的函数，构造一个web服务器，用于计算函数曲面然后返回SVG数据给客户端。服务器必须设置Content-Type头部：

// w.Header().Set("Content-Type", "image/svg+xml")
// （这一步在Lissajous例子中不是必须的，因为服务器使用标准的PNG图像格式，可以根据前面的512个字节自动输出对应的头部。）允许客户端通过HTTP请求参数设置高度、宽度和颜色等参数。
