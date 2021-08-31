from Compiler_Objc.labs.lab2.recursion.left import *
from Compiler_Objc.labs.lab2.useless.delete import *


def start():
    grammar = Grammar()

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
    print("No Terms:", grammar.Nonterms)
    for rul in grammar.Rules:
        print(rul)


if __name__ == '__main__':
    start()