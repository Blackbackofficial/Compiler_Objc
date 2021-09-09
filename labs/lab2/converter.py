from grammar import Grammar, Rules


# er_productions - подстановки продукций
# Nonterminals: ['S', 'A']
# Terminals: ['a', 'b', 'c', 'd']
# Start: S
# Productions:
# S -> Aa | b
# A -> Ac | Sd
# ER:  [A -> Aad, A -> bd]

class GrammarConverter:
    """
    Class to delete eps-productions and left recursion in grammar
    """

    @staticmethod
    def delete_left_recursion(g: Grammar) -> Grammar:
        grammar = Grammar(g.non_terminals, g.terminals, g.productions[:], g.start_symbol)

        non_terms = grammar.non_terminals[:]
        n = len(non_terms)
        for i in range(n):  # верхний цикл проходит по НЕТЕРМИНАЛАМ
            # print('{0}: {1}'.format(i, non_terms[i]))
            for j in range(i):
                cur_productions = []
                for p in grammar.productions:
                    if p.left == non_terms[i] and p.right[0] == non_terms[j] and len(p.right) > 1:  # Ai -> Aj a
                        er_productions = grammar.er_productions(p)
                        cur_productions.extend(er_productions[:])
                    else:
                        cur_productions.append(p)
                grammar.update_productions(cur_productions)

            grammar.delete_directly_left_recursion(non_terms[i])
        return grammar

    @staticmethod
    def factorization(g: Grammar) -> Grammar:
        grammar = Grammar(g.non_terminals, g.terminals, g.productions[:], g.start_symbol)
        for t in range(len(g.productions)):
            sym = GrammarConverter.find_factor(grammar)
            for key in sym:
                GrammarConverter.delete_factor(g, key)
        return g

    @staticmethod
    def SortRules(g: Grammar, letter: str) -> list:  # a - где есть этот нетерминал, b - где нет этого не терминала
        b = list()
        for i in range(len(g.productions)):
            rule = g.productions[i]
            if rule.left == letter:
                b.append(rule)
        return b

    @staticmethod
    def find_factor(g: Grammar) -> set:
        letters = set()
        for i in range(len(g.productions)):
            for j in range(i + 1, len(g.productions)):
                if g.productions[i].left == g.productions[j].left:
                    if g.productions[i].right[0] == g.productions[j].right[0]:
                        letters.add(g.productions[i].left)

        if len(letters) == 0:
            return letters

        return letters

    @staticmethod
    def delete_factor(g: Grammar, letter: str):
        b = GrammarConverter.SortRules(g, letter)
        for i in range(len(b)):  # удаляем правила не участвующие в факторе
            for j in range(i + 1, len(b)):
                if b[i].left == b[j].left:
                    if b[i].right[0] == b[j].right[0]:
                        continue
                    else:
                        b.pop(j)

        b.sort(key=lambda x: x.right[0])

        replace = []
        rules = []
        replace.append(b[0])
        for i in range(1, len(b)):
            if b[i].right[0] == replace[0].right[0]:
                replace.append(b[i])

        if len(replace) != 1:
            mx = GrammarConverter.FindMax(replace)
            rel = Rules(replace[0].left, replace[0].right[:mx])
            rel.right.append(rel.left + "`")
            rules.append(rel)
            for j in range(len(replace)):
                if len(replace[j].right) != 1:  # если терминал терм то -eps
                    replace[j].right = replace[j].right[mx:]
                    replace[j].left = replace[j].left + "`"
                elif replace[j].right[0] in g.terminals:
                    print("это терминал", replace[j].right[0])  # если не нужно то убрать

        g.non_terminals.append(replace[0].left)

        g.productions.extend(rules)

    @staticmethod
    def FindMax(rules) -> int:
        mx = len(rules[0].right)

        for i in range(len(rules)):
            if mx > len(rules[i].right):
                mx = len(rules[i].right)
            for j in range(mx):
                if rules[i].right[j] != rules[0].right[j]:
                    mx = j
                    break
        return mx

    # DELETE USELESS
    @staticmethod
    def delete_useless(g: Grammar) -> Grammar:
        symb = GrammarConverter.FindUseless(g)
        newNonterms = []

        for i in range(len(g.non_terminals)):
            if g.non_terminals[i] in symb:
                newNonterms.append(g.non_terminals[i])

        g.Nonterms = newNonterms

        newRules = []

        for i in range(len(g.productions)):
            flag = True
            for j in range(len(g.productions[i].right)):
                if not GrammarConverter.isTerm(g.productions[i].right[j]):
                    if not g.productions[i].right[j] in symb:
                        flag = False
            if flag:
                newRules.append(g.productions[i])
        g.productions = newRules
        g.non_terminals.clear()
        g.non_terminals.extend(symb)
        return g

    def FindUseless(g: Grammar):
        usefull = set()
        for i in range(len(g.productions)):
            flag = True
            for j in range(len(g.productions[i].right)):
                if not GrammarConverter.isTerm(g.productions[i].right[j]):
                    flag = False
            if flag:
                usefull.add(g.productions[i].left)

        oldsize = 0
        while oldsize != len(usefull):
            oldsize = len(usefull)
            for i in range(len(g.productions)):
                flag = True
                for j in range(len(g.productions[i].right)):
                    if not GrammarConverter.isTerm(g.productions[i].right[j]):
                        if not g.productions[i].right[j] in usefull:
                            flag = False
                            break
                if flag:
                    usefull.add(g.productions[i].left)

        return usefull

    def isTerm(self: str):
        return self.lower() == self
