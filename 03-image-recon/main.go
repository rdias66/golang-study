package main

import (
	"fmt"
	"gocv.io/x/gocv"
)

func main() {
  imageName := "user inputed image, not figured out yet"
	imagePath := "./image/" + imageName // Replace with the path to your image file

	// Load the image
	image := gocv.IMRead(imagePath, gocv.IMReadColor)
	if image.Empty() {
		fmt.Printf("Error reading image from: %s\n", imagePath)
		return
	}
	defer image.Close()

	// Load the pre-trained model
  modelPath := "./pre-trained-model/mobilenet_v2_1.0_224.pb"
	net := gocv.ReadNet(modelPath) // Replace with the path to your pre-trained model file
	if net.Empty() {
		fmt.Printf("Error reading model from: %s\n", modelPath)
		return
	}
	defer net.Close()

	// Preprocess the image
	blob := gocv.BlobFromImage(image, 1.0, image.Pt(224, 224), gocv.NewScalar(0, 0, 0, 0), true, false)
	defer blob.Close()

	// Set the input for the network
	net.SetInput(blob, "data")

	// Forward pass through the network
	prob := net.Forward("prob")

	// Get the top 5 predictions
	_, maxVal, _, maxLoc := gocv.MinMaxLoc(prob)
	fmt.Printf("Predicted class: %v, confidence: %.2f\n", maxLoc.X, maxVal)

	// Cleanup
	prob.Close()
}
