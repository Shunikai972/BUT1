package ihm.td5.modele

import ihm.td5.librairie.Auteur
import ihm.td5.librairie.Livre



/**
 * La bibliothèque est une classe qui combine un index de livres et un index
 * d'auteurs
 *
 * @author lanoix-a
 * @version 2019
 * réécrit en kotlin en 2022 et modifié par jacquin-c
 */

class Bibliotheque {


    private val livres: MutableList<Livre>
    var courant: Int
    private var nbLivres: Int
    private val auteurs : MutableSet<Auteur>


    /* créer une bibliotheque de livres
    */
    init {
        this.livres = ArrayList<Livre>()
        this.courant = 0
        this.nbLivres = -1
        this.auteurs = HashSet<Auteur>()
    }

    /**retourne le nombre de livre dans la  bibliothèque
     *@return le nombrede livre dans la bibliothèque
     *
     */
    fun donneNbLivres(): Int{
        return this.nbLivres
    }


    /**retourne le livre à l'indice index dans la bibliothèque de livres
     *@param index du livre recherché
     *@return le livre recherché
     * @throws NoSuchElementException si le livre à l'indice index n'existe pas
     */
    fun donneLivre(index: Int): Livre {
        if (index  >=0 && index < this.livres.size)
            return  this.livres[index]
        else
            throw NoSuchElementException()
    }

    /**
     * ajoute le livre à la bibliothèque de livres ; ajoute également l'auteur
     * du livres à l'index des auteurs
     *
     * @param livre livre à ajouter à la bibliothèque
     */
    fun ajoutLivre(livre: Livre) {
        if (nbLivres < this.livres.size) {
            this.livres.add(livre)
            this.nbLivres++
        }
        ajoutAuteur(livre.auteur)
    }


    /**
     * supprime le livre courant de la bibliothèques de livres ;
     * se repositionne sur le 1er livre de la bibliotèque
     */
    fun suppressionLivre() {

        this.livres.removeAt(this.courant)
        this.courant = 0
        this.nbLivres--
    }

    /**modifie le livre courant de la bibliothèques de livres
     *
     */
   fun modifierLivre(livre: Livre) {
    this.livres[this.courant]=livre
   }


/*
Teste si la liste de livres est vide

 */

    fun estVide(): Boolean{
        if (this.nbLivres<0){
            return true}
        else{
            return false
        }
    }




    /**
     * @return le livre actuellement consulté
     */
    fun donneLivre(): Livre {
        return this.livres[this.courant]
    }


    /**
     * change le livre courant par son sucesseur
     *
     * @throws NoSuchElementException si le livre courant n'a pas de sucesseur
     */
    fun passeAuLivreSuivant()  {
        if (ilYaLivreSuivant())
            this.courant++
        else
            throw NoSuchElementException()
    }




    /**
     * change le livre courant par son prédécésseur
     *
     * @throws NoSuchElementException si le livre courant n'a pas de prédécesseur
     */

    fun passeAuLivrePrecedent() {
        if (ilYaLivrePrecedent())
            this.courant--
        else
            throw NoSuchElementException()
    }

    /**
     * @return vrai s'il y a un livre avant le livre courant dans l'index des
     * livres
     */
    fun ilYaLivrePrecedent(): Boolean {
        return this.courant > 0
    }
    /**
     * @return vrai s'il y a encore un livre après le livre courant dans l'index
     * des livres
     */
    fun ilYaLivreSuivant(): Boolean {
        return this.courant < this.livres.size - 1
    }


    /**
     * @return les livres disponibles dans la bibliothèque
     */
    fun donneTousLesLivres(): ArrayList<Livre> {
        return ArrayList(this.livres)
    }

    /**
     * @param auteur auteur à ajouter à l'index des auteurs
     * @return vrai si l'ajout a bien eu lieu
     */
    fun ajoutAuteur(auteur: Auteur): Boolean {
        return this.auteurs.add(auteur)
    }

    /**
     * @return les auteurs disponibles dans la bibliothèque
     */
    fun donneTousLesAuteurs(): ArrayList<Auteur> {
        return ArrayList<Auteur>(this.auteurs)
    }

    /**
     * donne l'auteur corresplivreSuivantondant à la chaine de caractères
     *
     * @param str une chaine de caractère
     * @return l'auteur correspondant à la chaine de caractères
     * @throws NoSuchElementException si la chaine de caractères ne correspond à aucun auteur dans
     *                   l'index des auteurs
     */
    fun donneAuteur(str: String): Auteur {
        for (auteur in this.auteurs){
           if(str==auteur.toString())
               return auteur
        }
        throw NoSuchElementException()
    }

    /**
     * remplit la bibliothèque avec des données de tests
     */
    fun preremplir() {
        ajoutLivre(Livre("L'Appel de Cthulhu", Livre.HORROR, Auteur("Lovecraft", "Howard Phillips")))
        ajoutLivre(Livre("Les lames du Cardinal", Livre.FANTASY, Auteur("Pevel", "Pierre")))
        val pratchett = Auteur("Pratchett", "Terry")
        ajoutLivre(Livre("Mortimer", Livre.FANTASY, pratchett))
        ajoutLivre(Livre("Procrastination", Livre.FANTASY, pratchett))
        ajoutLivre(Livre("Les ch'tits hommes libres", Livre.FANTASY, pratchett))
        ajoutLivre(Livre("La longue terre", Livre.SF, pratchett))
        val simmons = Auteur("Simmons", "Dan")
        ajoutLivre(Livre("Hyperion", Livre.SF, simmons))
        ajoutLivre(Livre("L'echiquier du mal", Livre.HORROR, simmons))
        ajoutLivre(Livre("Gagner la guerre", Livre.FANTASY, Auteur("Jaworski", "Jean-Philippe")))
    }

}
