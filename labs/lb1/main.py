# from graphviz import Digraph

from FA import *
from parser import *
import graphviz


def model_mfa(word, dfa):
    now = dfa.start
    i = 0
    while i < len(word):
        print("state: " + now.name + ", char: " + word[i])
        now = now.transitions[word[i]]
        print("next state: " + now.name)
        i += 1
    if now.is_term:
        print("state: " + now.name + " is final")
    else:
        print("state: " + now.name + " is not final")


def main():
    pars = Parser(TEST)
    nfa = NFA(pars.ex)
    with open('../lb1/output/NFA.gv', "w+") as f:
        data1 = nfa.print_NFA(f)
        f.write(data1)
        f.close()

    dfa = DFA(nfa.nfa_stack[0].start, TEST_ALPH, nfa.nfa_stack[0].end)
    with open('../lb1/output/DFA.gv', "w+") as f:
        data2 = dfa.print_DFA(f)
        f.write(data2)
        f.close()

    for i in dfa.queue:
        i.label = False

    mfa = MFA(dfa, TEST_ALPH)
    with open('../lb1/output/MFA.gv', "w") as f:
        data3 = mfa.print_MFA(f)
        f.write(data3)
        f.close()

    graphviz.render('dot', 'png', '../lb1/output/NFA.gv')
    graphviz.render('dot', 'png', '../lb1/output/DFA.gv')
    graphviz.render('dot', 'png', '../lb1/output/MFA.gv')
    # model_mfa(INPUT_TEST, dfa)


if __name__ == "__main__":
    main()
