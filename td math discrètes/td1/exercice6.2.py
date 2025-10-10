# 6.2 Opérations ensemblistes

def ajout(x, E):
    """Retourne l'ensemble E avec l'élément x ajouté (sans doublon)."""
    if x not in E:
        return E + [x]
    else:
        return E.copy()

def union(A, B):
    """Retourne l'union des ensembles A et B."""
    result = A.copy()
    for x in B:
        if x not in result:
            result.append(x)
    return result

def intersection(A, B):
    """Retourne l'intersection des ensembles A et B."""
    return [x for x in A if x in B]

def retire(x, E):
    """Retourne l'ensemble E sans l'élément x."""
    return [elem for elem in E if elem != x]

def diff(A, B):
    """Retourne la différence A \ B."""
    return [x for x in A if x not in B]

if __name__ == "__main__":
    A = [1, 2, 3]
    B = [2, 3, 4]
    print("ajout(4, A):", ajout(4, A))            # [1, 2, 3, 4]
    print("union(A,B):", union(A, B))              # [1, 2, 3, 4]
    print("intersection(A,B):", intersection(A, B)) # [2, 3]
    print("retire(2, A):", retire(2, A))          # [1, 3]
    print("diff(A,B):", diff(A, B))                # [1]