package assets

import (
	"log"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	executablePath   string
	EnergyDrinkAsset *ebiten.Image
)

func Init() {
	// Initialize executablePath
	var err error
	executablePath, err = os.Getwd()
	if err != nil {
		log.Fatalf("Error getting executable path: %v", err)
	}

	// Get the directory of the executable
	executableDir := filepath.Dir(executablePath)

	// Construct the path to the asset
	assetPath := filepath.Join(executableDir, "gosnake/assets/energy_drink.png")

	// Load the image
	var imgErr error
	EnergyDrinkAsset, _, imgErr = ebitenutil.NewImageFromFile(assetPath)
	if imgErr != nil {
		log.Fatalf("Error loading image: %v", imgErr)
	}
}
