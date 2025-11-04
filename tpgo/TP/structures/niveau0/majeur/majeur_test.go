package majeur

import "testing"

func Test0Etudiant(t *testing.T) {
	p := personne{
		prenom: "ptest",
		nom:    "ntest",
		age:    20,
	}
	p.verifieMajeur()
	if !p.estMajeur {
		t.Fail()
	}
}

func Test1Etudiant(t *testing.T) {
	p := personne{
		prenom: "ptest",
		nom:    "ntest",
		age:    10,
	}
	p.verifieMajeur()
	if p.estMajeur {
		t.Fail()
	}
}

// Ajoutés après le test

type test struct {
	p         personne
	estMajeur bool
}

var testSet []test = []test{
	{p: personne{prenom: "ptest", nom: "ntest", age: 20}, estMajeur: true},
	{p: personne{prenom: "ptest", nom: "ntest", age: 10}, estMajeur: false},
	{p: personne{prenom: "Bob", nom: "Bobob", age: 18}, estMajeur: true},
}

func TestGeneral(t *testing.T) {
	for numTest, aTester := range testSet {
		p := aTester.p
		p.verifieMajeur()
		if aTester.estMajeur != p.estMajeur || aTester.p.nom != p.nom || aTester.p.prenom != p.prenom || aTester.p.age != p.age {
			t.Error("Erreur sur l'entrée", numTest, "de testSet")
		}
	}
}
