import numpy as np
import matplotlib.pyplot as plt
import math as mt

def BaseLagrange(x, listX, i):
    produit = 1
    for idx in range(len(listX)):
        if idx != i:
            produit = produit * (x - listX[idx]) / (listX[i] - listX[idx])
    return produit

def InterLagrange(x, listX, listY):
    somme = 0
    for idx in range(len(listX)):
        somme += listY[idx] * BaseLagrange(x, listX, idx)
    return somme

def dicho(a, b, f, e):
    liste_valeurs = []
    while True:
        m = (a + b) / 2
        liste_valeurs.append(m)
        if abs(f(m)) < e:
            break
        if f(a) * f(m) < 0:
            b = m
        else:
            a = m
    return liste_valeurs

def fausse_pos(a, b, f, e):
    approx = []
    while True:
        if abs(f(b) - f(a)) < 1e-14:
            break
        m = (a * f(b) - b * f(a)) / (f(b) - f(a))
        approx.append(m)
        if abs(f(m)) < e:
            break
        if f(a) * f(m) < 0:
            b = m
        else:
            a = m
    return approx

def newton(a, f, df, e):
    liste_valeurs = []
    x = a
    while True:
        x = x - f(x) / df(x)
        liste_valeurs.append(x)
        if abs(f(x)) < e:
            break
    return liste_valeurs


### QUELQUES TESTS INDICATIFS
print('test BaseLagrange : ', np.isclose(BaseLagrange(0.2, np.linspace(-1, 1, 10), 6), 0.37638881280000036))
print('test2 BaseLagrange : ', np.isclose(BaseLagrange(0.2, np.linspace(-1, 1, 11), 6), 1))
print('test3 BaseLagrange : ', np.isclose(BaseLagrange(0.2, np.linspace(-1, 1, 11), 4), 0))
print('test1 InterLagrange : ', np.isclose(InterLagrange(np.pi/3, np.linspace(0, np.pi, 7), np.cos(np.linspace(0, np.pi, 7))), 0.5))
print('test2 InterLagrange : ', np.isclose(InterLagrange(np.pi/3, np.linspace(0, np.pi, 5), np.cos(np.linspace(0, np.pi, 5))), 0.4969732592091243))

list_res_dicho1 = [1.5, 1.25, 1.375, 1.4375, 1.40625, 1.421875, 1.4140625, 1.41796875, 1.416015625, 1.4150390625, 1.41455078125, 1.414306640625, 1.4141845703125, 1.41424560546875, 1.414215087890625]
print('test dicho : ', list(np.isclose(np.array(dicho(1, 2, lambda x: x**2 - 2, 10**(-5))), np.array(list_res_dicho1))).count(True) == len(list_res_dicho1))

list_res_fausse_pos1 = [1.3333333333333333, 1.4, 1.411764705882353, 1.4137931034482758, 1.414141414141414, 1.4142011834319526, 1.41421143847487]
print('test fausse_pos1 : ', list(np.isclose(np.array(fausse_pos(1, 2, lambda x: x**2 - 2, 10**(-5))), np.array(list_res_fausse_pos1))).count(True) == len(list_res_fausse_pos1))

list_res_fausse_pos2 = [-4.999786766511124, -0.4997547420662774, 3.549115617573128, 7.185585307380122, 10.418455455972435, 13.155600249841392, 15.051534712120477, 15.655196579114605, 15.599764611538863, 15.605493655785121, 15.605550280417384, 15.605550839747492]
print('test fausse_pos2 : ', list(np.isclose(np.array(fausse_pos(-10, 40, lambda x: 5/(1+np.exp(-0.5*x+10))-0.5, 10**(-8))), np.array(list_res_fausse_pos2))).count(True) == len(list_res_fausse_pos2))

list_res_newton = [1.5, 1.4166666666666667, 1.4142156862745099]
print('test newton : ', list(np.isclose(np.array(newton(2, lambda x: x**2 - 2, lambda x: 2*x, 10**(-5))), np.array(list_res_newton))).count(True) == len(list_res_newton))
