from Compiler_Objc.labs.lab1.classes import *


# ex for ka: (xy*|ab|(x|a*))(x|y*), (a|b)*abb#
def main():
    regexp = input('Enter regular expression: ') + '#'
    dka = get_dka(regexp)
    choice = 1
    while choice:
        print('\nChoose action from menu below: ')
        print('   1) Evaluate string with dka')
        print('   2) Enter new regexp')
        print('   3) Minimize dka')
        print('   4) Draw dka')
        print('   0) Exit from program')
        try:
            choice = int(input())
        except Exception:
            raise Exception('Wrong command')
        if choice == 1:
            string = input("Enter input string for dka: ")
            dka.eval(string)
        if choice == 2:
            regexp = input('Enter regular expression: ') + '#'
            dka = get_dka(regexp)
        if choice == 3:
            dka.minimize()
        if choice == 4:
            dka.draw_fsm()


if __name__ == '__main__':
    main()
