import numpy as np

"""
fonction auxiliaire de 'affichage'
"""
def get_number_rep(x) :
    if x == int(x) :
        return str(int(x))
    else :
        return "%.2f" % x
"""
fonction auxiliaire de 'affichage'
"""
def get_str_rep(x,max_size,i,first,print0) :
    str_x = get_number_rep(abs(x))
    n = len(str_x)
    if x == -1. :
        if i == first :
            return " "*(max_size)+"-x_"+str(i+1)
        else :
            return "-"+" "*(max_size)+"x_"+str(i+1)
    elif x == 1. :
        if i == first :
            return " "*(1+max_size)+"x_"+str(i+1)
        else :
            return "+"+" "*(max_size)+"x_"+str(i+1)
    elif x == 0.:
        if print0 :
            return " "*(max_size+3)+"0"
        else :
            return " "*(max_size+1+3)
    elif x < 0 :
        if i == first :
            return " "*(max_size-n)+"-"+str_x+"x_"+str(i+1)
        else :
            return "-"+" "*(max_size-n)+str_x+"x_"+str(i+1)
    elif i == first :
        return " "*(1+max_size-n)+str_x+"x_"+str(i+1)
    else :
        return "+"+" "*(max_size-n)+str_x+"x_"+str(i+1)

"""
fonction permettant l'affichage des systèmes linéaires
"""
def affichage(A,y) :
    res = ""
    max_size = 0
    first = [-1]*len(A)
    for i in range(len(A)) :
        for j in range(len(A[i])) :
            max_size = max(max_size,len(get_number_rep(A[i][j])))
            if first[i] ==-1 and A[i][j] != 0 :
                first[i] = j
    for i in range(len(A)) :
        res+="|"
        for j in range(len(A[i])) :
            res+=get_str_rep(A[i][j],max_size,j,first[i],first[i]==-1 and j == len(A[i])-1)
        res=res+" = "+(get_number_rep(y[i]))+"\n"
    print(res)

"""
fonction retournant le système représenté par A et y dans lequel les lignes i et k ont été permutées.
"""

def permutation(A, y, i, k):
    A2 = A.copy()  
    y2 = y.copy()

    A2[i], A2[k] = A2[k].copy(), A2[i].copy()
    y2[i], y2[k] = y2[k], y2[i]

    return A2, y2
A = np.array([[1., 2., 3.],
              [4., 5., 6.],
              [7., 8., 9.]])
y = np.array([10., 11., 12.])

# Permuter la première et la dernière ligne
A_perm, y_perm = permutation(A, y, 0, 2)

print("=== Test permutation ===")
print("Matrice originale:\n", A)
print("Vecteur original:", y)
print("Matrice permutée:\n", A_perm)
print("Vecteur permuté:", y_perm)

# Vérifications
assert np.array_equal(A_perm[0], A[2])
assert np.array_equal(A_perm[2], A[0])
assert np.array_equal(y_perm[0], y[2])
assert np.array_equal(y_perm[2], y[0])
print("Permutation OK ✅\n")
"""
fonction retournant le système représenté par A et y dans lequel la variable n°j a été supprimée des lignes
n°i+1 à n en utilisant la ligne n°i
"""
def elimination(A, y, i, j):
    A2 = A.copy()
    y2 = y.copy()

    pivot = A2[i, j]

    for k in range(i+1, A2.shape[0]):
        facteur = A2[k, j] / pivot
        A2[k, :] = A2[k, :] - facteur * A2[i, :]
        y2[k] = y2[k] - facteur * y2[i]

    return A2, y2
# Exemple simple
A = np.array([[2., 1., -1.],
              [-3., -1., 2.],
              [-2., 1., 2.]], dtype=float)
y = np.array([8., -11., -3.], dtype=float)

# Éliminer la variable 0 (colonne 0) en utilisant la première ligne (i=0, j=0)
A_elim, y_elim = elimination(A, y, 0, 0)

print("=== Test elimination ===")
print("Matrice originale:\n", A)
print("Vecteur original:", y)
print("Matrice après élimination:\n", A_elim)
print("Vecteur après élimination:", y_elim)

# Vérifications
# Les éléments de la première colonne en dessous de la ligne 0 doivent être 0
assert np.allclose(A_elim[1:, 0], 0)
print("Élimination OK ✅\n")

def get_next_pivot(A,i,j) :
    # à compléter
    return k,l

def Gauss(A,y) :
    # à compléter
    return A,y

def solveTriSup(A,y) :
    return 0

A=np.array([[1.,1.,1.],[1.,-1.,2.],[-1.,2.,1.]])
y=np.array([2.,9.,-2.])

affichage(A,y)
A,y = Gauss(A,y)

# systèmes demandés dans les exercices de TP
As=[]
ys=[]
As.append(np.zeros((2,2)))
ys.append(np.zeros(3))
As.append(np.array([[-1.,-3.,0],[1.,2.,3.],[2.,-2.,1.]]))
ys.append(np.array([3.,4.,4.]))
As.append(np.array([[-2.,10.,3.],[-5.,34.,12.],[-1.,8.,3.]]))
ys.append(np.array([-18.,-54.,-12.]))
As.append(np.array([[-3.,-1.,4.,3.],[0.,0.,4.,2.],[-2.,-2.,-2.,1.]]))
ys.append(np.array([8.,10.,-1.]))

As.append(np.array([[2.,15.,3.],[0.,21.,7.],[4.,9.,-1.]]))
ys.append(np.array([25.,35.,10.]))
As.append(np.array([[4.,3.,4.],[0.,0.,-1.],[3.,-3.,1.],[0,-1.,3.]]))
ys.append(np.array([1.,0.,6.,1.]))
As.append(np.array([[0.,0.,-2.],[-2.,2.,-2.],[-4.,4.,-4.],[-2.,2.,-4.]]))
ys.append(np.array([3.,-4.,-8.,-1.]))

for i in range(len(As)) :
    print("------- système n°"+str(i+1)+" -------------")
    affichage(As[i],ys[i])
    As[i],ys[i] = Gauss(As[i],ys[i])
    print(solveTriSup(As[i],ys[i]))
