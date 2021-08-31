from Compiler_Objc.labs.lab2.models.grammar import *


def FindUseless(g: Grammar):
    usefull = set()
    for i in range(len(g.Rules)):
        flag = True
        for j in range(len(g.Rules[i].Right)):
            if not isTerm(g.Rules[i].Right[j]):
                flag = False
        if flag:
            usefull.add(g.Rules[i].Left)

    oldsize = 0
    while oldsize != len(usefull):
        oldsize = len(usefull)
        for i in range(len(g.Rules)):
            flag = True
            for j in range(len(g.Rules[i].Right)):
                if not isTerm(g.Rules[i].Right[j]):
                    if not g.Rules[i].Right[j] in usefull:
                        flag = False
                        break
            if flag:
                usefull.add(g.Rules[i].Left)

    return usefull


def DeleteUseless(g: Grammar):
    symb = FindUseless(g)

    newNonterms = []

    for i in range(len(g.Nonterms)):
        if g.Nonterms[i] in symb:
            newNonterms.append(g.Nonterms[i])

    g.Nonterms = newNonterms

    newRules = []

    for i in range(len(g.Rules)):
        flag = True
        for j in range(len(g.Rules[i].Right)):
            if not isTerm(g.Rules[i].Right[j]):
                if not g.Rules[i].Right[j] in symb:
                    flag = False
        if flag:
            newRules.append(g.Rules[i])
    g.Rules = newRules
    return


def isTerm(letter: str):
    return letter.lower() == letter
