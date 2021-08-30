from Compiler_Objc.lab2.models.grammar import *


def LeftRecursion(g: Grammar):
    MakeUndirectRecursive(g)

    setsss = FindRecusive(g)
    for i in range(len(g.Nonterms)):
        if g.Nonterms[i] in setsss:
            MakeIndirectRecursion(g, g.Nonterms[i])
    return


def FindRecusive(g: Grammar):
    letters = []
    for i in range(len(g.Nonterms)):
        a, b = SortRules(g, g.Nonterms[i])
        for j in range(len(a)):
            if a[j].Left == a[j].Right[0]:
                letters.append(a[j].Left)
    return set(letters)


def MakeIndirectRecursion(g: Grammar, letter: str):
    rules = []

    a, b = SortRules(g, letter)
    newLetter = letter + '*'

    for i in range(len(b)):
        rule = b[i]
        if rule.Right[len(rule.Right) - 1] != "eps":
            rule.Right.append(newLetter)

    for i in range(len(a)):
        a[i].Left = newLetter
        a[i].Right = a[i].Right[1:]
        a[i].Right.append(newLetter)

    empty = Rule()
    empty.Left = newLetter
    empty.Right = ["eps"]
    a.append(empty)

    for i in range(len(g.Rules)):
        if g.Rules[i].Left != letter and g.Rules[i].Left != newLetter:
            rules.append(g.Rules[i])
    rules.extend(a)
    rules.extend(b)

    g.Nonterms.append(newLetter)

    g.Rules = rules
    return


def CheckDirectRecursion(g: Grammar):
    for i in range(len(g.Nonterms)):
        a, b = SortRules(g, g.Nonterms[i])
        for j in range(len(a)):
            if a[j].Left == a[j].Right[0]:
                return True
    return False


def MakeUndirectRecursive(g: Grammar):
    rules = []
    for i in range(len(g.Nonterms)):
        ai, bi = SortRules(g, g.Nonterms[i])
        ai.extend(bi)
        for j in range(i):
            aj, bj = SortRules(g, g.Nonterms[j])
            aj.extend(bj)
            rule = ReplaceRules(ai, aj, g.Nonterms[j])
            rules.extend(rule)

    a0, b0 = SortRules(g, g.Nonterms[0])
    rules.extend(a0)
    rules.extend(b0)

    g.Rules = rules
    return rules


def ReplaceRules(rule1, rule2, letter: str):
    rules3 = []
    for i in range(len(rule1)):
        if letter == (rule1[i]).Right[0]:
            for j in range(len(rule2)):
                rule = Rule()
                rule.Left = rule1[i].Left
                rule.Right = rule2[j].Right
                rule.Right.extend(rule1[i].Left[1:])
                rules3.append(rule)
        else:
            rules3.append(rule1[i])
    return rules3


def SortRules(g: Grammar, letter: str):
    a, b = [], []
    for i in range(len(g.Rules)):
        rule = g.Rules[i]
        if rule.Left == letter:
            if rule.Left == rule.Right[0] and len(rule.Right) != 1:
                a.append(rule)
            else:
                b.append(rule)
    return a, b


def SetRules(g: Grammar):
    setss = set()
    newRules = []
    for i in range(len(g.Rules)):
        if g.Rules[i] not in setss:
            setss.add(g.Rules[i])
            newRules.append(g.Rules[i])

    g.Rules = newRules


def FindFactor(g: Grammar):
    letters = set()
    for i in range(len(g.Rules)):
        for j in range(i + 1, len(g.Rules)):
            if g.Rules[i].Left == g.Rules[j].Left:
                if g.Rules[i].Right[0] == g.Rules[j].Right[0]:
                    letters.add(g.Rules[i].Left)

    return letters


def DeleteFactor(g: Grammar, letter: str):
    a, b = SortRules(g, letter)

    b.sort(key=lambda x: x.Right[0])

    replace = []
    rules = []
    replace.append(b[0])
    for i in range(1, len(b)):
        if b[i].Right[0] == replace[0].Right[0]:
            replace.append(b[i])
        else:
            mx = FindMax(replace)
            for j in range(replace):
                replace[j].Right = replace[j].Right[mx:]
                replace[j].Left = replace[j].Left + '^'
            g.Nonterms = replace[0].Left + "^"
            rules.extend(replace)
            replace = []

    if len(replace) != 1:
        mx = FindMax(replace)
        for j in range(replace):
            replace[j].Right = replace[j].Right[mx:]
            replace[j].Left = replace[j].Left + '^'

    for i in range(len(g.Rules)):
        if g.Rules[i].Left != letter and g.Rules[i].Left != letter + '^':
            rules.append(g.Rules[i])
    rules.extend(a)

    g.Nonterms = replace[0].Left + "^"

    g.Rules = rules

    return


def FindMax(rules):
    mx = len(rules[0])

    for i in range(len(rules)):
        if mx > len(rules[i].Right):
            mx = rules[i].Right
        for j in range(max):
            if rules[i].Right[j] != rules[0].Right[j]:
                mx = j
                break
    return mx
