class Grammar:
    def __init__(self):
        self.Name = ""
        self.Nonterms = []
        self.Terms = []
        self.Rules = []
        self.Start = None


class Rule:
    def __init__(self):
        self.Left = ""
        self.Right = []

    def __repr__(self):
        return "%s -> %s;" % (self.Left, self.Right)
