// Code generated by 'yaegi extract gioui.org/op/paint'. DO NOT EDIT.

package main

import (
	"gioui.org/op/paint"
	"reflect"
)

func init() {
	Symbols["gioui.org/op/paint"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"Fill":       reflect.ValueOf(paint.Fill),
		"FillShape":  reflect.ValueOf(paint.FillShape),
		"NewImageOp": reflect.ValueOf(paint.NewImageOp),

		// type definitions
		"ColorOp":          reflect.ValueOf((*paint.ColorOp)(nil)),
		"ImageOp":          reflect.ValueOf((*paint.ImageOp)(nil)),
		"LinearGradientOp": reflect.ValueOf((*paint.LinearGradientOp)(nil)),
		"PaintOp":          reflect.ValueOf((*paint.PaintOp)(nil)),
	}
}
