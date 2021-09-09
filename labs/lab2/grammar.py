from collections import defaultdict


class Rules:
    def __init__(self, left, right):
        self.left = left
        self.right = right

    def __repr__(self):
        return f"{self.left} -> {''.join(self.right)}"


class Grammar:
    """
    Class describing grammar
    """
    EPS = "EPS"

    def __init__(self, non_terminals: list, terminals: list, productions: list, start_symbol: str):
        self.non_terminals = non_terminals
        self.terminals = terminals
        self.productions = productions
        self.start_symbol = start_symbol

    def er_productions(self, production):
        left = production.left
        right = production.right
        first_right = right.pop(0)
        new_productions = []

        for p in self.productions:
            if p.left == first_right:
                new_productions.append(Rules(left, p.right + right))

        print("ER: ", new_productions)
        return new_productions

    def delete_directly_left_recursion(self, left):
        # print('\n~~~~~~~~~~~~~~~~~~~~')
        # print('delete_directly_left_recursion')
        # print('~~~~~~~~~~~~~~~~~~~~')
        # print('left: {0}'.format(left))
        # print('self.productions: {0}'.format(self.productions))
        for p in self.productions:
            if p.left == left and p.right[0] == left:
                break
        else:
            return  # no left recursion

        new_productions = []
        new_left = left + "'"
        self.non_terminals.append(new_left)

        eps_in_productions = False
        for p in self.productions:
            # print('\ncur production {0}'.format(p))
            if p.left == left:
                if p.right[0] == left:
                    new_productions.append(Rules(new_left, p.right[1:] + [new_left]))
                    if not eps_in_productions:
                        new_productions.append(Rules(new_left, self.EPS))
                        eps_in_productions = True

                    # new_productions.append(Rules(new_left, p.right[1:]))
                    # print('new_productions {0}'.format(new_productions))
                else:
                    new_productions.append(Rules(left, p.right[:] + [new_left]))
                    # new_productions.append(Rules(left, p.right[:]))
                    # print('new_productions {0}'.format(new_productions))
            else:
                new_productions.append(p)
                # print('new_productions {0}'.format(new_productions))

        self.update_productions(new_productions)

    def update_productions(self, new_productions):
        self.productions.clear()
        self.productions.extend(new_productions[:])

    def _repr_productions(self):
        productions = defaultdict(list)
        for production in self.productions:
            productions[production.left].append("".join(production.right))

        repr_string = ""
        for left, rights in productions.items():
            repr_string += f"\n{left} -> {' | '.join(rights)}"

        return repr_string

    def __repr__(self):
        return "Nonterminals: {0}\nTerminals: {1}\nStart: {2}\n" \
               "Productions:{3}".format(self.non_terminals, self.terminals, self.start_symbol, self._repr_productions())
