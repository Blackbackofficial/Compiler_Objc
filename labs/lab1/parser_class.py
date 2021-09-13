from lexer_class import *
from tokens_class import *


class Parser:

    def __init__(self, source):
        self.lexer = Lexer(source)
        self.tokens = []
        work_concat = Tokens('CONCAT', '\x08')
        work1 = self.lexer.get_token()
        while work1.name != 'NONE':
            self.tokens.append(work1)
            work2 = self.lexer.get_token()
            if self.check_concat(work1, work2):
                self.tokens.append(work_concat)
            work1 = work2
        self.ex = []
        self.st = []
        self.RPN()

    @staticmethod
    def check_concat(work1, work2):
        if work1.name == 'CHAR' and work2.name == 'CHAR':
            return True
        elif work1.name == 'CHAR' and work2.name == 'LEFT_PAREN':
            return True
        elif work1.name == 'RIGHT_PAREN' and work2.name == 'CHAR':
            return True
        elif work1.name == 'STAR' and work2.name == 'LEFT_PAREN':
            return True
        elif work1.name == 'PLUS' and work2.name == 'LEFT_PAREN':
            return True
        elif work1.name == 'RIGHT_PAREN' and work2.name == 'LEFT_PAREN':
            return True
        elif work1.name == 'STAR' and work2.name == 'CHAR':
            return True
        elif work1.name == 'PLUS' and work2.name == 'CHAR':
            return True
        else:
            return False

    def RPN(self):
        i = 0
        while i < len(self.tokens):
            if self.tokens[i].name == 'CHAR' or self.tokens[i].name == 'STAR' or self.tokens[i].name == 'PLUS':
                self.ex.append(self.tokens[i])
            elif self.tokens[i].name == 'LEFT_PAREN' or self.tokens[i].name == 'ALT' or self.tokens[i].name == 'CONCAT':
                self.st.append(self.tokens[i])
            else:
                work = self.st.pop()
                while work.name != "LEFT_PAREN":
                    self.ex.append(work)
                    work = self.st.pop()
            i += 1
        while self.st:
            work = self.st.pop()
            self.ex.append(work)
