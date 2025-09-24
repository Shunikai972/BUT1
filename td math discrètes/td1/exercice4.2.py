

def et(x, y):
    if x == 1:
        if y == 1:
            return 1
        else:
            return 0
    else:
        return 0

def ou(x, y):
    if x == 1:
        return 1
    else:
        if y == 1:
            return 1
        else:
            return 0

def non(x):
    if x == 1:
        return 0
    else:
        return 1

def implique(x, y):
    
    if x == 1:
        if y == 1:
            return 1
        else:
            return 0
    else:
        return 1

def main():
    
    p = [0,0,0,0,1,1,1,1]
    q = [0,0,1,1,0,0,1,1]
    r = [0,1,0,1,0,1,0,1]

    res = []  

    print(" p q r | F1 F2 | F1∧F2")
    print("-----------------------")
    for i in range(len(p)):
        pi = p[i]
        qi = q[i]
        ri = r[i]

        
        not_p = non(pi)
        q_and_r = et(qi, ri)
        F1 = implique(not_p, q_and_r)

        
        not_r = non(ri)
        notp_and_notr = et(not_p, not_r)
        F2 = implique(qi, notp_and_notr)

        conj = et(F1, F2)
        res.append(conj)

        print(f" {pi} {qi} {ri} |  {F1}  {F2} |   {conj}")

    
    satisfiable = False
    
    for v in res:
        if v == 1:
            satisfiable = True
            break

    print("-----------------------")
    if satisfiable:
        print("La formule F1 ∧ F2 est SATISFIABLE (il existe au moins une valuation qui la rend vraie).")
    else:
        print("La formule F1 ∧ F2 est INSATISFIABLE (aucune valuation ne la rend vraie).")

if __name__ == "__main__":
    main()
