package particles_test

import (
	"testing"
	"project-particles/particles"
	"project-particles/config"
)

func TestNewSystem_NoRandomSpawn(t *testing.T) {
	config.General.InitNumParticles = 1
	config.General.RandomSpawn = false
	config.General.SpawnX = 100
	config.General.SpawnY = 200

	system := particles.NewSystem()

	if system.Content.Len() != 1 {
		t.Fatalf("expected 1 particles, got %d", system.Content.Len())
	}

	// Test si elles sont bien à la bonne position
	p := system.Content.Front().Value.(*particles.Particle)

	if p.PositionX != 100 || p.PositionY != 200 {
		t.Errorf("unexpected particle position (%f, %f)", p.PositionX, p.PositionY)
	}
}

func TestNewSystem_InitNumParticles(t *testing.T) {
	config.General.InitNumParticles = 10

	system := particles.NewSystem()

	//Test si il y en a 10
	if system.Content.Len() != 10 {
		t.Fatalf("expected 10 particles, got %d", system.Content.Len())
	}
}