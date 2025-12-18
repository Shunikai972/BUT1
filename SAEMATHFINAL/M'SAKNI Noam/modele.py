
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





def creer_grille_init(list_grille,n):
    grille = [0] *n*n*n*n
    for ligne, colonne, val in list_grille:
        ligne = ligne -1
        colonne = colonne -1
        index = ligne *n*n + colonne
        grille[index] = val
    return grille
    '''Arguments : une liste de listes(contenant les coordonnées à renseigner et le nombre correspondant) et un entier donnant la taille de la grille
        Renvoie : une liste (list_grille_complete) avec les valeurs qui devront s'afficher dans la grille en la parcourant ligne après ligne de haut en bas et de gauche à droite
'''

def creer_grille_final(list_var,n):
    N = n*n
    grille = [0] * (N*N)
    for idx, val in enumerate(list_var):
        if val is True :
            k = (idx % N) + 1
            j = (idx // N) % N + 1
            i = (idx // (N*N)) + 1
            grille[(i-1)*N + (j-1)] = k
    return grille
    '''
    Renvoie : une liste (list_grille_complete) avec les valeurs qui devront s'afficher dans la grille (en fonction des valeurs logiques prises par les variables de list_var) en la parcourant ligne après ligne de haut en bas et de gauche à droite
'''

def for_conj_sudoku(n):
    N = n * n
    n4 = N * N

    def var(r, c, k):
        return r * n4 + c * N + k

    clauses = []
    seen = set()

    def add(cl):
        t = tuple(cl)
        if t not in seen:
            seen.add(t)
            clauses.append(cl)

    for k in range(1, N + 1):
        for start in range(n4):

            r1 = start // N
            c1 = start % N

            for k2 in range(k + 1, N + 1):
                add([-var(r1, c1, k), -var(r1, c1, k2)])

            for suite in range(start + 1, n4):
                r2 = suite // N
                c2 = suite % N
                same_blk = (r1 // n, c1 // n) == (r2 // n, c2 // n)

                if c1 == c2 or r1 == r2 or same_blk:
                    add([-var(r1, c1, k), -var(r2, c2, k)])

        for r in range(N):
            add([var(r, c, k) for c in range(N)])
        for c in range(N):
            add([var(r, c, k) for r in range(N)])
        for br in range(n):
            for bc in range(n):
                add([var(br*n + dr, bc*n + dc, k)
                     for dr in range(n)
                     for dc in range(n)])

    return clauses
    '''
    Renvoie : la formule (liste de listes) associée à une grille de sudoku de taille n selon les attentes formulées dans le sujet
    '''

def init_list_var(grille,n):
    N = n * n
    list_var = []
    list_None = [None] * N
    for i in range(len(grille)):
        if grille[i] == 0 :
            list_var = list_var + list_None
        else :
            list_false = [False] * N
            list_false[grille[i] - 1] = True
            list_var = list_var + list_false
    return list_var
    '''
    Renvoie : une liste list_var initialisant une valuation tenant compte des valeurs non nulles déjà renseignées dans list_grille_complete
'''


if __name__ == '__main__':
    #REALISER VOS TESTS ICI
    #TEST evaluer_clause
    clause1=[1,-2,3,-4]
    list_var1=[True,True,False,None]
    test("essai cas 1 evaluer_clause : ",evaluer_clause(clause1,list_var1),True)
    clause2=[1,-2,3,-4]
    list_var2=[False,True,False,None]
    test("essai cas 2 evaluer_clause : ",evaluer_clause(clause2,list_var2),None)
    clause3=[1,-2,3,-4]
    list_var3=[None,True,False,True]
    test("essai cas 3 evaluer_clause : ",evaluer_clause(clause3,list_var3),None)
    clause4=[1,-3]
    list_var4=[False,False,True]
    test("essai cas 4 evaluer_clause : ",evaluer_clause(clause4,list_var4),False)
    clause5=[]
    list_var5=[False,False,True]
    test("essai cas 5 evaluer_clause : ",evaluer_clause(clause5,list_var5),False)
    clause6=[1,2,3]
    list_var6=[False,False,True]
    test("essai cas 6 evaluer_clause : ",evaluer_clause(clause6,list_var6),True)
    clause7=[1,-2,3,4]
    list_var7=[None,True,False,True]
    test("essai cas 7 evaluer_clause : ",evaluer_clause(clause7,list_var7),True)


    #TEST evaluer_cnf
    for1=[[1,2],[2,-3,4],[-1,-2],[-1,-2,-3],[1]]
    list_var_for1_test1=[True,False,False,None]
    test('test1 evaluer_cnf : ',evaluer_cnf(for1,list_var_for1_test1),True)
    list_var_for1_test2=[None,False,False,None]
    test('test2 evaluer_cnf : ',evaluer_cnf(for1,list_var_for1_test2),None)
    list_var_for1_test3=[True,False,True,False]
    test('test3 evaluer_cnf : ',evaluer_cnf(for1,list_var_for1_test3),False)

    #TEST determine_valuations
    list_var1=[True,None,False,None]
    print(test_determine_valuations('res_test_determine_valuations cas 1 : ',list_var1,[[True, True, False, True], [True, False, False, True], [True, True, False, False], [True, False, False, False]]))
    list_var2=[None,False,True,None,True,False]
    print(test_determine_valuations('res_test_determine_valuations cas 2 : ',list_var2,[[True, False, True, True, True, False], [False, False, True, True, True, False], [True, False, True, False, True, False], [False, False, True, False, True, False]]))
    list_var3=[False,True,True,False]
    print(test_determine_valuations('res_test_determine_valuations cas 3 : ',list_var3,[[False, True, True, False]]))
    list_var4=[None,None,None]
    print(test_determine_valuations('res_test_determine_valuations cas 4 : ',list_var4,[[True, True, True], [False, True, True], [True, False, True], [False, False, True], [True, True, False], [False, True, False], [True, False, False], [False, False, False]]))


    #TEST resol_sat_force_brute
    for1=[[1,2],[2,-3,4],[-1,-2],[-1,-2,-3],[1],[-1,2,3]]
    list_var_for1=[None,None,None,None]
    test('test1 resol_sat_force_brute : ',resol_sat_force_brute(for1,list_var_for1),(True,[True, False, True, True]))
    for2=[[1,4,-5],[-1,-5],[2,-3,5],[2,-4],[2,4,5],[-1,-2],[-1,2,-3],[-2,4,-5],[1,-2]]
    list_var_for2=[None,None,None,None,None]
    test('test2 resol_sat_force_brute : ',resol_sat_force_brute(for2,list_var_for2),(False,[]))
    for3=[[-1,-2],[-1,2,-3,4],[2,3,4],[3],[1,-4],[-1,2],[1,2]]
    list_var_for3=[None,None,None,None]
    test('test3 resol_sat_force_brute : ',resol_sat_force_brute(for3,list_var_for3),(True,[False, True, True, False]))
    for4=[[-1,-2],[-1,2,-3,4],[2,3,4],[3],[1,-4],[-1,2],[1,2]]
    list_var_for4=[None,None,None,True]
    test('test4 resol_sat_force_brute : ',resol_sat_force_brute(for4,list_var_for4),(False,[]))


    #TEST enlever_litt_for
    for1=[[1,2,4,-5],[-1,2,3,-4],[-1,-2,-5],[-3,4,5],[-2,3,4,5],[-4]]
    litt1=4
    test('essai cas 1 enlever_litt_for : ',enlever_litt_for(for1,litt1),[[-1, 2, 3], [-1, -2, -5], []])

    #TEST init_formule_simpl_for
    list_var_for1=[False, None, None, False, None]
    for1=[[-5, -3, 4, -1], [3], [5, -2], [-2, 1, -4], [1, -3]]
    cor_for1=[[3], [5, -2], [-3]]
    test_for('test1_init_formule_simpl_for : ',init_formule_simpl_for(for1,list_var_for1),cor_for1)
    list_var_for2= [False, True, False, True, False]
    for2= [[3, 2, 1], [-1, -2, 5]]
    cor_for2=[]
    test_for('test2_init_formule_simpl_for : ',init_formule_simpl_for(for2,list_var_for2),cor_for2)
    list_var_for3= [None, None, None, True, None]
    for3= [[-5, -1], [-1, -3], [4], [-4, 1], [-2, -1, 3]]
    cor_for3=[[-5, -1], [-1, -3], [1], [-2, -1, 3]]
    test_for('test3_init_formule_simpl_for : ',init_formule_simpl_for(for3,list_var_for3),cor_for3)


    #TEST retablir_for
    formule_init=[[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    list_chgmts1=[(0, True), (1, True), (2, False)]
    form1=[[-5], [4, 5], [-4, 5]]
    list_chgmts2=[(0, True), (1, True), (2, False), (3, True), (4, False)]
    form2=[[]]
    list_chgmts3=[(0, True), (1, True), (2, False), (3, False)]
    form3=[[-5], [5]]
    test('essai cas 1 retablir_for : ',retablir_for(formule_init,list_chgmts1),form1)
    test('essai cas 2 retablir_for : ',retablir_for(formule_init,list_chgmts2),form2)
    test('essai cas 3 retablir_for : ',retablir_for(formule_init,list_chgmts3),form3)


    #TEST progress
    list_var=[True, None, None, None, None]
    list_chgmts=[(0, True)]
    l1=[True, True, None, None, None]
    l2=[(0, True), (1, True)]
    test("essai cas 1 progress : ",progress(list_var,list_chgmts),(l1,l2))
    list_var=[True, False, False, None, None]
    list_chgmts=[(0, True), (1, False), (2, False)]
    l1=[True, False, False, True, None]
    l2=[(0, True), (1, False), (2, False), (3, True)]
    test("essai cas 2 progress : ",progress(list_var,list_chgmts),(l1,l2))
    list_var=[None, None, None, None, None]
    list_chgmts=[]
    l1=[True, None, None, None, None]
    l2=[(0, True)]
    test("essai cas 3 progress : ",progress(list_var,list_chgmts),(l1,l2))
    list_var=[False, None, None, None, None]
    list_chgmts=[(0, False)]
    l1=[False, True, None, None, None]
    l2=[(0, False), (1, True)]
    test("essai cas 4 progress : ",progress(list_var,list_chgmts),(l1,l2))
    list_var=[True, False, None, None, None]
    list_chgmts=[]
    l1=[True, False, True, None, None]
    l2=[(2, True)]
    test("essai cas 5 progress : ",progress(list_var,list_chgmts),(l1,l2))
    list_var=[True, False, False, None, None]
    list_chgmts=[(2, False)]
    l1=[True, False, False, True, None]
    l2=[(2, False), (3, True)]
    test("essai cas 6 progress : ",progress(list_var,list_chgmts),(l1,l2))

    #TEST progress_simpl_for
    formule= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    list_var= [None, None, None, None, None]
    list_chgmts= []
    cor_form,cor_l1,cor_l2= ([[2, 3, -4], [-2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]],[True, None, None, None, None],[(0, True)])
    test('essai1_progress_simpl_for : ',progress_simpl_for(formule,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))
    formule= [[-5], [5]]
    list_var= [True, True, True, False, None]
    list_chgmts= [(0, True), (1, True), (2, True), (3, False)]
    cor_form,cor_l1,cor_l2= ([[]],[True, True, True, False, True],[(0, True), (1, True), (2, True), (3, False), (4, True)])
    test('essai2_progress_simpl_for : ',progress_simpl_for(formule,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))
    formule= [[3, -4], [-3, 4, 5], [-4, 5]]
    list_var= [True, False, None, None, None]
    list_chgmts= [(0, True), (1, False)]
    cor_form,cor_l1,cor_l2= ([[4, 5], [-4, 5]],[True, False, True, None, None],[(0, True), (1, False), (2, True)])
    test('essai3_progress_simpl_for : ',progress_simpl_for(formule,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))


    '''#TEST progress_simpl_for_dpll
    formule= [[-5], [4, 5], [-4, 5]]
    list_var= [True, True, False, None, None]
    list_chgmts= [(0, True), (1, True), (2, False)]
    list_sans_retour= []
    cor_for,cor_l1,cor_l2,cor_l3= ([[4], [-4]],[True, True, False, None, False],[(0, True), (1, True), (2, False), (4, False)],[4])
    test('essai1_progress_simpl_for_dpll : ',progress_simpl_for_dpll(formule,list_var,list_chgmts,list_sans_retour),(cor_for,cor_l1,cor_l2,cor_l3))
    formule= [[-5,4], [2,4, 5], [-2, 5]]
    list_var= [True, None, None, None, None]
    list_chgmts= [(0, True)]
    list_sans_retour= [0]
    cor_for,cor_l1,cor_l2,cor_l3= ([[-2,5]],[True, None, None, True, None],[(0, True),(3, True)],[0,3])
    test('essai2_progress_simpl_for_dpll : ',progress_simpl_for_dpll(formule,list_var,list_chgmts,list_sans_retour),(cor_for,cor_l1,cor_l2,cor_l3))
    formule=[[1,2,4,-5],[-1,2,3,-4],[-1,-2,-5],[-3,4,5],[-2,3,4,5],[-4,5]]
    list_var=[None,None,None,None,None]
    list_chgmts= []
    list_sans_retour= []
    cor_for,cor_l1,cor_l2,cor_l3=([[2, 3, -4], [-2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]], [True, None, None, None, None], [(0, True)], [])
    test('essai3_progress_simpl_for_dpll : ',progress_simpl_for_dpll(formule,list_var,list_chgmts,list_sans_retour),(cor_for,cor_l1,cor_l2,cor_l3))
    '''
    #TEST retour
    list_var= [True, True, None, None, None]
    list_chgmts= [(0, True), (1, True)]
    l1= [True, False, None, None, None]
    l2= [(0, True), (1, False)]
    test("essai cas 1 retour : ",retour(list_var,list_chgmts),(l1,l2))
    list_var= [True, False, None, None, None]
    list_chgmts= [(0, True), (1, False)]
    l1= [False, None, None, None, None]
    l2= [(0, False)]
    test("essai cas 2 retour : ",retour(list_var,list_chgmts),(l1,l2))
    list_var= [True, False, False, True, None]
    list_chgmts= []
    l1= [True, False, False, True, None]
    l2= []
    test("essai cas 3 retour : ",retour(list_var,list_chgmts),(l1,l2))
    list_var= [True, False, False, False, False]
    list_chgmts= [(0, True), (1, False), (2, False), (3, False), (4, False)]
    l1= [False, None, None, None, None]
    l2= [(0, False)]
    test("essai cas 4 retour : ",retour(list_var,list_chgmts),(l1,l2))
    list_var= [True, True, False, True, None]
    list_chgmts= [(1, True)]
    l1= [True, False, False, True, None]
    l2= [(1, False)]
    test("essai cas 5 retour : ",retour(list_var,list_chgmts),(l1,l2))
    list_var= [True, False, False, True, None]
    list_chgmts= [(1, False)]
    l1= [True, None, False, True, None]
    l2= []
    test("essai cas 6 retour : ",retour(list_var,list_chgmts),(l1,l2))


    #TEST retour_simpl_for
    formule_init= [[1, -5,], [-1], [-5], [1]]
    list_var= [True, True, False, False, True]
    list_chgmts= [(0, True), (4, True)]
    cor_form,cor_l1,cor_l2= ([[]],[True, True, False, False, False],[(0, True), (4, False)])
    test('essai1_retour_simpl_for : ',retour_simpl_for(formule_init,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))
    formule_init= [[3,4], [-3]]
    list_var= [False, True, True, None, None]
    list_chgmts= [(2, True)]
    cor_form,cor_l1,cor_l2= ([[4]],[False, True, False, None, None],[(2, False)])
    test('essai2_retour_simpl_for : ',retour_simpl_for(formule_init,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))
    formule_init= [[1, -5, -4], [4, -1,2], [-5, 4], [1, 4,-2], [-4, 5]]
    list_var= [True, False, False, False, None]
    list_chgmts= [(0, True),(1,False),(3, False)]
    cor_form,cor_l1,cor_l2= ([[-5, -4], [-5, 4], [4, -2], [-4, 5]], [False, None, False, None, None], [(0, False)])
    test('essai3_retour_simpl_for : ',retour_simpl_for(formule_init,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))
    formule_init= [[5, -3], [-5, 3]]
    list_var= [False, True, True, False, False]
    list_chgmts= [(2, True),(4,False)]
    cor_form,cor_l1,cor_l2= ([[-5]],[False, True, False, False, None],[(2, False)])
    test('essai4_retour_simpl_for : ',retour_simpl_for(formule_init,list_var,list_chgmts),(cor_form,cor_l1,cor_l2))


    '''#TEST retour_simpl_for_dpll
    formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    list_var= [True, True, False, True, False]
    list_chgmts= [(0, True), (1, True), (2, False), (4, False), (3, True)]
    list_sans_retour= [4, 3]
    cor_form,cor_l1,cor_l2,cor_l3= ([[3, -4], [-3, 4, 5], [-4, 5]], [True, False, None, None, None], [(0, True), (1, False)], [])
    test('essai1_retour_simpl_for_dpll : ',retour_simpl_for_dpll(formule_init,list_var,list_chgmts,list_sans_retour),(cor_form,cor_l1,cor_l2,cor_l3))
    formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    list_var= [True, True, True, True, False]
    list_chgmts= [(0, True), (1, True), (2, True), (3, True), (4, False)]
    list_sans_retour= []
    cor_form,cor_l1,cor_l2,cor_l3= ([[-5], [5]], [True, True, True, False, None], [(0, True), (1, True), (2, True), (3, False)], [])
    test('essai2_retour_simpl_for_dpll : ',retour_simpl_for_dpll(formule_init,list_var,list_chgmts,list_sans_retour),(cor_form,cor_l1,cor_l2,cor_l3))
    formule_init= [[1], [-2], [-1], [-4,-2]]
    list_var= [True, None, False, None, True]
    list_chgmts= [(0, True)]
    list_sans_retour= [0]
    cor_form,cor_l1,cor_l2,cor_l3= ([[1], [-2], [-1], [-4, -2]], [None, None, False, None, True], [], [])
    test('essai3_retour_simpl_for_dpll : ',retour_simpl_for_dpll(formule_init,list_var,list_chgmts,list_sans_retour),(cor_form,cor_l1,cor_l2,cor_l3))
    '''

    #TEST resol_parcours_arbre
    formule_init= [[1, 4, -5], [-1, -5], [2, -3, 5], [2, -4], [2, 4, 5], [-1, -2], [-1, 2, -3], [-2, 4, -5], [1, -2]]
    list_var= [True, True, False, True, None]
    list_chgmts= [(1, True)]
    cor_resol=(False, [])
    test('essai1_resol_parcours_arbre : ',resol_parcours_arbre(formule_init,list_var,list_chgmts),cor_resol)
    formule_init= [[5], [3, -5, -1, -2], [1, -2, -5], [2, -5, 1, -3], [3]]
    list_var= [True, False, None, False, True]
    list_chgmts= [(0, True)]
    cor_resol=(True,[True, False, True, False, True])
    test('essai2_resol_parcours_arbre : ',resol_parcours_arbre(formule_init,list_var,list_chgmts),cor_resol)
    formule_init= [[-5, 2, -3, -4], [1, -5], [5, 2], [3, -2, 4], [5, -2, -1]]
    list_var= [False, True, False, None, None]
    list_chgmts= [(1, True)]
    cor_resol=(True,[False, True, False, True, False])
    test('essai3_resol_parcours_arbre : ',resol_parcours_arbre(formule_init,list_var,list_chgmts),cor_resol)
    formule_init= [[1, 4, -5], [-1, -5], [2, -3, 5], [2, -4], [2, 4, -5], [1, 2], [-1, 2, -3], [-2, 4, -5], [1, -2]]
    list_var= [False, True, False, False, None]
    list_chgmts= [(0,False),(1, True)]
    cor_resol=(False, [])
    test('essai4_resol_parcours_arbre : ',resol_parcours_arbre(formule_init,list_var,list_chgmts),cor_resol)
    formule_init= [[1, 4, -5], [-1, -5], [2, -3, 5], [2, -4], [2, 4, -5], [-1, -2], [-1, 2, -3], [-2, 4, -5], [1, -2]]
    list_var= [False, True, False, False, None]
    list_chgmts= [(0,False),(1, True)]
    cor_resol=(True, [False, False, False, False, False])
    test('essai5_resol_parcours_arbre : ',resol_parcours_arbre(formule_init,list_var,list_chgmts),cor_resol)


    #TEST resol_parcours_arbre_simpl_for
    formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    formule= [[2, 3, -4], [-2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    list_var= [True, None, None, None, None]
    list_chgmts= [(0, True)]
    cor_resol=(True, [True, False, True, True, True])
    test('essai1_resol_parcours_arbre_simpl_for : ',resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts),cor_resol)
    formule_init= [[5], [3, -5, -1], [1, -5], [3]]
    formule= [[5], [-5]]
    list_var= [False, True, True, False, None]
    list_chgmts= [(0,False),(2, True),(3,False)]
    cor_resol=(False, [])
    test('essai2_resol_parcours_arbre_simpl_for : ',resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts),cor_resol)
    formule_init= [[-5, 2, -3, -4], [-5], [5, 2], [3, -2, 4]]
    formule= [[-5], [4]]
    list_var= [False, True, False, None, None]
    list_chgmts= [(1, True),(2,False)]
    cor_resol=(True, [False, True, False, True, False])
    test('essai3_resol_parcours_arbre_simpl_for : ',resol_parcours_arbre_simpl_for(formule_init,formule,list_var,list_chgmts),cor_resol)


    '''#TEST resol_parcours_arbre_simpl_for_dpll
    formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    formule= [[2, 3, -4], [-2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    list_var= [True, None, None, None, None]
    list_chgmts= [(0, True)]
    list_sans_retour= []
    cor_resol=(True, [True, False, True, None, True])
    test('essai1_resol_parcours_arbre_simpl_for_dpll : ',resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour),cor_resol)
    formule_init= [[1, 2, 4, -5], [-1, 2, 3, -4], [-1, -2, -5], [-3, 4, 5], [-2, 3, 4, 5], [-4, 5]]
    formule= [[3, -4]]
    list_var= [True, False, None, None, True]
    list_chgmts= [(0, True), (1, False), (4, True)]
    list_sans_retour= [4]
    cor_resol=(True, [True, False, True, None, True])
    test('essai2_resol_parcours_arbre_simpl_for_dpll : ',resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour),cor_resol)
    formule_init= [[-5, 2, -3, -4], [1, -5], [5, 2], [3, -2, 4], [5, -2, -1]]
    formule= [[2], [-2, 4]]
    list_var= [False, None, False, None, False]
    list_chgmts= [(4, False)]
    list_sans_retour= [4]
    cor_resol=(True, [False, True, False, True, False])
    test('essai3_resol_parcours_arbre_simpl_for_dpll : ',resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour),cor_resol)
    formule_init= [[5], [3, -5, -1, -2], [1, 2, -5], [2, -5, -1, -3], [3],[-3,-2]]
    formule= [[3, -1, -2], [1, 2], [2, -1, -3], [3],[-3,-2]]
    list_var= [None, None, None, False, True]
    list_chgmts= [(4, True)]
    list_sans_retour=[4]
    cor_resol=(False, [])
    test('essai4_resol_parcours_arbre_simpl_for_dpll : ',resol_parcours_arbre_simpl_for_dpll(formule_init,formule,list_var,list_chgmts,list_sans_retour),cor_resol)
    '''

    #TEST creer_grille_init
    list_grille3=[[1,3,2],[1,6,5],[2,5,4],[2,8,9],[2,9,3],[3,2,7],[3,9,6],[4,3,1],[4,4,8],[4,8,3],[5,1,7],[5,2,2],[5,5,6],[5,8,8],[5,9,4],[6,2,4],[6,6,2],[6,7,5],[7,1,3],[7,8,1],[8,1,4],[8,2,6],[8,5,7],[9,4,9],[9,7,8]]
    cor_grille3=[0, 0, 2, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 4, 0, 0, 9, 3, 0, 7, 0, 0, 0, 0, 0, 0, 6, 0, 0, 1, 8, 0, 0, 0, 3, 0, 7, 2, 0, 0, 6, 0, 0, 8, 4, 0, 4, 0, 0, 0, 2, 5, 0, 0, 3, 0, 0, 0, 0, 0, 0, 1, 0, 4, 6, 0, 0, 7, 0, 0, 0, 0, 0, 0, 0, 9, 0, 0, 8, 0, 0]
    test("essai creer_grille_init : ",creer_grille_init(list_grille3,3),cor_grille3)


    #TEST creer_grille_final
    list_var_fin=[False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, True, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, True, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, True, False, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, False, False, False, False, False, False, True, False, False, False, False, True, False, False, False, False, False, False, False, False, False, True, False, False, False, False]
    cor_grille_final=[9, 3, 2, 6, 1, 5, 4, 7, 8, 5, 8, 6, 2, 4, 7, 1, 9, 3, 1, 7, 4, 3, 8, 9, 2, 5, 6, 6, 9, 1, 8, 5, 4, 7, 3, 2, 7, 2, 5, 1, 6, 3, 9, 8, 4, 8, 4, 3, 7, 9, 2, 5, 6, 1, 3, 5, 9, 4, 2, 8, 6, 1, 7, 4, 6, 8, 5, 7, 1, 3, 2, 9, 2, 1, 7, 9, 3, 6, 8, 4, 5]
    test("essai creer_grille_final : ",creer_grille_final(list_var_fin,3),cor_grille_final)


    #TEST for_conj_sudoku
    corrige_for2=[[-1, -21], [-1, -5], [-1, -17], [-1, -9], [-1, -33], [-1, -13], [-1, -49], [-1, -2], [-1, -3], [-1, -4], [-5, -17], [-5, -9], [-17, -33], [-5, -13], [-17, -49], [-5, -6], [-5, -7], [-5, -8], [-9, -13], [-33, -49], [-9, -10], [-9, -11], [-9, -12], [-13, -14], [-13, -15], [-13, -16], [1, 5, 9, 13], [1, 17, 33, 49], [1, 5, 17, 21], [-9, -29], [-17, -21], [-5, -21], [-17, -25], [-5, -37], [-17, -29], [-5, -53], [-17, -18], [-17, -19], [-17, -20], [-13, -25], [-21, -25], [-21, -37], [-21, -29], [-21, -53], [-21, -22], [-21, -23], [-21, -24], [-25, -29], [-37, -53], [-25, -26], [-25, -27], [-25, -28], [-29, -30], [-29, -31], [-29, -32], [17, 21, 25, 29], [5, 21, 37, 53], [9, 13, 25, 29], [-33, -53], [-33, -37], [-9, -25], [-33, -41], [-9, -41], [-33, -45], [-9, -57], [-33, -34], [-33, -35], [-33, -36], [-37, -49], [-37, -41], [-25, -41], [-37, -45], [-25, -57], [-37, -38], [-37, -39], [-37, -40], [-41, -45], [-41, -57], [-41, -42], [-41, -43], [-41, -44], [-45, -46], [-45, -47], [-45, -48], [33, 37, 41, 45], [9, 25, 41, 57], [33, 37, 49, 53], [-41, -61], [-49, -53], [-13, -29], [-49, -57], [-13, -45], [-49, -61], [-13, -61], [-49, -50], [-49, -51], [-49, -52], [-45, -57], [-53, -57], [-29, -45], [-53, -61], [-29, -61], [-53, -54], [-53, -55], [-53, -56], [-57, -61], [-45, -61], [-57, -58], [-57, -59], [-57, -60], [-61, -62], [-61, -63], [-61, -64], [49, 53, 57, 61], [13, 29, 45, 61], [41, 45, 57, 61], [-2, -22], [-2, -6], [-2, -18], [-2, -10], [-2, -34], [-2, -14], [-2, -50], [-2, -3], [-2, -4], [-6, -18], [-6, -10], [-18, -34], [-6, -14], [-18, -50], [-6, -7], [-6, -8], [-10, -14], [-34, -50], [-10, -11], [-10, -12], [-14, -15], [-14, -16], [2, 6, 10, 14], [2, 18, 34, 50], [2, 6, 18, 22], [-10, -30], [-18, -22], [-6, -22], [-18, -26], [-6, -38], [-18, -30], [-6, -54], [-18, -19], [-18, -20], [-14, -26], [-22, -26], [-22, -38], [-22, -30], [-22, -54], [-22, -23], [-22, -24], [-26, -30], [-38, -54], [-26, -27], [-26, -28], [-30, -31], [-30, -32], [18, 22, 26, 30], [6, 22, 38, 54], [10, 14, 26, 30], [-34, -54], [-34, -38], [-10, -26], [-34, -42], [-10, -42], [-34, -46], [-10, -58], [-34, -35], [-34, -36], [-38, -50], [-38, -42], [-26, -42], [-38, -46], [-26, -58], [-38, -39], [-38, -40], [-42, -46], [-42, -58], [-42, -43], [-42, -44], [-46, -47], [-46, -48], [34, 38, 42, 46], [10, 26, 42, 58], [34, 38, 50, 54], [-42, -62], [-50, -54], [-14, -30], [-50, -58], [-14, -46], [-50, -62], [-14, -62], [-50, -51], [-50, -52], [-46, -58], [-54, -58], [-30, -46], [-54, -62], [-30, -62], [-54, -55], [-54, -56], [-58, -62], [-46, -62], [-58, -59], [-58, -60], [-62, -63], [-62, -64], [50, 54, 58, 62], [14, 30, 46, 62], [42, 46, 58, 62], [-3, -23], [-3, -7], [-3, -19], [-3, -11], [-3, -35], [-3, -15], [-3, -51], [-3, -4], [-7, -19], [-7, -11], [-19, -35], [-7, -15], [-19, -51], [-7, -8], [-11, -15], [-35, -51], [-11, -12], [-15, -16], [3, 7, 11, 15], [3, 19, 35, 51], [3, 7, 19, 23], [-11, -31], [-19, -23], [-7, -23], [-19, -27], [-7, -39], [-19, -31], [-7, -55], [-19, -20], [-15, -27], [-23, -27], [-23, -39], [-23, -31], [-23, -55], [-23, -24], [-27, -31], [-39, -55], [-27, -28], [-31, -32], [19, 23, 27, 31], [7, 23, 39, 55], [11, 15, 27, 31], [-35, -55], [-35, -39], [-11, -27], [-35, -43], [-11, -43], [-35, -47], [-11, -59], [-35, -36], [-39, -51], [-39, -43], [-27, -43], [-39, -47], [-27, -59], [-39, -40], [-43, -47], [-43, -59], [-43, -44], [-47, -48], [35, 39, 43, 47], [11, 27, 43, 59], [35, 39, 51, 55], [-43, -63], [-51, -55], [-15, -31], [-51, -59], [-15, -47], [-51, -63], [-15, -63], [-51, -52], [-47, -59], [-55, -59], [-31, -47], [-55, -63], [-31, -63], [-55, -56], [-59, -63], [-47, -63], [-59, -60], [-63, -64], [51, 55, 59, 63], [15, 31, 47, 63], [43, 47, 59, 63], [-4, -24], [-4, -8], [-4, -20], [-4, -12], [-4, -36], [-4, -16], [-4, -52], [-8, -20], [-8, -12], [-20, -36], [-8, -16], [-20, -52], [-12, -16], [-36, -52], [4, 8, 12, 16], [4, 20, 36, 52], [4, 8, 20, 24], [-12, -32], [-20, -24], [-8, -24], [-20, -28], [-8, -40], [-20, -32], [-8, -56], [-16, -28], [-24, -28], [-24, -40], [-24, -32], [-24, -56], [-28, -32], [-40, -56], [20, 24, 28, 32], [8, 24, 40, 56], [12, 16, 28, 32], [-36, -56], [-36, -40], [-12, -28], [-36, -44], [-12, -44], [-36, -48], [-12, -60], [-40, -52], [-40, -44], [-28, -44], [-40, -48], [-28, -60], [-44, -48], [-44, -60], [36, 40, 44, 48], [12, 28, 44, 60], [36, 40, 52, 56], [-44, -64], [-52, -56], [-16, -32], [-52, -60], [-16, -48], [-52, -64], [-16, -64], [-48, -60], [-56, -60], [-32, -48], [-56, -64], [-32, -64], [-60, -64], [-48, -64], [52, 56, 60, 64], [16, 32, 48, 64], [44, 48, 60, 64]]
    test_for('test_for_conj_sudoku : ',for_conj_sudoku(2),corrige_for2)


    #TEST init_list_var
    grille2= [0, 1, 0, 0, 4, 2, 0, 0, 0, 0, 2, 0, 0, 3, 0, 0]
    cor_list_var_grille2= [None, None, None, None, True, False, False, False, None, None, None, None, None, None, None, None, False, False, False, True, False, True, False, False, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, None, False, True, False, False, None, None, None, None, None, None, None, None, False, False, True, False, None, None, None, None, None, None, None, None]
    test('test_init_list_var : ',init_list_var(grille2,2),cor_list_var_grille2)











