def inclus(A, B):
    """renvoie true si a inclu b"""
    for x in A:
        if x not in B:
            return False
    return True

def egal(A, B):
    """renvoie true si a égal b."""
    return inclus(A, B) and inclus(B, A)

def disjoint(A, B):
    """renvoie true si a et b sont disjoint."""
    for x in A:
        if x in B:
            return False
    return True



# testeuh
if __name__ == "__main__":
    A = [1, 2, 3]
    B = [2, 3, 4]

    print("inclus(A,B):", inclus(A, B))           # F
    print("inclus([2,3], B):", inclus([2, 3], B)) # true
    print("egal(A,B):", egal(A, B))               # F
    print("disjoint(A, B):", disjoint(A, B))      # F
