from Compiler_Objc.labs.lab2.recursion.left import *
from Compiler_Objc.labs.lab2.useless.delete import *


def main():
    grammar = Grammar()

    grammar.Nonterms = ["E", "T", "F"]
    #
    # rule = Rule()
    # rule.Left = "E"
    # rule.Right = ["E", "+", "T"]
    # grammar.Rules.append(rule)
    #
    # rule = Rule()
    # rule.Left = "E"
    # rule.Right = ["T"]
    # grammar.Rules.append(rule)
    #
    # rule = Rule()
    # rule.Left = "T"
    # rule.Right = ["T", "*", "F"]
    # grammar.Rules.append(rule)
    #
    # rule = Rule()
    # rule.Left = "T"
    # rule.Right = ["F"]
    # grammar.Rules.append(rule)
    #
    # rule = Rule()
    # rule.Left = "F"
    # rule.Right = ["a"]
    # grammar.Rules.append(rule)

    # Факторизация
    rule = Rule()
    rule.Left = "E"
    rule.Right = ["T", "+", "F"]
    grammar.Rules.append(rule)

    rule = Rule()
    rule.Left = "E"
    rule.Right = ["T", "+", "T"]
    grammar.Rules.append(rule)

    rule = Rule()
    rule.Left = "T"
    rule.Right = ["F"]
    grammar.Rules.append(rule)

    rule = Rule()
    rule.Left = "F"
    rule.Right = ["a"]
    grammar.Rules.append(rule)

    print("Start: ")
    for rul in grammar.Rules:
        print(rul)

    LeftRecursion(grammar)
    print("Left Recursion: ")
    for rul in grammar.Rules:
        print(rul)
    SetRules(grammar)
    print("Set Rules: ")
    for rul in grammar.Rules:
        print(rul)
    MakeUnfactored(grammar)
    SetRules(grammar)
    print("Set Rules: ")
    for rul in grammar.Rules:
        print(rul)
    print("Make Unfactored: ")
    for rul in grammar.Rules:
        print(rul)
    DeleteUseless(grammar)
    print("Delete Useless: ")
    for rul in grammar.Rules:
        print(rul)
    print("No Terms:", grammar.Nonterms)
    for rul in grammar.Rules:
        print(rul)


if __name__ == '__main__':
    main()
