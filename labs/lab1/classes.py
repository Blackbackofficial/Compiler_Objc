from queue import SimpleQueue
from typing import List, Set
from graphviz import Digraph


class Tree(object):
    def __init__(self, data, left=None, right=None):
        self.left = left
        self.right = right
        self.data = data
        self.idx = None

    def __str__(self):
        return self.data

    def set_idx(self, count=1):
        if not (self.left or self.right):
            self.idx = count
            return count + 1
        if self.left:
            count = self.left.set_idx(count)
        if self.right:
            count = self.right.set_idx(count)
        return count

    def nullable(self) -> bool:
        if self.data == '|':
            return self.left.nullable() or self.right.nullable()
        if self.data == '*':
            return True
        if self.data == 'cat':
            return self.left.nullable() and self.right.nullable()
        return False

    def firstpos(self):
        if self.data == '|':
            return self.left.firstpos() + self.right.firstpos()
        if self.data == '*':
            return self.left.firstpos()
        if self.data == 'cat':
            if self.left.nullable():
                return self.left.firstpos() + self.right.firstpos()
            return self.left.firstpos()
        return [self.idx]

    def lastpos(self):
        if self.data == '|':
            return self.left.lastpos() + self.right.lastpos()
        if self.data == '*':
            return self.left.lastpos()
        if self.data == 'cat':
            if self.right.nullable():
                return self.left.lastpos() + self.right.lastpos()
            return self.right.lastpos()
        return [self.idx]

    def followpos(self, pos):
        res = set()
        if self.data == '*' and pos in self.lastpos():
            res.update(set(self.firstpos()))
        if self.data == 'cat' and pos in self.left.lastpos():
            res.update(set(self.right.firstpos()))
        if self.left:
            res.update(self.left.followpos(pos))
        if self.right:
            res.update(self.right.followpos(pos))
        return res


class DKA(object):
    def __init__(self, states, lang, moves, q0, finish_states):
        self.states = states
        self.moves = moves
        self.lang = lang
        self.q0 = q0
        self.finish_states = finish_states

        self.moves_table = [[[] for x in states] for y in states]
        for move in moves:
            i = self.states.index(move['in'])
            j = self.states.index(move['out'])
            self.moves_table[i][j].append(move['move'])

    def is_terminal(self, state):
        return state in self.finish_states

    def draw_fsm(self):
        f = Digraph('finite_state_machine', filename='fsm.gv')
        for state in self.states:
            moves = [move for move in self.moves if move['in'] == state or move['out'] == state]
            if not moves:
                continue
            if state in self.finish_states:
                f.node(f'q{self.states.index(state)}', shape='doublecircle')
            else:
                f.node(f'q{self.states.index(state)}')
        for move in self.moves:
            f.edge(f"q{self.states.index(move['in'])}", f"q{self.states.index(move['out'])}", move['move'])
        f.view()

    def eval(self, string):
        current_state = self.states.index(self.q0)
        for i in range(len(string)):
            if string[i] not in self.lang:
                print("Wrong symbol detected")
                return False
            found = False
            for moves in self.moves_table[current_state]:
                if string[i] in moves:
                    current_state = self.moves_table[current_state].index(moves)
                    found = True
            if not found:
                print("String can not be evaluated")
                return False
        if self.states[current_state] not in self.finish_states:
            print("String finish in non-finish state")
            return False
        print("String evaluated successfully")
        return True

    def minimize(self):
        states = [set()] + self.states
        moves = [x for x in self.moves]
        moves_table = [[self.lang for x in states]] + [[self.lang] + moves for moves in self.moves_table]
        count = len(states)
        back_moves = [[moves_table[x][y] for x in range(count)] for y in range(count)]

        q = SimpleQueue()
        marked = [[False for x in range(count)] for y in range(count)]
        for i in range(count):
            for j in range(count):
                if not marked[i][j] and self.is_terminal(states[i]) != self.is_terminal(states[j]):
                    marked[i][j] = True
                    marked[j][i] = True
                    q.put({i, j})

        while not q.empty():
            u, v = q.get()
            for c in self.lang:
                for r in [i for i in range(count) if back_moves[u][i] == c]:
                    for s in [i for i in range(count) if back_moves[v][i] == c]:
                        if not marked[r][s]:
                            marked[r][s] = True
                            marked[s][r] = True
                            q.put({r, s})

        component = [-1 for i in range(count)]
        for i in range(count):
            if not marked[0][i]:
                component[i] = 0

        components_count = 0
        for i in range(count):
            if component[i] == -1:
                components_count += 1
                component[i] = components_count
                for j in range(i+1, count):
                    if not marked[i][j]:
                        component[j] = components_count
        unique_components = set([x for x in component if x != 0])
        to_delete = []
        for i in unique_components:
            # JOIN STATES AND MOVES
            component_states = [self.states[j-1] for j in range(1, len(component)) if component[j] == i]
            to_delete.extend(component_states[1:])
            for move in moves:
                if move['in'] in component_states:
                    move['in'] = component_states[0]
                if move['out'] in component_states:
                    move['out'] = component_states[0]
        states = [i for i in self.states if i not in to_delete]
        finish_states = [i for i in self.finish_states if i not in to_delete]

        new_moves = []
        for move in moves:
            if move not in new_moves:
                new_moves.append(move)

        self.states = states
        self.moves = new_moves
        self.finish_states = finish_states

        self.moves_table = [[[] for x in states] for y in states]
        for move in new_moves:
            i = self.states.index(move['in'])
            j = self.states.index(move['out'])
            self.moves_table[i][j].append(move['move'])


class NKA(object):
    def __init__(self, states, lang, moves, q0, finish_states):
        self.states = states
        self.moves = moves
        self.lang = lang
        self.q0 = q0
        self.finish_states = finish_states

        self.moves_table = [[[] for x in states] for y in states]
        for move in moves:
            i = self.states.index(move['in'])
            j = self.states.index(move['out'])
            self.moves_table[i][j].append(move['move'])

    def is_terminal(self, state):
        return state in self.finish_states

    def draw_fsm(self):
        f = Digraph('finite_state_machine', filename='fsm.gv')
        for state in self.states:
            moves = [move for move in self.moves if move['in'] == state or move['out'] == state]
            if not moves:
                continue
            if state in self.finish_states:
                f.node(f'q{self.states.index(state)}', shape='doublecircle')
            else:
                f.node(f'q{self.states.index(state)}')
        for move in self.moves:
            f.edge(f"q{self.states.index(move['in'])}", f"q{self.states.index(move['out'])}", move['move'])
        f.view()

    def eps_closure(self, state):
        closure = list([state])

        q = SimpleQueue()
        q.put(state)
        while not q.empty():
            s = q.get()
            for move in self.moves:
                if move['in'] == s and move['move'] == 'eps':
                    if move['out'] not in closure:
                        q.put(move['out'])
                        closure.append(move['out'])

        return closure

    def reachable(self, state, t):
        reachable = list()

        for move in self.moves:
            if move['in'] == state and move['move'] == t:
                if move['out'] not in reachable:
                    reachable.append(move['out'])

        return reachable


regexp_ops_sym = ["|", "*"]


def check_brackets(regexp: str) -> bool:
    group_deep = 0
    for i in range(len(regexp)):
        if regexp[i] == '(':
            group_deep += 1
        elif regexp[i] == ')':
            if group_deep == 0:
                return False
            group_deep -= 1

    return group_deep == 0


def remove_brackets(regexp: str) -> str:
    if regexp[0] == '(' and regexp[-1] == ')':
        if check_brackets(regexp[1:-1]):
            return remove_brackets(regexp[1:-1])
    return regexp


def find_outter_op(regexp) -> int:
    level = 0
    ops_idx = {"or": 0, "cat": 0, "star": 0}
    prev = None
    for i in range(len(regexp)):
        if level == 0:
            if regexp[i] == '*':
                ops_idx['star'] = i
            elif regexp[i] == '|':
                ops_idx['or'] = i
            elif prev and prev != '|' and prev != '(':
                ops_idx['cat'] = i
        if regexp[i] == '(':
            level += 1
        elif regexp[i] == ')':
            level -= 1
        prev = regexp[i]

    if ops_idx['or']:
        return ops_idx['or']
    if ops_idx['cat']:
        return ops_idx['cat']
    return ops_idx['star']


def get_moves(regexp: str) -> Set[str]:
    filter_list = ['(', ')', '*', '|', '#']
    return set([i for i in regexp if i not in filter_list])


def get_syntax_tree(regexp: str) -> Tree:
    regexp = remove_brackets(regexp)
    if len(regexp) == 1:
        return Tree(regexp)
    op_idx = find_outter_op(regexp)
    if regexp[op_idx] == '|':
        left = get_syntax_tree(regexp[0:op_idx])
        right = get_syntax_tree(regexp[op_idx+1::])
        return Tree('|', left, right)
    elif regexp[op_idx] == '*':
        left = get_syntax_tree(regexp[0:op_idx])
        return Tree('*', left)
    else:
        left = get_syntax_tree(regexp[0:op_idx])
        right = get_syntax_tree(regexp[op_idx::])
        return Tree('cat', left, right)


def get_unmarked(states: List[dict]) -> int:
    for i in range(len(states)):
        if not states[i]['marked']:
            return i
    return -1


def fill_mark_map(tree: Tree, arr: List[str]):
    if not tree.left and not tree.right:
        arr[tree.idx-1] = tree.data
        return
    if tree.left:
        fill_mark_map(tree.left, arr)
    if tree.right:
        fill_mark_map(tree.right, arr)


def get_dka(regexp: str) -> DKA:
    tree = get_syntax_tree(regexp)
    lang = get_moves(regexp)
    count = tree.set_idx()
    followpos_table = [tree.followpos(i) for i in range(1, count)]
    state_map = ['' for i in range(1, count)]
    fill_mark_map(tree, state_map)
    q0 = {'state': set(tree.firstpos()), 'marked': False}
    d_state = [q0]
    d_tran = []
    unmarked_idx = get_unmarked(d_state)

    while unmarked_idx > -1:
        d_state[unmarked_idx]['marked'] = True

        for move in lang:
            union = set()

            for i in d_state[unmarked_idx]['state']:
                if state_map[i - 1] == move:
                    union.update(followpos_table[i-1])

            found = False
            for state in d_state:
                if state['state'] == union:
                    found = True

            if not found:
                d_state.append({'state': union, 'marked': False})
            if union:
                d_tran.append({'in': d_state[unmarked_idx]['state'], 'out': union, 'move': move})

        unmarked_idx = get_unmarked(d_state)

    finish_states = []
    for state in d_state:
        found = False
        for pos in state['state']:
            if state_map[pos-1] == '#':
                found = True
        if found:
            finish_states.append(state)

    return DKA([state['state'] for state in d_state], lang, d_tran, q0['state'],
               [state['state'] for state in finish_states])


def get_nka(tree: Tree, last_q_num):
    if tree.data not in regexp_ops_sym + ['cat']:
        start_q = last_q_num
        fin_q = last_q_num + 1
        states = [start_q, fin_q]
        move = {'in': start_q, 'out': fin_q, 'move': tree.data}
        return NKA(states, [tree.data], [move], start_q, [fin_q]), last_q_num + 2
    elif tree.data == '|':
        left_nka, last_q_num = get_nka(tree.left, last_q_num)
        right_nka, last_q_num = get_nka(tree.right, last_q_num)
        start_q = last_q_num
        fin_q = last_q_num + 1
        states = left_nka.states + right_nka.states + [start_q, fin_q]
        lang = set(left_nka.lang + right_nka.lang)
        moves = [{'in': start_q, 'out': left_nka.q0, 'move': 'eps'},
                 {'in': start_q, 'out': right_nka.q0, 'move': 'eps'},
                 {'in': left_nka.finish_states[0], 'out': fin_q, 'move': 'eps'},
                 {'in': right_nka.finish_states[0], 'out': fin_q, 'move': 'eps'}] + left_nka.moves + right_nka.moves
        return NKA(states, lang, moves, start_q, [fin_q]), last_q_num + 2
    elif tree.data == 'cat':
        left_nka, last_q_num = get_nka(tree.left, last_q_num)
        right_nka, last_q_num = get_nka(tree.right, last_q_num)
        states = left_nka.states + right_nka.states[1:]
        lang = set(left_nka.lang)
        lang.union(set(right_nka.lang))
        moves = left_nka.moves + right_nka.moves
        for move in moves:
            if move['in'] == right_nka.states[0]:
                move['in'] = left_nka.states[-1]
            if move['out'] == right_nka.states[0]:
                move['out'] = left_nka.states[-1]
        return NKA(states, lang, moves, left_nka.q0, [right_nka.finish_states[0]]), last_q_num
    elif tree.data == '*':
        left_nka, last_q_num = get_nka(tree.left, last_q_num)
        start_q = last_q_num
        fin_q = last_q_num + 1
        states = [start_q] + left_nka.states + [fin_q]
        lang = left_nka.lang
        moves = [{'in': start_q, 'out': left_nka.q0, 'move': 'eps'},
                 {'in': left_nka.finish_states[0], 'out': fin_q, 'move': 'eps'},
                 {'in': left_nka.finish_states[0], 'out': left_nka.q0, 'move': 'eps'},
                 {'in': start_q, 'out': fin_q, 'move': 'eps'}] + left_nka.moves
        return NKA(states, lang, moves, start_q, [fin_q]), last_q_num + 2


def nka_to_dka(nka: NKA) -> DKA:
    q0 = {'state': set(nka.eps_closure(nka.q0)), 'marked': False}
    d_state = [q0]
    d_tran = []
    unmarked_idx = get_unmarked(d_state)

    while unmarked_idx > -1:
        d_state[unmarked_idx]['marked'] = True

        for sym in nka.lang:
            union = set()

            for i in d_state[unmarked_idx]['state']:
                reachable = nka.reachable(i, sym)
                for j in reachable:
                    union.update(nka.eps_closure(j))

            found = False
            for state in d_state:
                if state['state'] == union:
                    found = True

            if not found:
                d_state.append({'state': union, 'marked': False})
            if union:
                d_tran.append({'in': d_state[unmarked_idx]['state'], 'out': union, 'move': sym})

        unmarked_idx = get_unmarked(d_state)

    old_finish_states = []
    for state in nka.states:
        reachable = nka.reachable(state, '#')
        if reachable:
            old_finish_states.append(state)

    finish_states = []
    for state in d_state:
        for i in state['state']:
            if i in old_finish_states and state not in finish_states:
                finish_states.append(state)

    return DKA([state['state'] for state in d_state], nka.lang, d_tran, q0['state'],
               [state['state'] for state in finish_states])
