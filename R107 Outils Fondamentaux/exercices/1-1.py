import math
def modulo(a,b): 
    reste = a  
    quotient = 0
    while reste >= b:
        reste = reste - b 
        quotient = quotient + 1  
    print("Quotient :", quotient)
    print("Reste :", reste)
    return(quotient, reste)

def euclide_etendu(a, b):
    u0, u1 = 1, 0
    v0, v1 = 0, 1
    
    while b != 0:
        quotient, reste = modulo(a, b)
        
        u2 = u0 - quotient * u1
        v2 = v0 - quotient * v1
        u0, u1 = u1, u2
        v0, v1 = v1, v2
        a, b 
        b = reste  
    return a, u0, v0

pgcd, u, v = euclide_etendu(120, 23)
print("PGCD :", pgcd)
print("u =", u, "v =", v)
print("Vérification : au + bv =", 120*u + 23*v)



def est_premier(p):
    if p < 2:
        return False
    i = 2
    while i * i <= p:
        if p % i == 0:
            return False
        i += 1
    return True
 
def n_eme_premier(n):
    compteur = 0
    nombre = 1
    while compteur < n:
        nombre += 1
        if est_premier(nombre):
            compteur += 1
    return nombre
#EXO5
import numpy as np
np.zeros(7)          # crée un tableau de 7 éléments, tous égaux à 0

np.ones(6)           # crée un tableau de 6 éléments, tous égaux à 1

np.identity(3)       # crée une matrice identité de taille 3×3 
                     # (diagonale = 1, le reste = 0)

np.array([3,7,-1,2]) # crée un tableau 1D contenant les valeurs 3, 7, -1 et 2

np.array([[3,7],[-1,2]])  # crée un tableau 2D (matrice 2×2) avec ces lignes

np.arange(10,30,5)   # crée un tableau allant de 10 à 30 (exclu) avec un pas de 5
                     # → [10, 15, 20, 25]

np.linspace(0,2,9)   # crée 9 valeurs également espacées entre 0 et 2 (inclus)

np.sin(np.linspace(0,2*np.pi,20))  
                     # calcule le sinus de 20 valeurs régulièrement espacées
                     # entre 0 et 2π (une période complète)
#EX5.2
import numpy as np

a = np.array([1, 2, 3])
b = np.array([10, 20, 30])

print(a + b)   # [11 22 33]



a = np.array([1, 2, 3])
print(a + 4)   # [5 6 7]


a = np.array([1, 2, 3])
b = np.array([10, 20, 30])

print(a * b)   # [10 40 90]



a = np.array([1, 2, 3])
print(3 * a)   # [3 6 9]


print(np.add(a, b))  # [11 22 33]


A = np.array([[1, 2],
              [3, 4]])

B = np.array([[10, 20],
              [30, 40]])

print(A.dot(B))
# [[ 70 100]
#  [150 220]]

print(A @ B)
# [[ 70 100]
#  [150 220]]


A = np.array([[1, 2, 3],
              [4, 5, 6]])

print(A.transpose())
# [[1 4]
#  [2 5]
#  [3 6]]



from numpy.linalg import matrix_power

A = np.array([[1, 2],
              [3, 4]])

print(matrix_power(A, 2))
# [[ 7 10]
#  [15 22]]




from numpy.linalg import det

A = np.array([[1, 2],
              [3, 4]])

print(det(A))   # -2.0000000000000004



from numpy.linalg import inv

print(inv(A))
# [[-2.   1. ]
#  [ 1.5 -0.5]]





A = np.array([[1, 2, 3],
              [4, 5, 6]])

print(A.shape)   # (2, 3)

#5.3
"""
Opérations vectorielles ou matricielles classiques :

a + b

3 * a et a * 3 (multiplication par un scalaire)

a.dot(b) (produit scalaire ou produit matriciel)

a @ b (produit matriciel)

a.transpose() (transposée)

np.linalg.matrix_power(a, 2) (puissance d’une matrice)

np.linalg.det(a) (déterminant)

np.linalg.inv(a) (inverse d’une matrice)

⚠️ Peu ou pas définies en algèbre linéaire :

a + 4 (ajout d’un scalaire à une matrice → pas d’interprétation en algèbre linéaire)

a * b (multiplication terme à terme, non linéaire)

np.add(a, b) (même remarque que a + b, mais conceptuellement oui en algèbre linéaire : somme de matrices)

a.shape (simple information de structure) """

def diagonale(A):
    """
    Prend en entrée une matrice A (liste de listes) de taille p × n
    et retourne une matrice B ne conservant que les éléments diagonaux.
    """
    p = len(A)
    n = len(A[0]) if p > 0 else 0

    B = []
    for i in range(p):
        ligne = []
        for j in range(n):
            if i == j:
                ligne.append(A[i][j])
            else:
                ligne.append(0)
        B.append(ligne)
    return B
