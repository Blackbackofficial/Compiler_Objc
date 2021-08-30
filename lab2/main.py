from models.grammar import Grammar, Rule
from recursion.left import *
from useless.delete import *


def start():
    grammar = Grammar()
    print("DEWDWDEWDWDEWED")

    grammar.Nonterms = ["E", "T", "F"]
    rule = Rule()
    rule.Left = "E"
    rule.Right = ["E", "+", "T"]
    grammar.Rules.append(rule)

    rule = Rule()
    rule.Left = "E"
    rule.Right = ["T"]
    grammar.Rules.append(rule)

    rule = Rule()
    rule.Left = "T"
    rule.Right = ["T", "*", "F"]
    grammar.Rules.append(rule)

    rule = Rule()
    rule.Left = "T"
    rule.Right = ["F"]
    grammar.Rules.append(rule)

    rule = Rule()
    rule.Left = "F"
    rule.Right = ["a"]
    grammar.Rules.append(rule)

    LeftRecursion(grammar)
    SetRules(grammar)
    DeleteUseless(grammar)
    print("%v", grammar)


if __name__ == '__main__':
    start()
