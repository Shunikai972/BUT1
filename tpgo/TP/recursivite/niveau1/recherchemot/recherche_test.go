package recherche

import "testing"

func TestVideEtudiant(t *testing.T) {
	if rechercheMot([]string{}, "bonjour") {
		t.Fail()
	}
}

func TestExisteEtudiant(t *testing.T) {
	if !rechercheMot([]string{"salut", "bonjour", "bonsoir"}, "bonjour") {
		t.Fail()
	}
}

func TestNExistePasEtudiant(t *testing.T) {
	if rechercheMot([]string{"salut", "bon", "jour", "bonsoir"}, "bonjour") {
		t.Fail()
	}
}

// Ajouté après le DS

func TestGeneral(t *testing.T) {
	dico := []string{"Une", "pierre", "deux", "maisons", "trois", "ruines", "quatre", "fossoyeurs", "un", "jardin", "des", "fleurs", "un", "raton", "laveur", "une", "douzaine", "d’huîtres", "un", "citron", "un", "pain", "un", "rayon", "de", "soleil", "une", "lame", "de", "fond", "six", "musiciens", "une", "porte", "avec", "son", "paillasson", "un", "monsieur", "décoré", "de", "la", "légion", "d’honneur", "un", "autre", "raton", "laveur", "un", "sculpteur", "qui", "sculpte", "des", "napoléon", "la", "fleur", "qu’on", "appelle", "souci", "deux", "amoureux", "sur", "un", "grand", "lit", "un", "receveur", "des", "contributions", "une", "chaise", "trois", "dindons", "un", "ecclésiastique", "un", "furoncle", "une", "guêpe", "un", "rein", "flottant", "une", "écurie", "de", "courses", "un", "fils", "indigne", "deux", "frères", "dominicains", "trois", "sauterelles", "un", "strapontin", "deux", "filles", "de", "joie", "un", "oncle", "Cyprien", "une", "Mater", "dolorosa", "trois", "papas", "gâteau", "deux", "chèvres", "de", "Monsieur", "Seguin", "un", "talon", "Louis", "XV", "un", "fauteuil", "Louis", "XVI", "un", "buffet", "Henri", "II", "deux", "buffets", "Henri", "III", "trois", "buffets", "Henri", "IV", "un", "tiroir", "dépareillé", "une", "pelote", "de", "ficelle", "deux", "épingles", "de", "sûreté", "un", "monsieur", "âgé", "une", "Victoire", "de", "samothrace", "un", "comptable", "deux", "aides-", "comptables", "un", "homme", "du", "monde", "deux", "chirurgiens", "trois", "végétariens", "un", "cannibale", "une", "expédition", "coloniale", "un", "cheval", "entier", "une", "demi-", "pinte", "de", "bon", "sang", "une", "mouche", "tsé-tsé", "un", "homard", "à", "l’américaine", "un", "jardin", "à", "la", "française", "deux", "pommes", "à", "l’anglaise", "un", "face-à-main", "un", "valet", "de", "pied", "un", "orphelin", "un", "poumon", "d’acier", "un", "jour", "de", "gloire", "une", "semaine", "de", "bonté", "un", "mois", "de", "marie", "une", "année", "terrible", "une", "minute", "de", "silence", "une", "seconde", "d’inattention", "et…", "cinq", "ou", "six", "ratons", "laveurs", "un", "petit", "garçon", "qui", "entre", "à", "l’école", "en", "pleurant", "un", "petit", "garçon", "qui", "sort", "de", "l’école", "en", "riant", "une", "fourmi", "deux", "pierres", "à", "briquet", "dix-sept", "éléphants", "un", "juge", "d’instruction", "en", "vacances", "assis", "sur", "un", "pliant", "un", "paysage", "avec", "beaucoup", "d’herbe", "verte", "dedans", "une", "vache", "un", "taureau", "deux", "belles", "amours", "trois", "grandes", "orgues", "un", "veau", "marengo", "un", "soleil", "d’austerlitz", "un", "siphon", "d’eau", "de", "Seltz", "un", "vin", "blanc", "citron", "un", "Petit", "Poucet", "un", "grand", "pardon", "un", "calvaire", "de", "pierre", "une", "échelle", "de", "corde", "deux", "sœurs", "latines", "trois", "dimensions", "douze", "apôtres", "mille", "et", "une", "nuits", "trente-deux", "positions", "six", "parties", "du", "monde", "cinq", "points", "cardinaux", "dix", "ans", "de", "bons", "et", "loyaux", "services", "sept", "péchés", "capitaux", "deux", "doigts", "de", "la", "main", "dix", "gouttes", "avant", "chaque", "repas", "trente", "jours", "de", "prison", "dont", "quinze", "de", "cellule", "cinq", "minutes", "d’entracte", "et…", "plusieurs", "ratons", "laveurs."}

	autres := []string{"Cabbage", "Rolls", "Cakes", "Calzones", "Camping", "Recipes", "Canning", "and", "Preserving", "Carrot", "Cakes", "Casseroles", "Ceviche", "Cheese", "Balls", "Cheese", "Fondue", "Cheesecakes", "Chef", "John", "Cherry", "Pie", "Chess", "Pie", "Chicken", "Adobo", "Chicken", "and", "Dumplings", "Chicken", "Cacciatore", "Chicken", "Cordon", "Bleu", "Chicken", "Marsala", "Chicken", "Noodle", "Soups", "Chicken", "Parmesan", "Chicken", "Piccata", "Chicken", "Salads", "Chicken", "Teriyaki", "Chilaquiles", "Chiles", "Rellenos", "Chili", "Recipes", "Chinese", "Chocolate", "Cakes", "Chocolate", "Chip", "Cookies", "Chocolate", "Fudge", "Chowders", "Christmas", "Christmas", "Cookies", "Cinnamon", "Rolls", "Cobblers", "Cocktails", "Coffee", "Cakes", "Coleslaws", "Comfort", "Food", "Cookies", "Cooking", "for", "a", "Crowd", "Cooking", "for", "One", "Cooking", "for", "Two", "Copycat", "Recipes", "Cornbread", "Corned", "Beef", "Crab", "Cakes", "Crackers", "Cranberry", "Sauces", "Creme", "Brulee", "Crisps", "and", "Crumbles", "Cupcakes"}

	// dans le dico
	for i := 0; i < len(dico); i++ {
		mot := dico[i]
		if !rechercheMot(dico, mot) {
			t.Errorf("Vous dites que \"%s\" n'est pas dans le dico", mot)
		}
	}

	// pas dans le dico
	for i := 0; i < len(autres); i++ {
		if rechercheMot(dico, autres[i]) {
			t.Errorf("Vous dites que \"%s\" est dans le dico", autres[i])
		}
	}
}
