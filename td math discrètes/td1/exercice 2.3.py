import math
def g(x):
    a = 0
    if x < 0:
        print("nombre invalide : négatif")
        return
    else :
        for i in range (0,x + 1):
            if  i % 7 == 0:
                print(i)
print("début")
x = int(input("quelle est valeure"))    
g(x)