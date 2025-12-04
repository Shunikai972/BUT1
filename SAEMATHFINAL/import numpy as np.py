import numpy as np
import copy
import time
def test(mess,eval,res):
    print(mess,(eval==res)*'OK'+(eval!=res)*'Try again')
def test_determine_valuations(mess,list_var,res):
    test=mess+'Ok'
    list_testee=determine_valuations(list_var)
    for el in list_testee :
        if el not in res:
            test=mess+'Try again'
            return test
    for el in res:
        if el not in list_testee :
            test=mess+'Try again'
            return test
    for i in range(len(list_testee)-1):
        if list_testee[i] in list_testee[i+1:]:
            test=mess+'wowowow y a du doublon là-dedans'
            return test
    return test  

def test_for(mess,formu,res_for):
    res=True
    if (formu==[] and res_for!=[]) or (formu!=[] and res_for==[]):
        res=False
    for el1 in formu:
        for el2 in res_for:
            res=(set(el1)==set(el2))
            if res :
                break
        if not res :
            print(mess+'Try again !')
            return
    for el2 in res_for:
        for el1 in formu:
            res=(set(el2)==set(el1))
            if res :
                break
        if not res :
            print(mess+'Try again !')
            return
    res=False
    for i in range(len(formu)-1):
        for el in formu[i+1:]:
            if set(formu[i])==set(el):
                print(mess+'wowowow y a du doublon là-dedans')
                return 
    print(mess+'Ok')

#A VOUS DE JOUER#
def evaluer_clause(clause,list_var):
    if len(clause) == 0:
        return False
    none = False
    for i in clause : 
        val = list_var[abs(i)-1]
        if val == None: 
            none = True
            continue
        if i < 0:
            val = not val
        if val == True: 
            return True
        if none : 
            return None
        return False

def evaluer_cnf(formule,list_var):
    for clause in formule:
        if evaluer_clause(clause,list_var) == False:
            return False
        else: 
            if evaluer_clause(clause,list_var) == None:
                return None
    return True
def determine_valuation(list_var):
    if None not in list_var:
        return[list_var]
    true = list_var.copy()
    false = list_var.copy()
    for i in range(len(list_var)):
        if list_var[i] == None:
            true[i] == True
            false[i] = False
            break
    return determine_valuation(true) + determine_valuation(false)
def resol_sat_force_brute(formule,list_var):
    for val in determine_valuation(list_var):
        if evaluer_cnf(formule,val) == True:
            return(True,val)
    return False, []
def enlever_litt_for(formule, litteral):
    nouvelle_formule = []
    
    for clause in formule:
        clause_simplifie = []
        supprimer_clause = False

        for e in clause:
            if e == litteral:
                supprimer_clause = True
            elif e != -litteral:
                clause_simplifie.append(e)
        if not supprimer_clause:
            nouvelle_formule.append(clause_simplifie)
    return nouvelle_formule

def init_formule_simpl_for(formule_init,list_var):
    for i in range(len(list_var)):
        if list_var[i] is False:
            formule_init = enlever_litt_for(formule_init, -(i+1))
        if list_var[i] is True:
            formule_init = enlever_litt_for(formule_init, i+1)
    return formule_init
def retablir_for(formule_init,list_chgmts):
    for i, bool in list_chgmts:
        i = i + 1
    if bool is False: 
        i = -i
    formule_init = enlever_litt_for(formule_init, i)
    return formule_init
def progress (list_var, list_chgmts):
    for i in range(len(list_var)):
        if list_var[i] is None:
            list_var[i] = True
            list_chgmts.append(i, True)
            break
    return list_var, list_chgmts
def progress_simpl_for(formule, list_var, list_chgmts):
    list_var, list_chgmts = progress(list_var, list_chgmts)
    return retablir_for(formule, list_chgmts), list_var, list_chgmts
def retour(list_var, list_chgmts):
    if not list_chgmts:
        return list_var, list_chgmts
    i = len(list_chgmts) -1
    while i >= 0: 
        idx, val = list_chgmts[i]
        if val:
            list_chgmts[i] = (idx, False)
            list_var[idx] = False
            return list_var, list_chgmts
        list_chgmts.pop(i)
        list_var[idx] = None
        i -= 1
    return list_var, list_chgmts
def retour_simpl_for(formule_init,list_var,list_chgmts):
    list_var, list_chgmts = retour(list_var, list_chgmts)
    formule_init = retablir_for(formule_init,list_chgmts)
    return formule_init, list_var, list_chgmts
