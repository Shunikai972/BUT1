import math
import numpy as np
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


def transpose(A):
    p = len(A)
    n = len(A[0])
    B = []

    for j in range(n):
        ligne = []
        for i in range(p):
            ligne.append(A[i][j])
        B.append(ligne)
    return B


def produit(A, B):
    p = len(A)
    l = len(A[0])
    
    if len(B) != l:
        return []
    
    n = len(B[0])
    
    C = [[0 for _ in range(n)] for _ in range(p)]
    
    for i in range(p):
        for j in range(n):
            for k in range(l):
                C[i][j] += A[i][k] * B[k][j]
    
    return C


def vandermonde(v):
    n = len(v)
    A = []
    for i in range(n):
        ligne = []
        for j in range(n):
            ligne.append(v[i] ** j)
        A.append(ligne)
    return A


def permute_colonnes(A):
    p = len(A)
    n = len(A[0])
    B = []
    for i in range(p):
        ligne = []
        for j in range(n):
            j_prime = n - 1 - j
            ligne.append(A[i][j_prime])
        B.append(ligne)
    return B


def permute_lignes(A):
    p = len(A)
    n = len(A[0])
    B = []
    for i in range(p):
        i_prime = p - 1 - i
        B.append(A[i_prime][:])
    return B


def minimum(A):
    p = len(A)
    n = len(A[0])
    min_val = A[0][0]
    i_min = 0
    j_min = 0

    for i in range(p):
        for j in range(n):
            if A[i][j] < min_val:
                min_val = A[i][j]
                i_min = i
                j_min = j
    
    return min_val, (i_min, j_min)

# ----------------------------
#  TESTS + PRINTS
# ----------------------------

print("\n===== TEST modulo =====")
print(modulo(10, 3))
print(modulo(25, 5))
print(modulo(7, 10))
print(modulo(0, 4))


print("\n===== TEST euclide_etendu =====")
pgcd, u, v = euclide_etendu(120, 23)
print("PGCD =", pgcd, "| u =", u, "v =", v)
print("Vérification :", 120*u + 23*v)

pgcd, u, v = euclide_etendu(56, 98)
print("PGCD =", pgcd, "| u =", u, "v =", v)
print("Vérification :", 56*u + 98*v)


print("\n===== TEST est_premier =====")
for n in [1,2,3,4,5,97,100]:
    print(n, ":", est_premier(n))


print("\n===== TEST n_eme_premier =====")
for n in [1,2,5,10]:
    print("Premier numéro", n, "=", n_eme_premier(n))


print("\n===== TEST diagonale =====")
A = [[1,2,3],[4,5,6],[7,8,9]]
print(diagonale(A))


print("\n===== TEST transpose =====")
A = [[1,2],[3,4],[5,6]]
print("transpose :", transpose(A))


print("\n===== TEST produit =====")
A = [[1,2],[3,4]]
B = [[10,20],[30,40]]
print(produit(A, B))


print("\n===== TEST vandermonde =====")
print(vandermonde([1,2,3]))


print("\n===== TEST permute_colonnes =====")
A = [[1,2,3],[4,5,6]]
print(permute_colonnes(A))


print("\n===== TEST permute_lignes =====")
A = [[1,2],[3,4],[5,6]]
print(permute_lignes(A))


print("\n===== TEST minimum =====")
A = [[3,5,9],[8,-2,7],[4,1,0]]
print(minimum(A))


def permutation(A, y, i, k):
    A2 = A.copy()  
    y2 = y.copy()

    A2[i], A2[k] = A2[k].copy(), A2[i].copy()
    y2[i], y2[k] = y2[k], y2[i]

    return A2, y2


def elimination(A, y, i, j):
    A2 = A.copy()
    y2 = y.copy()

    pivot = A2[i, j]

    for k in range(i+1, A2.shape[0]):
        facteur = A2[k, j] / pivot
        A2[k, :] = A2[k, :] - facteur * A2[i, :]
        y2[k] = y2[k] - facteur * y2[i]

    return A2, y2

A = np.array([[1, 1, 1],
              [1, -1, 2],
              [-1, 2, 1]], dtype=float)

y = np.array([2, 9, -2], dtype=float)

A_perm, y_perm = permutation(A, y, 0, 2)
print(A_perm)
print(y_perm)

A_elim, y_elim = elimination(A, y, 0, 0)
print(A_elim)
print(y_elim)
