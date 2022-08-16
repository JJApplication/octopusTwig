/*
Create: 2022/8/15
Project: octopusTwig
Github: https://github.com/landers1037
Copyright Renj
*/

// Package main
package main

import (
	octopusTwig "github.com/JJApplication/octopus_twig"
)

func main() {
	twig := octopusTwig.TwigCore{Name: octopusTwig.OctopusTwig}
	twig.Init()
	twig.Start()
}
