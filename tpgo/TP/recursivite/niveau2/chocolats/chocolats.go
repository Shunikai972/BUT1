package chocolats

/*
Une marque de barres de chocolat fait une promotion~: si on retourne k embalages on en obtient une gratuite. On se demande alors combien de barres de chocolat on peut obtenir quand on dispose de n euros et que chacune coûte m euros.
La fonction miam doit répondre (de manière récursive) à cette question.

# Entrées
- n : la somme dont on dispose en euros
- m : le prix d'une barre de chocolat en euros, toujours supérieur à 0
- k : le nombre d'embalages qu'il faut pour obtenir une barre de chocolat gratuite, toujours supérieur à 1

# Sortie
- choco : le nombre de barre de chocolat qu'on peut obtenir en tout

# Exemple
*/
func miam(n, m, k uint) (choco uint) {
	return choco
}


package chocolats

/*
Une marque de barres de chocolat fait une promotion : si on retourne k emballages on en obtient une gratuite.
On se demande alors combien de barres de chocolat on peut obtenir quand on dispose de n euros et que chacune coûte m euros.
La fonction miam doit répondre (de manière récursive) à cette question.

# Entrées
- n : la somme dont on dispose en euros
- m : le prix d'une barre de chocolat en euros, toujours supérieur à 0
- k : le nombre d'emballages qu'il faut pour obtenir une barre de chocolat gratuite, toujours supérieur à 1

# Sortie
- choco : le nombre de barres de chocolat qu'on peut obtenir en tout
*/
func miam(n, m, k uint) (choco uint) {
    // Si l'argent disponible (n) ne permet même pas d'acheter une barre (m),
    // on ne peut rien obtenir. C'est le cas de base (ou cas d'arrêt) de la récursion.
    if n < m {
        return 0
    }

    // --- Étape 1 : Achat initial ---
    // Nombre de barres que l'on peut acheter avec l'argent initial (n).
    barresAchetees := n / m

    // --- Étape 2 : Récursion pour la promotion ---
    // Pour chaque barre achetée, on a 1 emballage.
    // On passe le nombre d'emballages disponibles pour le premier échange promotionnel.
    barresPromo := echangeEmballages(barresAchetees, k)

    // --- Étape 3 : Total ---
    // Le total est la somme des barres achetées initialement et des barres obtenues par promotion.
    return barresAchetees + barresPromo
}

// Fonction récursive auxiliaire pour gérer l'échange d'emballages.
// Elle calcule le nombre total de barres gratuites obtenues.
func echangeEmballages(emballagesDisponibles, k uint) (barresGratuites uint) {
    // Cas de base (ou cas d'arrêt) :
    // S'il n'y a pas assez d'emballages pour obtenir une barre gratuite (i.e., < k),
    // on arrête la récursion.
    if emballagesDisponibles < k {
        return 0
    }

    // --- Calcul des barres gratuites obtenues à cette étape ---
    // Nombre de barres gratuites obtenues en échangeant les emballages.
    barresGratuitesObtenues := emballagesDisponibles / k

    // --- Calcul des nouveaux emballages disponibles pour la prochaine étape ---
    // 1. Les emballages *restants* de l'échange actuel (le reste de la division).
    emballagesRestants := emballagesDisponibles % k
    
    // 2. Les emballages *générés* par les nouvelles barres gratuites (chaque barre gratuite donne 1 emballage).
    emballagesNouveaux := barresGratuitesObtenues

    // Le total d'emballages pour le prochain tour de promotion.
    nouveauxEmballages := emballagesRestants + emballagesNouveaux

    // --- Appel récursif ---
    // On rappelle la fonction avec le total de nouveaux emballages.
    // Cela nous donne le nombre de barres gratuites obtenues lors des échanges suivants.
    barresGratuitesRecursives := echangeEmballages(nouveauxEmballages, k)

    // --- Total des barres gratuites ---
    // Le total est la somme des barres obtenues à cette étape *plus* toutes celles obtenues par récursion.
    return barresGratuitesObtenues + barresGratuitesRecursives
}