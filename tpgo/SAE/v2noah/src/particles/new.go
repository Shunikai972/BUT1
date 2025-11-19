package particles

import (
	"container/list"
	"math/rand/v2"
	"project-particles/config"
)

// config.General.WindowSizeX
// config.General.WindowSizeY
// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()
	for i := 0; i < config.General.InitNumParticles; i++ {

		var posX, posY float64

		maxX := float64(config.General.WindowSizeX) - (10*config.General.ScaleX) //*10 pour prendre en compte la taille du .png
		maxY := float64(config.General.WindowSizeY) - (10*config.General.ScaleY)

		if config.General.RandomSpawn {
			posX = rand.Float64()*(maxX)
			posY = rand.Float64()*(maxY)
		} else {
			posX = float64(config.General.SpawnX)
			posY = float64(config.General.SpawnY)
		}

		l.PushFront(&Particle{
			ScaleX:     config.General.ScaleX,
			ScaleY:     config.General.ScaleY,
			PositionX:  posX,
			PositionY:  posY,
			ColorRed:   rand.Float64(),
			ColorGreen: rand.Float64(),
			ColorBlue:  rand.Float64(),
			Opacity:    1,
			SpeedX:     rand.Float64()*2 - 1, //Vitesse random entre [-1,1]
			SpeedY:     rand.Float64()*2 - 1,
		})
	}
	return System{Content: l}
}
