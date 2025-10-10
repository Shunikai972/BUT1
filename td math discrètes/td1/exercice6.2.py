#6.2

def ajout(x, E):
    """ensemble avec élément x ajoutéééééééééééééé."""
    if x not in E:
        return E + [x]
    else:
        return E.copy()

def union(A, B):
    """union des ensembles a et bbbbbbbbbbbbbbbbbbbb"""
    result = A.copy()
    for x in B:
        if x not in result:
            result.append(x)
    return result

def intersection(A, B):
    """insterctions des ensembles a et bbbbbbbbbbbbbbbbbbbb"""
    return [x for x in A if x in B]

def retire(x, E):
    """nsemble avec élément x enlevééééééééé"""
    return [elem for elem in E if elem != x]

def diff(A, B):
    """ différence de A \ B."""
    return [x for x in A if x not in B]
#testeuh2
if __name__ == "__main__":
    A = [1, 2, 3]
    B = [2, 3, 4]
    print("ajout(4, A):", ajout(4, A))            
    print("union(A,B):", union(A, B))              
    print("intersection(A,B):", intersection(A, B)) 
    print("retire(2, A):", retire(2, A))          
    print("diff(A,B):", diff(A, B))                #