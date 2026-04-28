# TD3: Gestion des événements



Nous allons dans ce TD mettre en place des gestionnaires d’événements. 

Le code de la vue de l’application vous est fourni dans le fichier *TD3.kt*. Pour chaque question, il faudra développer des écouteurs dans des fichiers séparés (stockés dans le répertoire *ecouteurs* et ensuite abonner le nœud concerné de la vue (objet graphique ou conteneur) à cet écouteur.

![exo1.png](image%2Fexo1.png)

### 1) Clic sur le bouton *go*

Lorsque l’utilisateur clique sur le bouton *go*, il faut que le label nommé *labelNbClicBouton* reflète le nombre de clic déjà réalisé sur ce bouton.

Vous développerez l’écouteur dans le fichier *BoutonGoEcouteur.kt*. Il est déjà partiellement écrit.

**Explications:**

La classe *BoutonGoEcouteur* implémente l’interface *EventHandler<ActionEvent>*

Pour connaître le type de l‘événement que devra gérer l’écouteur, il suffit de regarder dans le tableau à la fin du cours où sont répertoriés, les actions utilisateurs, les composants cibles et les événements associés.
Le clic sur un bouton génère un évènement de type *ActionEvent*

La méthode *handle(…)* de cette classe, que vous devez redéfinir, est la méthode déclenchée lorsque l’événement a lieu.

Vous remarquerez que la classe *BoutonGoEcouteur* a un constructeur qui a comme paramètre, la vue de l’application. Dans cette vue, certains des nœuds de la scène sont en attributs. Suivant la visibilité autorisée, on pourra donc dans l’écouteur, via la vue, manipuler ses attributs (changer la valeur de la propriété d’un nœud liée à la couleur, le contenu d’un label …)

Une fois la méthode *handle()* redéfinie, il ne reste plus qu’à associer l’écouteur au nœud concerné dans la vue.

> A vous de jouer ! Développez le premier écouteur

### 2) Choix d'une couleur dans la *ComboBox*

Lorsque l’utilisateur choisit une couleur dans la *ComboBox choixCouleur*, le panneau *panneauCouleurs* se colore de la couleur choisie. Il faudra utiliser l’attribut *selectionModel* de la *comboBox* pour pouvoir accéder à l’item sélectionné

**Ecouteur** => ***ChoixCouleurEcouteur***


### 3) Clic sur le panneau *panneauCouleurs*

Lorsque l’utilisateur clique sur le panneau *panneauCouleurs*, il faut que le label nommé *labelNbClicPanneau* reflète le nombre de clic déjà réalisé sur ce panneau

**Ecouteur** => ***ClicPanneauEcouteur***

> Pour ceci, il faut regarder la classe [*MouseEvent*](https://docs.oracle.com/javase/8/javafx/api/javafx/scene/input/MouseEvent.html)

### 4) Afficher/cacher une image

- quand on clique sur le bouton "afficher image", l'image *image/fleur.jpg* s'affiche dans le *panneauBas* après les 2 entiers. Le texte du bouton est modifié, il devient "cacher image".
- quand on clique sur le bouton "cacher image", l'image disparait et le bouton reprend son texte d'origine

**Ecouteur** => ***BoutonAfficheImageEcouteur***


### 5) Recopie du texte à la volée

On veut maintenant que le texte saisi dans le *TextField zoneTexte* soit recopié à la volée dans le *TextArea textarea*.

**Ecouteur** => ***RecopieurTexteEcouteur***

> Pour ceci, il faut regarder la classe [*KeyEvent*](https://docs.oracle.com/javase/8/javafx/api/javafx/scene/input/KeyEvent.html)

### 6) Deux clics sur le champs texte *zoneTexte*

Lorqu’on clique 2 fois sur le champ texte *zoneTexte* alors son contenu s’efface

**Ecouteur** => ***EffaceurTexteEcouteur***

> On voit ici qu'un nœud peut être abonné à plusieurs écouteurs


### 7) Fenêtre de dialogue

On veut maintenant, que si dans le *TextField zoneTexte* un caractère *\$* est saisi, alors une fenêtre de dialogue de type *ALERT* s’ouvre en stipulant le problème et le caractère *\$* est effacé de la zone de saisie


### 8) Entrée du pointeur de souris dans *panneauHaut*

On veut maintenant que lorsque le pointeur de souris entre dans *panneauHaut*, le panneau se colore en rose et lorsqu’il en sort, le panneau reprend sa couleur initiale.

> Pour ceci, il faut regarder la classe [*MouseEvent*](https://docs.oracle.com/javase/8/javafx/api/javafx/scene/input/MouseEvent.html)

> On peut accéder au " type " de l’événement *MouseEvent* via la méthode *getEventType()*. Elle renvoie un des types d’événement qui sont des variables "statiques" de la classe.
> Par exemple: *MouseEvent.MOUSE_PRESSED*

**Ecouteur**: ***PanneauHautEcouteur***



### 9) Utilisation de lambda expression

Dans la vue, désabonnez le bouton *go* de son écouteur.  Maintenant, obtenez le même comportement  en utilisant une lambda expression (voir cours).


### 10) Appui sur la touche ctrl

On veut maintenant que lorsque l’utilisateur appuie sur la touche *ctrl* de son clavier le *panneauBas* se colore en rose. Lorsqu’il relâche la touche, le *panneauBas* doit redevenir blanc. Définir où placer l’écouteur, quel type d’événement intercepter et les traitements à effectuer. Implémentez ceci.

**Ecouteur**: ***EcouteurToucheControl***



