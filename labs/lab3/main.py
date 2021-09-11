from classes import *
# from rd_parser import *
from print_tree import print_tree


class PrintTree(print_tree):
    def get_children(self, node):
        return node.children

    def get_node_str(self, node):
        return str(node.data)


# in_book.txt
# g1_algo_pascal.txt
def start():
    reader = GrammarReader()
    converter = GrammarConverter()
    reader.get_grammar('in_book.txt')
    converter.get_grammar(reader.grammar())

    # PRINT START GRAMMAR
    print("Start Grammar:\n")
    converter.grammar().sort_rules()
    print(converter.grammar()); print("\n")

    # DELETE LEFT RECURSION
    print("Delete left recursion:\n")
    converter.delete_left_rec()
    print(converter.grammar()); print("\n")

    # DELETE USELESS RULES
    print("Delete useless rules:\n")
    converter.delete_useless_rules()
    print(converter.grammar()); print("\n")

    # FIRST
    firstTable, first_nonterm = converter.grammar().firstTable()
    for i in range(len(firstTable)):
        q = f"FIRST({first_nonterm[i]}) = " + "{"
        for j in range(len(firstTable[i])):
            q += str(firstTable[i][j].mark) + " "
        q += "}"
        print(q)
    print('\n')

    # FOLLOW
    followTable = converter.grammar().followTable()
    for i in range(len(followTable)):
        q = f"FOLLOW({first_nonterm[i]}) = " + "{"
        for j in range(len(followTable[i])):
            q += str(followTable[i][j].mark) + " "
        q += "}"
        print(q)
    print('\n')

    # ENTER EVAL STRING
    string = input("Enter string: ")
    i, u = converter.grammar().recursive_decent_parse(string)
    # parser = RecursiveDecent()
    res, tree = converter.grammar().recursive_decent_parse(string)
    if res:
        PrintTree(tree)
        print("Save in output.txt")
    else:
        print("String cannot be evaluated with lab3")


if __name__ == '__main__':
    start()
