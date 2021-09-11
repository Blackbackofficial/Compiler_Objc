from __future__ import annotations
from abc import ABC, abstractmethod
from typing import List
from queue import SimpleQueue, LifoQueue
from abc import abstractmethod
import os

BASE_DIR = os.path.dirname(os.path.abspath(__file__))


class Tree(object):
    @abstractmethod
    def meta(self):
        pass


class CommonTree(Tree):
    def __init__(self, data=None):
        self.children = []
        self.data = data

    def meta(self):
        return None


def getNextChar(char, union):
    next_char = chr(ord(char) + 1)
    syms = [x.mark for x in union]
    while next_char in syms:
        next_char = getNextChar(next_char, union)
    return next_char


class GrammarSymbol(object):
    def __init__(self, mark, idx=0):
        self.mark = mark
        self.idx = idx

    def __str__(self):
        res = f"{self.mark}"
        if self.idx:
            res += f"_{self.idx}"
        return res

    def __eq__(self, another):
        return self.mark == another.mark and self.idx == another.idx

    def get_next_idx(self):
        return GrammarSymbol(self.mark, self.idx + 1)


EPS_SYM = GrammarSymbol('eps')
FIN_SYM = GrammarSymbol('$')


class Rule(object):
    def __init__(self, lhs, rhs):
        self.lhs = lhs
        if len(rhs) > 1:
            self.rhs = [x for x in rhs if x != EPS_SYM]
        else:
            self.rhs = rhs

    def __str__(self):
        res = f"{self.lhs} ->"
        for sym in self.rhs:
            res += f" {sym}"
        return res


class Grammar(object):
    def __init__(self, nonterminal, terminal, rules: List[Rule], start_symbol) -> None:
        self.nonterminal = nonterminal
        self.terminal = terminal + [EPS_SYM]
        self.rules = rules
        self.start_symbol = start_symbol
        self.first_table = None
        self.follow_table = None

    def recursive_decent_parse(self, string):
        is_eval, num, tree = self.recursive_decent_parse_iter(self.start_symbol, string.split())
        return (is_eval and num == len(string.split())), tree

    def recursive_decent_parse_iter(self, nonterm: GrammarSymbol, string: List[str]):
        productions = [rule for rule in self.rules if rule.lhs == nonterm]
        productions.sort(key=lambda x: len(x.rhs), reverse=True)
        evaled = 0
        for rule in productions:
            idx_string = 0
            idx_rhs = 0
            cur_tree = CommonTree(rule.lhs)
            while idx_rhs < len(rule.rhs):
                if rule.rhs[0] == EPS_SYM:
                    return True, evaled, CommonTree(EPS_SYM)
                if not string:
                    break
                if rule.rhs[idx_rhs] in self.terminal:
                    if string[idx_string] == str(rule.rhs[idx_rhs]):
                        cur_tree.children.append(CommonTree(rule.rhs[idx_rhs]))
                        idx_rhs += 1
                        idx_string += 1
                    else:
                        break
                else:
                    is_eval, num_evaled, new_tree = self.recursive_decent_parse_iter(rule.rhs[idx_rhs],
                                                                                     string[idx_string:])
                    if is_eval:
                        idx_string += num_evaled
                        idx_rhs += 1
                        cur_tree.children.append(new_tree)
                    else:
                        break

            if idx_rhs == len(rule.rhs):
                evaled = idx_string
                return True, evaled, cur_tree

        return False, evaled, None

    def sort_rules(self):
        self.rules = sorted(self.rules, key=lambda x: self.nonterminal.index(x.lhs))

    def __str__(self):
        res = ""
        for nonterm in self.nonterminal:
            rules = [rule for rule in self.rules if rule.lhs == nonterm]
            res += f"{nonterm} -> "
            res += '|'.join([' '.join([str(x) for x in rule.rhs]) for rule in rules])
            res += '\n'
        return res

    def first(self, syms: List[GrammarSymbol]):
        res = []
        if not syms:
            return [EPS_SYM]
        if syms[0] in self.terminal:
            res.append(syms[0])
        else:
            out_rules = [rule for rule in self.rules if rule.lhs == syms[0]]
            eps_rules = [rule for rule in out_rules if EPS_SYM in rule.rhs]

            for rule in out_rules:
                res.extend([x for x in self.first(rule.rhs) if x not in res])

            if eps_rules:
                res.extend([x for x in self.first(syms[1:]) if x not in res])
        return res

    def firstTable(self):
        first_sets = [list() for sym in self.nonterminal]
        list_nonterm = list()
        changed = True
        while changed:
            changed = False
            for rule in self.rules:
                idx = self.nonterminal.index(rule.lhs)
                old_len = len(first_sets[idx])

                first_sets[idx].extend([x for x in self.first(rule.rhs) if x not in first_sets[idx]])

                if len(first_sets[idx]) > old_len:
                    changed = True
                list_nonterm.append(rule.lhs.mark)

        unique_nonterm = []
        for number in list_nonterm:
            if number in unique_nonterm:
                continue
            else:
                unique_nonterm.append(number)

        self.first_table = first_sets
        return first_sets, unique_nonterm

    def followTable(self):
        follow_sets = [list() for sym in self.nonterminal]
        list_nonterm = list()
        follow_sets[self.nonterminal.index(self.start_symbol)].append(FIN_SYM)

        changed = True
        while changed:
            changed = False
            for rule in self.rules:
                for idx in range(len(rule.rhs)):
                    if rule.rhs[idx] in self.nonterminal:
                        follow_sets_idx = self.nonterminal.index(rule.rhs[idx])
                        follow_len = len(follow_sets[follow_sets_idx])

                        rest_first = self.first(rule.rhs[idx + 1:])

                        for sym in rest_first:
                            if sym == EPS_SYM:
                                lhs_follow = follow_sets[self.nonterminal.index(rule.lhs)]
                                for x in lhs_follow:
                                    if x not in follow_sets[follow_sets_idx]:
                                        follow_sets[follow_sets_idx].append(x)
                            elif sym not in follow_sets[follow_sets_idx]:
                                follow_sets[follow_sets_idx].append(sym)

                        if len(follow_sets[follow_sets_idx]) > follow_len:
                            changed = True
                list_nonterm.append(rule.lhs.mark)

        self.follow_table = follow_sets
        return follow_sets


class GrammarBuilder(ABC):
    @property
    @abstractmethod
    def grammar(self) -> Grammar:
        pass

    @abstractmethod
    def get_grammar(self, arg) -> None:
        pass


class GrammarReader(GrammarBuilder):
    def __init__(self) -> None:
        self._grammar = None

    def reset(self) -> None:
        self._grammar = None

    def grammar(self) -> Grammar:
        return self._grammar

    def get_grammar(self, filename) -> None:
        with open(f"{BASE_DIR}/grammars/{filename}", "r") as f:
            nonterminal_count = int(f.readline())
            nonterminal = [GrammarSymbol(symb) for symb in f.readline().split()]

            terminal_count = int(f.readline())
            terminal = [GrammarSymbol(symb) for symb in f.readline().split()]

            rules_count = int(f.readline())
            rules = []
            for i in range(rules_count):
                lhs, rhs = f.readline().split(" -> ")
                lhs = [GrammarSymbol(symb) for symb in lhs.split()][0]
                rhs = [GrammarSymbol(symb) for symb in rhs.split()]
                rules.append(Rule(lhs, rhs))
            rules = sorted(rules, key=lambda x: nonterminal.index(x.lhs))
            start_symbol = GrammarSymbol(f.readline())

        self._grammar = Grammar(nonterminal, terminal, rules, start_symbol)


class GrammarConverter(GrammarBuilder):
    def __init__(self) -> None:
        self._grammar = None

    def reset(self):
        self._grammar = None

    def grammar(self) -> Grammar:
        return self._grammar

    def get_grammar(self, grammar: Grammar) -> None:
        self._grammar = grammar

    def delete_left_rec(self):
        grammar = self._grammar

        new_rules = []
        new_nonterm = [x for x in grammar.nonterminal]
        # следующий символ to @ - A
        next_char = GrammarSymbol(getNextChar('@', grammar.nonterminal + grammar.terminal))

        for nonterm in grammar.nonterminal:
            rules = [y for y in grammar.rules if y.lhs == nonterm]
            cleared_rules = []
            for rule in rules:
                q = SimpleQueue()
                q.put(rule)
                while not q.empty():
                    cur_rule = q.get()
                    if cur_rule.rhs[0] in grammar.terminal or (new_nonterm.index(cur_rule.rhs[0])
                                                               >= new_nonterm.index(
                                cur_rule.lhs)):  # находится после новых нетермов
                        cleared_rules.append(cur_rule)
                    else:
                        rules_to_insert = [y.rhs + cur_rule.rhs[1:] for y in new_rules if y.lhs == cur_rule.rhs[0]]
                        if not rules_to_insert:
                            rules_to_insert = [y.rhs + cur_rule.rhs[1:] for y in grammar.rules if
                                               y.lhs == cur_rule.rhs[0]]
                        for i in rules_to_insert:
                            q.put(Rule(cur_rule.lhs, i))

            rules = cleared_rules

            # если прямая рекурсия
            rec_rules = [rule for rule in rules if nonterm == rule.rhs[0]]
            if rec_rules:
                com_rules = [rule for rule in rules if rule not in rec_rules]

                new_rules.extend([Rule(nonterm, rule.rhs + [next_char]) for rule in com_rules])
                new_rules.extend([Rule(nonterm, rule.rhs) for rule in com_rules])
                new_rules.extend([Rule(next_char, rule.rhs[1:] + [next_char]) for rule in rec_rules])
                new_rules.extend([Rule(next_char, rule.rhs[1:]) for rule in rec_rules])
                new_nonterm.insert(new_nonterm.index(nonterm) + 1, next_char)
                next_char = GrammarSymbol(getNextChar(next_char.mark, new_nonterm + grammar.terminal))
            else:
                new_rules.extend(rules)

        self._grammar.rules = new_rules
        self._grammar.nonterminal = new_nonterm

    def delete_non_gen_rules(self):
        grammar = self._grammar

        is_gen = {str(k): False for k in grammar.nonterminal}
        checked = {str(k): False for k in grammar.nonterminal}
        concerned_rules = {str(k): [i for i in range(len(grammar.rules)) if k in grammar.rules[i].rhs]
                           for k in grammar.nonterminal}
        q = SimpleQueue()
        counter = {}
        for rule in grammar.rules:
            str_rule = str(rule)
            counter[str_rule] = 0
            for sym in rule.rhs:
                if sym in grammar.nonterminal:
                    counter[str_rule] += 1

            if not counter[str_rule]:
                q.put(rule.lhs)
                is_gen[str(rule.lhs)] = True
                checked[str(rule.lhs)] = True

        while not q.empty():
            nonterm = q.get()
            for rule in concerned_rules[str(nonterm)]:
                rule = grammar.rules[rule]
                for i in range(rule.rhs.count(nonterm)):
                    counter[str(rule)] -= 1
                if not counter[str(rule)] and not checked[str(rule.lhs)]:
                    q.put(rule.lhs)
                    checked[str(rule.lhs)] = True
                    is_gen[str(rule.lhs)] = True

        new_rules = []
        for rule in grammar.rules:
            not_get = False
            for sym in rule.rhs:
                if sym in grammar.nonterminal:
                    if not is_gen[str(sym)]:
                        not_get = True
            if not not_get:
                new_rules.append(rule)

        self._grammar.rules = new_rules

    def delete_unreachable_rules(self):
        grammar = self._grammar

        reachable = [grammar.start_symbol]
        q = SimpleQueue()
        q.put(grammar.start_symbol)

        while not q.empty():
            nonterm = q.get()
            for rule in grammar.rules:
                if nonterm == rule.lhs:
                    for x in rule.rhs:
                        if x in grammar.nonterminal and x not in reachable:
                            q.put(x)
                            reachable.append(x)

        new_rules = [rule for rule in grammar.rules if rule.lhs in reachable]

        self._grammar.rules = new_rules
        self._grammar.nonterminal = reachable

    def delete_useless_rules(self):
        self.delete_non_gen_rules()
        self.delete_unreachable_rules()
