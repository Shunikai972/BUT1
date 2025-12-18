def retour_simpl_for(formule_init,list_var,list_chgmts):
    list_var, list_chgmts = retour(list_var, list_chgmts)
    return retablir_for(formule_init, list_chgmts), list_var, list_chgmts
    '''
Renvoie : form,l1,l2
    form : nouvelle formule
    l1 : nouvelle list_var
    l2 : nouvelle list_chgmts
'''


def resol_parcours_arbre(formule_init,list_var,list_chgmts):

    bool_var, list_var2 = resol_sat_force_brute(formule_init, list_var)
    if bool_var:
        return True, list_var2
    nouv_list_var, nouv_list_chgmts = progress(list_var, list_chgmts)
    if nouv_list_chgmts == list_chgmts:
        ret_list_var, ret_list_chgmts = retour(list_var, list_chgmts)
        if ret_list_chgmts == []:
            return False, []

        return resol_parcours_arbre(formule_init, ret_list_var, ret_list_chgmts)
    return resol_parcours_arbre(formule_init, nouv_list_var, nouv_list_chgmts)

def resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts):#la même distinction peut être faite entre formule et formule_init
    '''
    Renvoie SAT,l1 avec :
SAT=True ou False
l1=une liste de valuations rendant la formule vraie ou une liste vide
'''
# J'ai utiliser l'Initialisation du parcours mais de manière différente j'ai donc enlever votre partie de code
    formule, list_var, list_chgmts = progress_simpl_for(formule_init, list_var, list_chgmts)
    if [] in formule :
        formule2, list_var2, list_chgmts2 = retour_simpl_for(formule_init, list_var, list_chgmts)
        if list_chgmts2 == [] :
            return False, []
        else :
            return resol_parcours_arbre_simpl_for(formule_init, formule2, list_var2, list_chgmts2)
    if formule == [] :
        return True, list_var
    return resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts)
def creer_grille_init(list_grille, n):
    grille = [0] *n*n*n*n
    for ligne, colonne, val in list_grille:
        ligne = ligne -1
        colonne = colonne -1
        index = ligne *n*n+ colonne
        grille[index] = val
    return grille
def creer_grille_final(list_var,n):
    N = n*n
    grille = [0]*N*N
    for idx, val in enumerate(list_var):
        if val is True:
            k = (idx%N) +1
            j =(idx//N)%N+1
            i = (idx//(N*N))+1
            grille[(i-1)*N+(j-1)] = k
    return grille
def for_conj_sudoku(n):
    N = n*n
    n4 = N*N
    
    def var(r,c,k):
        return r*n4+c*N+k
    clauses = []
    seen = set()
    def add(cl):
        t = tuples(cl)
        if t not in seen:
            seen.add(t)
            clauses.append(cl)
    for k in range (1,N+1):
        for start in range(n4):
            r1 = start // n
            c1 = start % n
            for k2 in range (k+1, N+1):
                add([-var(r1,c1,k), -var(r1,c1,k2)])
            for suite in range(start + 1, n4):
                r2 = suite // Nc2 
                c2 = suite % n
                same_blk =(r1//n,c1//n) == (r2//n,c2//n)
                if c1 == c2 or r1 == R2 or same_blk:
                    add([-var(r1,c1,k),-var(r2,c2,k)])
    for r in range(N):
        add  