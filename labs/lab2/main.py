from converter import GrammarConverter
from grammar_file import GrammarFile


def start(filename):
    try:
        g = GrammarFile.create_grammar_from_file(filename)
    except Exception as e:
        print(e)
    else:
        print("Loaded grammar:\n{}".format(g))
        # Step 1
        grammar = GrammarConverter.delete_left_recursion(g)
        print("\nWithout left recursion:\n{}".format(grammar))
        # Step 2
        grammar = GrammarConverter.factorization(grammar)
        print("\nWithout left factorization:\n{}".format(grammar))
        # Step 3
        grammar = GrammarConverter.delete_useless(grammar)
        GrammarFile.save_grammar_to_file(grammar, 'out.json')
        print("\nWithout useless:\n{}".format(grammar))


if __name__ == "__main__":
    start('test/gr1.json')
