from json import load, dump
from grammar import Grammar, Rules


class GrammarFile:
    """
    Class to convert grammar to/from file
    """

    @staticmethod
    def create_grammar_from_file(filename) -> Grammar:

        with open(filename, "r") as f:
            grammar_json = load(f)

        try:
            non_terms = grammar_json["non_term"]
            terms = grammar_json["term"]
            start = grammar_json["start"]
            productions = grammar_json["productions"]
        except KeyError:
            raise Exception("Incorrect fields in `{}`".format(filename))

        if start not in non_terms:
            raise Exception("Incorrect start symbol in grammar")

        intersect = set(terms) & set(non_terms)
        if intersect:
            raise Exception("Terminals and non-terminals have the same symbols: {}".format(intersect))

        grammar_productions = []
        for p in productions:
            left = p["l"]
            right = p["r"]
            if left not in non_terms or not set(right).issubset(set(non_terms + terms)):
                raise Exception("Incorrect production: {} -> {}".format(left, "".join(right)))
            grammar_productions.append(Rules(left, right))

        return Grammar(non_terminals=non_terms, terminals=terms, start_symbol=start, productions=grammar_productions)

    @staticmethod
    def save_grammar_to_file(g: Grammar, filename: str) -> None:

        grammar_json = {"non_term": g.non_terminals, "term": g.terminals, "start": g.start_symbol, "productions": []}

        for production in g.productions:
            grammar_json["productions"].append({"l": production.left, "r": production.right})

        with open(filename, "w") as f:
            dump(grammar_json, f)
