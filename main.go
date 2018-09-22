package main

import (
	"image"
	"image/png"
	"log"
	"os"
)

func main() {
	// Output image
	rect := image.Rectangle{Max: image.Point{X: 800, Y: 800}}
	img := newImage(rect)

	// Mesh
	obj, err := loadObjFromFile("models/african_head.obj")
	if err != nil {
		log.Fatalln("Unable to load obj file:", err)
	}

	// Texture
	textureFile, err := os.Open("textures/african_head_diffuse.png")
	if err != nil {
		log.Fatalln("Unable to read texture from file:", err)
	}
	texture, err := png.Decode(textureFile)
	if err != nil {
		log.Fatalln("Unable to decode texture:", err)
	}
	texture = flipImageVertically(texture.Bounds(), texture)

	// Render
	//now := time.Now()
	//fps := 0
	//for time.Since(now) <= time.Second {
	render(img, obj, texture)
	//	fps++
	//}
	//fmt.Println("FPS:", fps)

	// Saving
	img = flipImageVertically(rect, img)
	saveImage(img)
}

func render(img *image.RGBA, obj *Obj, texture image.Image) {
	rect := img.Bounds()
	width := rect.Dx()
	height := rect.Dy()

	zBuffer := make([]float64, width*height)
	for i := 0; i < len(zBuffer); i++ {
		zBuffer[i] = -1.0
	}

	for _, face := range obj.Faces {
		triangle := Triangle{}

		for i := 0; i < 3; i++ {
			localVertex := face.Vertices[i]

			// Embed the 3D coordinate into 4D temporarily.
			vertex4 := Vertex4{
				X: localVertex.X,
				Y: localVertex.Y,
				Z: localVertex.Z,
				W: 1,
			}

			// Map from an object's local coordinate space into world coordinate space.
			cos90 := -0.44807361613
			sin90 := 0.8939966636
			modelMatrix := Matrix4{
				cos90, 0, sin90, 0,
				0, 1, 0, 0,
				-sin90, 0, cos90, 0,
				0, 0, 0, 1,
			}

			// Map from world space to camera space.
			cameraMatrix := Identity4()

			// Map from camera space to screen.
			screenMatrix := genScreenMatrix(0, 0, width, height)

			vertex4.transform(modelMatrix)
			vertex4.transform(cameraMatrix)
			vertex4.transform(screenMatrix)

			// Bring back 4D into 3D.
			vertex3 := vertex4.lower()

			triangle.points[i].X = int(vertex3.X)
			triangle.points[i].Y = int(vertex3.Y)
		}

		drawTriangle(
			img,
			triangle,
			zBuffer,
			texture,
			face,
		)
	}
}

func genScreenMatrix(x, y, w, h int) Matrix4 {
	d := 255

	return Matrix4{
		float64(w) / 2.0, 0, 0, float64(x) + float64(w)/2.0,
		0, float64(h) / 2.0, 0, float64(y) + float64(h)/2.0,
		0, 0, float64(d) / 2.0, float64(d) / 2.0,
		0, 0, 0, 1,
	}
}

func genCameraMatrix(eye Vertex3, center Vertex3, up Vertex3) Matrix4 {
	z := eye.minus(center).normalize(1.0)
	x := up.cross(z).normalize(1.0)
	y := z.cross(x).normalize(1.0)

	minv := Identity4()
	tr := Identity4()

	minv.m11 = x.X
	minv.m21 = y.X
	minv.m31 = z.X
	minv.m12 = x.Y
	minv.m22 = y.Y
	minv.m32 = z.Y
	minv.m13 = x.Z
	minv.m23 = y.Z
	minv.m33 = z.Z

	tr.m13 = -center.X
	tr.m23 = -center.Y
	tr.m33 = -center.Z

	return minv.Dot(tr)
}