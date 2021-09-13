TEST = "(a|b)*(ab)"
TEST_ALPH = ['a', 'b']
INPUT_TEST = "(a|b)*(ab)"


class MPoint:
    def __init__(self, name):
        self.name = name
        self.transitions = {}
        self.eps = []
        self.label = False
        self.is_term = False


class MBlock:
    def __init__(self, start, end):
        self.start = start
        self.end = end


class NFA:
    def __init__(self, token):
        self.nfa_stack = []
        self.stack = []
        self.tokens = token
        i = 1
        for t in self.tokens:
            if t.name == 'CHAR':
                i = self.char_handler(i, t)
            elif t.name == 'STAR':
                i = self.star_handler(i, t)
            elif t.name == 'PLUS':
                i = self.plus_handler(i, t)
            elif t.name == "ALT":
                i = self.alt_handler(i, t)
            else:
                self.concat_handler(t)
        self.nfa_stack[0].end.is_term = True
        self.nfa_stack[0].start.name = "Start"

    def char_handler(self, i, t):
        work1 = MPoint('s' + str(i))
        i += 1
        work2 = MPoint('s' + str(i))
        i += 1
        work1.transitions[t.value] = work2
        block = MBlock(work1, work2)
        self.nfa_stack.append(block)
        self.stack.append(work1)
        self.stack.append(work2)
        return i

    def star_handler(self, i, t):
        workblok = self.nfa_stack.pop()
        P1 = MPoint('s' + str(i))
        i += 1
        P2 = MPoint('s' + str(i))
        i += 1
        workblok.end.eps.append(workblok.start)
        workblok.end.eps.append(P2)
        P1.eps.append(workblok.start)
        P1.eps.append(P2)
        workblok = MBlock(P1, P2)
        self.nfa_stack.append(workblok)
        self.stack.append(P1)
        self.stack.append(P2)
        return i

    def plus_handler(self, i, t):
        workblok = self.nfa_stack.pop()
        P1 = MPoint('s' + str(i))
        i += 1
        P2 = MPoint('s' + str(i))
        i += 1
        workblok.end.eps.append(workblok.start)
        workblok.end.eps.append(P2)
        P1.eps.append(workblok.start)
        workblok = MBlock(P1, P2)
        self.nfa_stack.append(workblok)
        self.stack.append(P1)
        self.stack.append(P2)
        return i

    def alt_handler(self, i, t):
        P1 = MPoint('s' + str(i))
        i += 1
        P2 = MPoint('s' + str(i))
        i += 1
        workblok1 = self.nfa_stack.pop()
        workblok2 = self.nfa_stack.pop()
        P1.eps.append(workblok1.start)
        P1.eps.append(workblok2.start)
        workblok1.end.eps.append(P2)
        workblok2.end.eps.append(P2)
        workblok = MBlock(P1, P2)
        self.nfa_stack.append(workblok)
        self.stack.append(P1)
        self.stack.append(P2)
        return i

    def concat_handler(self, t):
        worckbloc2 = self.nfa_stack.pop()
        worckbloc1 = self.nfa_stack.pop()
        worckbloc1.end.eps.append(worckbloc2.start)
        worckbloc = MBlock(worckbloc1.start, worckbloc2.end)
        self.nfa_stack.append(worckbloc)

    def print_NFA(self, f):
        string = 'digraph G {\n'
        for i in self.stack:
            string += str(i.name) + '\n'
            for j in i.eps:
                string += str(i.name) + '->' + str(j.name) + '\n'
            for j in i.transitions:
                string += str(i.name) + '->' + str(i.transitions[j].name) + '[label=' + str(j) + ']' + '\n'
        return string + '}'

class DFA:
    def __init__(self, nfa, alph, nfa_end):
        self.start = MPoint('s1')
        j = 2
        work = []
        work.append(nfa)
        self.start.eps = self.e_closure(work)
        self.queue = []
        self.queue.append(self.start)
        self.all_e_closure = []
        self.all_e_closure.append(self.start.eps)
        Dstate = self.pop_anlable()
        while Dstate != 0:
            Dstate.label = True
            for i in alph:
                work = Dstate.eps
                w_move = self.move(work, i)
                w_e_closure = self.e_closure(w_move)
                if (w_e_closure not in self.all_e_closure):
                    self.all_e_closure.append(w_e_closure)
                    w_point = MPoint('s' + str(j))
                    j += 1
                    w_point.eps = w_e_closure
                    Dstate.transitions[i] = w_point
                    self.queue.append(w_point)
                else:
                    for k in self.queue:
                        if k.eps == w_e_closure:
                            Dstate.transitions[i] = k
            Dstate = self.pop_anlable()
        for i in self.queue:
            if nfa_end in i.eps:
                i.is_term = True
            if nfa in i.eps:
                i.name = "Start"

    def e_closure(self, T_stack):
        e_closur = []
        while T_stack != []:
            work = T_stack.pop()
            if work not in e_closur:
                e_closur.append(work)
            for i in work.eps:
                if i not in e_closur:
                    e_closur.append(i)
                    T_stack.append(i)
        return e_closur

    def move(self, T_stack, ch):
        mov = []
        for i in T_stack:
            work = i.transitions.get(ch)
            if (work != None) and (work not in mov):
                mov.append(work)
        return mov

    def pop_anlable(self):
        for i in self.queue:
            if i.label == False:
                return i
        return 0

    def print_DFA(self, f):
        string = 'digraph G {\n'
        for i in self.queue:
            string += str(i.name) + '\n'
            for j in i.transitions:
                string += str(i.name) + '->' + str(i.transitions[j].name) + '[label=' + str(j) + ']' + '\n'
        return string + '}'


class MFA:
    def __init__(self, dfa, alph):
        def tr_al_init(point, alph):
            for i in alph:
                point.transitions[i] = []

        # step 1
        queue_b = {}
        for i in dfa.queue:
            work = queue_b.get(i.name)
            if work == None:
                work = MPoint(i.name)
                tr_al_init(work, alph)
                queue_b[i.name] = work
            for j in i.transitions:
                work1 = queue_b.get(i.transitions[j].name)
                if work1 == None:
                    work1 = MPoint(i.transitions[j].name)
                    tr_al_init(work1, alph)
                    queue_b[i.transitions[j].name] = work1
                work1.transitions[j].append(work)
        # step 2
        dost = {}
        for i in dfa.queue:
            dost[i.name] = False

        def DFS(dost, point):
            dost[point.name] = True
            point.label = True
            for i in point.transitions:
                if point.transitions[i].label == False:
                    DFS(dost, point.transitions[i])

        DFS(dost, dfa.start)
        # step 3
        marked = {}
        for i in dfa.queue:
            work = {}
            for j in dfa.queue:
                work[j.name] = False
            marked[i.name] = work
        q = []
        for i in dfa.queue:
            for j in dfa.queue:
                if (marked[i.name][j.name] != True) and i.is_term != j.is_term:
                    marked[i.name][j.name] = True
                    marked[j.name][i.name] = True
                    q.append([i, j])
        # step 4
        while len(q) != 0:
            work = q.pop(0)
            work1 = work[1].name
            work = work[0].name
            for k in alph:
                for i in queue_b[work].transitions[k]:
                    for j in queue_b[work1].transitions[k]:
                        if marked[i.name][j.name] != True:
                            marked[i.name][j.name] = True
                            marked[j.name][i.name] = True
                            q.append([i, j])
        # step 5
        components = {}
        for i in dfa.queue:
            components[i] = -1
        for i in dfa.queue:
            if marked[dfa.queue[0].name][i.name] != True:
                components[i] = 0
        componentsCount = 0
        for i in dfa.queue:
            if dost[i.name] != True:
                continue
            if components[i] == -1:
                componentsCount += 1
                components[i] = componentsCount
                for j in dfa.queue:
                    if marked[i.name][j.name] != True:
                        components[j] = componentsCount
        # step 6
        self.queue = {}
        i = 0
        while i <= componentsCount:
            self.queue['s' + str(i)] = MPoint('s' + str(i))
            i += 1
        for i in components:
            work = self.queue.get('s' + str(components[i]))
            if i == dfa.start:
                self.start = work
            if i.is_term == True:
                work.is_term = True
            for j in i.transitions:
                work1 = self.queue.get('s' + str(components[i.transitions[j]]))
                if work.transitions.get(j) == None:
                    work.transitions[j] = work1

    def tr_al_init(self, point, alph):
        for i in alph:
            point.transitions[i] = []
        return point

    def print_MFA(self, f):
        string = 'digraph G {\n'
        for i in self.queue:
            string += str(i) + '\n'
            for j in self.queue[i].transitions:
                string += i + '->' + str(self.queue[i].transitions[j].name) + '[label=' + str(j) + ']' + '\n'
        return string + '}'
