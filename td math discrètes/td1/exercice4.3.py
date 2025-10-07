import math
def value(nb):
    for nombre in range(2**nb) :
        binaire = bin(nombre)[2::].zfill(nb)
        print(binaire)
value(32 )

w