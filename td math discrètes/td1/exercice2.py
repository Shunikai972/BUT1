import math
def g(x):
    a = 0
    if x < 0:
        a = -math.sqrt(x)
        print("cas négatif", a)
    else :
        a = math.sqrt(x)
        print("cas positif", a)
print("début")
x = int(input("quelle est valeure"))    
g(x)