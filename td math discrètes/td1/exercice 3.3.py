def carre (liste):
    squares = []
    for element in liste:
        squares.append(element*element)
    return squares
ma_liste = [1, 2, 3, 2, 4, 2, 5]
resultat = carre(ma_liste)
print(resultat) 
