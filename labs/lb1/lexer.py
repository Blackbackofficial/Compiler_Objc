from tokens import *


class Lexer:
    def __init__(self, source):
        self.source = source
        self.symbols = {'(': 'LEFT_PAREN', ')': 'RIGHT_PAREN', '*': 'STAR', '|': 'ALT', '\x08': 'CONCAT', '+': 'PLUS'}
        self.current = 0
        self.length = len(self.source)

    def get_token(self):
        if (self.current < self.length):
            work = self.source[self.current]
            self.current += 1
            if work not in self.symbols.keys():
                return Tokens ('CHAR', work)
            else:
                return Tokens(self.symbols[work], work)
        else:
            return Tokens('NONE', '')