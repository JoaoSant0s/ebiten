// Copyright 2014 Hajime Hoshi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +build example

package main

import (
	_ "image/jpeg"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 320
	screenHeight = 240
)

var (
	count        int
	gophersImage *ebiten.Image
)

func update(screen *ebiten.Image) error {
	count++
	w, h := gophersImage.Size()
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(screenWidth-w)/2, float64(screenHeight-h)/2)
	op.ColorM.RotateHue(float64(count%360) * 2 * math.Pi / 360)
	screen.DrawImage(gophersImage, op)
	return nil
}

func main() {
	var err error
	gophersImage, _, err = ebitenutil.NewImageFromFile("_resources/images/gophers.jpg", ebiten.FilterNearest)
	if err != nil {
		log.Fatal(err)
	}
	if err := ebiten.Run(update, screenWidth, screenHeight, 2, "Hue (Ebiten Demo)"); err != nil {
		log.Fatal(err)
	}
}
