package ateddybearpicnic

import (
	"testing"
)

func Test42Etudiant(t *testing.T) {
	if !bears(42) {
		t.Error("Vous dites qu'avec un nombre d'ours initial de 42 il n'est pas possible de gagner la partie alors qu'elle est directement gagnée")
	}
}

func Test250Etudiant(t *testing.T) {
	if !bears(250) {
		t.Error("Vous dites qu'avec un nombre d'ours initial de 250 il n'est pas possible de gagner la partie alors que c'est possible avec la séquence suivante : 250 est divisible par 5, se débarasser de 42 ours pour arriver à 208. Ensuite 208 est pair, se débarasser de la moitié des ours pour arriver à 104. Puis 104 est pair, se débarasser de la moitié des ours pour arriver à 52. Enfin, 52 est divisible par 4, multiplier les deux derniers chiffres donne 10, ce qui permet de se débarasser de 10 ours pour arriver à 42.")
	}
}

func Test53Etudiant(t *testing.T) {
	if bears(53) {
		t.Error("Vous dites qu'avec un nombre d'ours initial de 53 il est possible de gagner la partie, ce qui n'est pas vrai.")
	}
}

// Ajouté après le DS

type test struct {
	numReceived uint
	winnable    bool
}

var testSet []test = []test{
	{numReceived: 41, winnable: false},
	{numReceived: 42, winnable: true},
	{numReceived: 53, winnable: false},
	{numReceived: 76, winnable: false},
	{numReceived: 84, winnable: true},
	{numReceived: 208, winnable: true},
	{numReceived: 250, winnable: true},
	{numReceived: 533, winnable: false},
	{numReceived: 843, winnable: true},
	{numReceived: 941, winnable: false},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		if bears(aTester.numReceived) != aTester.winnable {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
