package particles

import (
	"container/list"
	"project-particles/config"
	"math/rand/v2"
)

// NewSystem est une fonction qui initialise un système de particules et le
// retourne à la fonction principale du projet, qui se chargera de l'afficher.
// C'est à vous de développer cette fonction.
// Dans sa version actuelle, cette fonction affiche une particule blanche au
// centre de l'écran.
func NewSystem() System {
	l := list.New()
for i := 0; i < config.General.InitNumParticles; i++{
	if config.General.RandomSpawn{
		l.PushFront(&Particle{
			ScaleX:    1, ScaleY: 1,
			PositionX: (rand.Float64() * float64(config.General.WindowSizeX))- config.General.ScaleX,
			PositionY: (rand.Float64() * float64(config.General.WindowSizeY)) - config.General.ScaleY,
			
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
		})
	} else{
		l.PushFront(&Particle{
			PositionX: float64(config.General.WindowSizeX),
			PositionY: float64(config.General.WindowSizeY),
			ScaleX:    1, ScaleY: 1,
			ColorRed: 1, ColorGreen: 1, ColorBlue: 1,
			Opacity: 1,
		})
	}
}	
	return System{Content: l}
}
