TEST = "(a|b)*abb"
TEST_ALPHA = ['a', 'b']
INPUT_TEST = "aaabbbbaabbabd"


class MPoint:  # для подсчета состояний
    def __init__(self, name):
        self.name = name
        self.transitions = {}
        self.eps = []
        self.label = False
        self.is_term = False


class MBlock:  # блок состояний
    def __init__(self, start, end):
        self.start = start
        self.end = end


class NFA:
    def __init__(self, token):  # в token передаем обратную польскую
        self.nfa_stack = []
        self.stack = []
        self.tokens = token
        i = 1
        for t in self.tokens:
            if t.name == 'CHAR':
                i = self.char_handler(i, t)
            elif t.name == 'STAR':
                i = self.star_handler(i)
            elif t.name == 'PLUS':
                i = self.plus_handler(i)
            elif t.name == "ALT":
                i = self.alt_handler(i)
            else:
                self.concat_handler()
        self.nfa_stack[0].end.is_term = True
        self.nfa_stack[0].start.name = "Start"

    def char_handler(self, i, t):
        work1 = MPoint('s' + str(i))
        i += 1  # для обозначения номера состояния (s1, s2, s3,...)
        work2 = MPoint('s' + str(i))
        i += 1
        work1.transitions[t.value] = work2  # {S1: b -> S2}
        block = MBlock(work1, work2)
        self.nfa_stack.append(block)
        self.stack.append(work1)
        self.stack.append(work2)
        return i

    def star_handler(self, i):
        work_block = self.nfa_stack.pop()
        P1 = MPoint('s' + str(i))
        i += 1
        P2 = MPoint('s' + str(i))
        i += 1
        work_block.end.eps.append(work_block.start)  # Добавляем переход по eps из конца в начало
        work_block.end.eps.append(P2)
        P1.eps.append(work_block.start)
        P1.eps.append(P2)
        work_block = MBlock(P1, P2)
        self.nfa_stack.append(work_block)
        self.stack.append(P1)
        self.stack.append(P2)
        return i

    def plus_handler(self, i):
        work_block = self.nfa_stack.pop()
        P1 = MPoint('s' + str(i))
        i += 1
        P2 = MPoint('s' + str(i))
        i += 1
        work_block.end.eps.append(work_block.start)
        work_block.end.eps.append(P2)
        P1.eps.append(work_block.start)
        work_block = MBlock(P1, P2)
        self.nfa_stack.append(work_block)
        self.stack.append(P1)
        self.stack.append(P2)
        return i

    def alt_handler(self, i):  # ИЛИ
        P1 = MPoint('s' + str(i))
        i += 1
        P2 = MPoint('s' + str(i))
        i += 1
        work_block1 = self.nfa_stack.pop()
        work_block2 = self.nfa_stack.pop()
        P1.eps.append(work_block1.start)
        P1.eps.append(work_block2.start)
        work_block1.end.eps.append(P2)
        work_block2.end.eps.append(P2)
        work_block = MBlock(P1, P2)
        self.nfa_stack.append(work_block)
        self.stack.append(P1)
        self.stack.append(P2)
        return i

    def concat_handler(self):
        work_block2 = self.nfa_stack.pop()
        work_block1 = self.nfa_stack.pop()
        work_block1.end.eps.append(work_block2.start)
        work_block1 = MBlock(work_block1.start, work_block2.end)
        self.nfa_stack.append(work_block1)

    def print_NFA(self):
        string = 'digraph G {\n'
        for i in self.stack:
            string += str(i.name) + '\n'
            for j in i.eps:
                string += str(i.name) + '->' + str(j.name) + '\n'
            for j in i.transitions:
                string += str(i.name) + '->' + str(i.transitions[j].name) + '[label=' + str(j) + ']' + '\n'
        return string + '}'


class DFA:
    def __init__(self, nfa_start, alpha, nfa_end):
        j = 1
        self.start = MPoint('s'+str(j))
        j += 1
        work = [nfa_start]
        self.start.eps = self.e_closure(work)
        self.queue = []
        self.queue.append(self.start)
        self.all_e_closure = []
        self.all_e_closure.append(self.start.eps)
        d_state = self.pop_label()  # изначально в dstates(множестве состояний) содержится начальное состояние
        while d_state != 0:  # пока имеется не помеченное состояние
            d_state.label = True
            for i in alpha:
                work = d_state.eps
                w_move = self.move(work, i)
                w_e_closure = self.e_closure(w_move)
                if w_e_closure not in self.all_e_closure:
                    self.all_e_closure.append(w_e_closure)
                    w_point = MPoint('s' + str(j))
                    j += 1
                    w_point.eps = w_e_closure
                    d_state.transitions[i] = w_point
                    self.queue.append(w_point)
                else:
                    for k in self.queue:
                        if k.eps == w_e_closure:
                            d_state.transitions[i] = k
            d_state = self.pop_label()
        for i in self.queue:
            if nfa_end in i.eps:
                i.is_term = True
            if nfa_start in i.eps:
                i.name = "Start"

    @staticmethod
    def e_closure(T_stack):
        e_closure = []
        while T_stack:  # пока стек не пуст
            work = T_stack.pop()  # снять со стека верхний элемент
            if work not in e_closure:
                e_closure.append(work)
            for i in work.eps:
                if i not in e_closure:
                    e_closure.append(i)
                    T_stack.append(i)  # помещаем обратно в стек
        return e_closure

    @staticmethod
    def move(T_stack, ch):
        mov = []
        for i in T_stack:
            work = i.transitions.get(ch)
            if (work is not None) and (work not in mov):
                mov.append(work)
        return mov

    def pop_label(self):
        for i in self.queue:
            if not i.label:
                return i
        return 0

    def print_DFA(self):
        string = 'digraph G {\n'
        for i in self.queue:
            string += str(i.name) + '\n'
            for j in i.transitions:
                string += str(i.name) + '->' + str(i.transitions[j].name) + '[label=' + str(j) + ']' + '\n'
        return string + '}'


class MFA:
    def __init__(self, dfa, alpha):
        def tr_al_init(point, alpha):
            for r in alpha:
                point.transitions[r] = list()

        # step 1
        queue_b = {}
        for i in dfa.queue:
            work = queue_b.get(i.name)
            if work is None:
                work = MPoint(i.name)
                tr_al_init(work, alpha)
                queue_b[i.name] = work
            for j in i.transitions:
                work1 = queue_b.get(i.transitions[j].name)
                if work1 is None:
                    work1 = MPoint(i.transitions[j].name)
                    tr_al_init(work1, alpha)
                    queue_b[i.transitions[j].name] = work1
                work1.transitions[j].append(work)
        # step 2
        dost = {}
        for i in dfa.queue:
            dost[i.name] = False

        def DFS(d, point):
            d[point.name] = True
            point.label = True
            for el in point.transitions:
                if not point.transitions[el].label:
                    DFS(d, point.transitions[el])

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
                if not marked[i.name][j.name] and i.is_term != j.is_term:
                    marked[i.name][j.name] = True
                    marked[j.name][i.name] = True
                    q.append([i, j])
        # step 4
        while len(q) != 0:
            work = q.pop(0)
            work1 = work[1].name
            work = work[0].name
            for k in alpha:
                for i in queue_b[work].transitions[k]:
                    for j in queue_b[work1].transitions[k]:
                        if not marked[i.name][j.name]:
                            marked[i.name][j.name] = True
                            marked[j.name][i.name] = True
                            q.append([i, j])
        # step 5
        components = {}
        for i in dfa.queue:
            components[i] = -1
        for i in dfa.queue:
            if not marked[dfa.queue[0].name][i.name]:
                components[i] = 0
        componentsCount = 0
        for i in dfa.queue:
            if not dost[i.name]:
                continue
            if components[i] == -1:
                componentsCount += 1
                components[i] = componentsCount
                for j in dfa.queue:
                    if not marked[i.name][j.name]:
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
            if i.is_term:
                work.is_term = True
            for j in i.transitions:
                work1 = self.queue.get('s' + str(components[i.transitions[j]]))
                if work.transitions.get(j) is None:
                    work.transitions[j] = work1

    def print_MFA(self):
        string = 'digraph G {\n'
        for i in self.queue:
            string += str(i) + '\n'
            for j in self.queue[i].transitions:
                string += i + '->' + str(self.queue[i].transitions[j].name) + '[label=' + str(j) + ']' + '\n'
        return string + '}'
