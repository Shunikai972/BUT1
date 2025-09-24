x = int(input("valx "))
y = int(input("valy "))
def et(x,y):
    if(x== 1):
        if(y==1):
            return(1)
        else:
            return(0)
    elif(y == 1):
        if(x==1):
            return(1)
        else:
            return(0)
    else:
        return(0)
resultat = et(x, y)
def ou(x,y):
    if(x== 1):
        return(1)
    elif(y == 1):
        if(x==1):
            return(1)
    elif(x==0):
        if(y == 0):
            return(0)
def non(x):
    if(x== 1):
        return(0)
    elif(x == 0):
        return(1)
def implique(x, y):
    if x == 1:
        if y == 0:
            return 0
        else:
            return 1
    else:
        return 1

resultat = et(x, y)
resultat2 = ou(x, y)
resultat3 = non(x)
resultat4 = implique(x, y)
if resultat == 0:
    print(" RESULTAT ET: FALSE")
else:
    print("RESULTAT ET: TRUE")
if resultat2 == 0:
    print(" RESULTAT OU: FALSE")
else:
    print("RESULTAT OU: TRUE")
if resultat3 == 0:
    print(" RESULTAT NON: FALSE")
else:
    print("RESULTAT NON: TRUE")
if resultat4 == 0:
    print(" RESULTAT IMPLIQUE: FALSE")
else:
    print("RESULTAT IMPLIQUE: TRUE")








