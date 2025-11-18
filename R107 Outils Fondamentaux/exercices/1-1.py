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