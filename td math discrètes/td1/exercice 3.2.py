def compte (el, liste):
    compteur = 0
    for element in liste:
        if element == el:
            compteur += 1
    return compteur

ma_liste = [1, 2, 3, 2, 4, 2, 5]
resultat = compte(2, ma_liste)
print(resultat)
