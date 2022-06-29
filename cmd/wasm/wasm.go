package main

import (
	"fmt"
	"syscall/js"
)

// Attempt to create a simple javascript class:
//
// class Rectangle {
//   constructor(height, width) {
//     this.height = height;
//     this.width = width;
//   }
//
//   // Method
//   area() {
//     return this.height * this.width;
//   }
// }

func Err(msg string) any {
	errorConstructor := js.Global().Get("Error")
	errorObject := errorConstructor.New(msg)
	return errorObject
}

func rectangleConstructor(this js.Value, args []js.Value) any {
	if this.IsUndefined() {
		return Err("must use new to initialize class Rectangle")
	}
	var h, w js.Value
	if len(args) < 2 {
		h = js.ValueOf(0)
		w = js.ValueOf(0)
	} else {
		h, w = args[0], args[1]
	}
	this.Set("height", h)
	this.Set("width", w)
	return this
}

func rectangleArea(this js.Value, args []js.Value) any {
	height := this.Get("height").Float()
	width := this.Get("width").Float()
	return height * width
}

func rectangleString(this js.Value, args []js.Value) any {
	return fmt.Sprintf("Rectangle(%.2f, %.2f)",
		this.Get("height").Float(),
		this.Get("width").Float())
}

func helloWorld(this js.Value, args []js.Value) any {
	return "Hello, World!"
}

func classRectangle() (string, js.Func) {
	const name = "Rectangle"
	rectangle := js.FuncOf(rectangleConstructor)
	rectangle.Set("prototype", map[string]any{
		"constructor": rectangle,
		"area":        js.FuncOf(rectangleArea),
		"toString":    js.FuncOf(rectangleString),
	})
	return name, rectangle
}

func funcMap() map[string]js.Func {
	rectName, rectConstructor := classRectangle()
	return map[string]js.Func{
		rectName:     rectConstructor,
		"helloWorld": js.FuncOf(helloWorld),
	}
}

func main() {
	for name, fn := range funcMap() {
		js.Global().Set(name, fn)
	}
	select {}
}
