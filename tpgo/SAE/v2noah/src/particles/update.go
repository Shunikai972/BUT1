package particles

import (
	"math/rand/v2"
	"project-particles/config"
)

// Update mets à jour l'état du système de particules (c'est-à-dire l'état de
// chacune des particules) à chaque pas de temps. Elle est appellée exactement
// 60 fois par seconde (de manière régulière) par la fonction principale du
// projet.
// C'est à vous de développer cette fonction.
func (s *System) Update() {
	for l := s.Content.Front(); l != nil; l = l.Next() {
		p := l.Value.(*Particle)

		p.PositionX -= p.SpeedX * config.General.SpeedFactor
		p.PositionY -= p.SpeedY * config.General.SpeedFactor

		if p.PositionX < 0 || p.PositionX > float64(config.General.WindowSizeX)-(10*p.ScaleX) || p.PositionY < 0 || p.PositionY > float64(config.General.WindowSizeY)-(10*p.ScaleY) {
			p.PositionX = rand.Float64() * (float64(config.General.WindowSizeX) - (10*p.ScaleX))
			p.PositionY = rand.Float64() * (float64(config.General.WindowSizeY) - (10*p.ScaleY))
		}
	}
}
