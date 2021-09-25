from FA import *
from parser_class import *
import graphviz


def model_mfa(word, dfa):
    i = 0
    while i < len(word):
        print("state: " + dfa.start.name + ", char: " + word[i])
        dfa.start = dfa.start.transitions[word[i]]
        print("next state: " + dfa.start.name)
        i += 1
    if dfa.start.is_term:
        print("state: " + dfa.start.name + " the string is valid")
    else:
        print("state: " + dfa.start.name + " the string is NOT valid")


def main():
    pars = Parser(TEST)
    nfa = NFA(pars.ex)
    with open('../lab1/output/NFA.gv', "w+") as f:
        data1 = nfa.print_NFA()
        f.write(data1)
        f.close()

    dfa = DFA(nfa.nfa_stack[0].start, TEST_ALPHA, nfa.nfa_stack[0].end)
    with open('../lab1/output/DFA.gv', "w+") as f:
        data2 = dfa.print_DFA()
        f.write(data2)
        f.close()

    for i in dfa.queue:
        i.label = False

    mfa = MFA(dfa, TEST_ALPHA)
    with open('../lab1/output/MFA.gv', "w+") as f:
        data3 = mfa.print_MFA()
        f.write(data3)
        f.close()

    graphviz.render('dot', 'png', '../lab1/output/NFA.gv')
    graphviz.render('dot', 'png', '../lab1/output/DFA.gv')
    graphviz.render('dot', 'png', '../lab1/output/MFA.gv')
    model_mfa(INPUT_TEST, dfa)


if __name__ == "__main__":
    main()
