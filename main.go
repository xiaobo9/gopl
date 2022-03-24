// Echo2 prints its command-line arguments.
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xiaobo9/gopl/ch1/echo"
	"github.com/xiaobo9/gopl/ch1/fetch"
	"github.com/xiaobo9/gopl/ch1/lissajous"
	"github.com/xiaobo9/gopl/ch1/server"
	"github.com/xiaobo9/gopl/ch2/boiling"
	"github.com/xiaobo9/gopl/ch2/popcount"
	"github.com/xiaobo9/gopl/ch3/mandelbrot"
	"github.com/xiaobo9/gopl/ch3/surface"
	"github.com/xiaobo9/gopl/zeller"
)

// init log
func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

// type Action func()
// var mapFuncs map[string]Action

var actionMap map[string]func()

// init name to function map
func init() {
	// var mapFuncs = make(map[string]Action)
	// var mapFuncs = make(map[string]func())

	actionMap = map[string]func(){
		"echo":     func() { echo.Echo() },
		"fetch":    func() { fetch.FetchSingle() },
		"fetchAll": func() { fetch.Fetchall() },
		"server":   func() { server.Server() },
		"gcd":      func() { fmt.Println(gcd(15, 20)) },
		"boiling":  func() { boiling.Boiling() },
		"popcount": func() { popcount.PopCountTest() },
		"zeller":   func() { zeller.Zeller("2022-03-26") },
		// 图片
		"mandelbrot": func() { mandelbrot.Mandelbrot() },
		"surface":    func() { surface.SurfaceSvg() },
		"lissajous":  func() { lissajous.Lissajous() },
	}

}

func main() {
	if len(os.Args) <= 1 {
		log.Println("没有指定动作")
		return
	}
	actionName := os.Args[1]
	log.Println("action name: ", actionName)
	action := actionMap[actionName]
	if action != nil {
		(action)()
		return
	}
	log.Println("没有找到指定动作: ", actionName)
}

func gcd(x, y int) int {
	for y != 0 {
		fmt.Printf("%d %d %d\n", x, y, x%y)
		x, y = y, x%y
	}
	return x
}
