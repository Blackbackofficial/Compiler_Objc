IfStatement() {
    int i = 0;
    int count = 0;
    if (count != 0) {
        int y = 89;
        ForStatement();
    } else {
        i = 7;
    }
//     y = 0;
    DoWhileStatement();
}

ForStatement() {
    int i = 0;
    for (i; i <= 1; i++) {
        int qwerty = 0;
        WhileStatement();
        ForStatement();
        qwerty = 89;
    }
}

WhileStatement() {
    int i = 19;
    while (i < 20) {
        int j = i;
        j++;
//         DoWhileStatement();
        break;
    }
}

// DoWhileStatement() {
//     int i = 600;
//     int foo = 10;
//     do {
//         int razr = foo;
//         razr = 1 + --razr;
//         foo = razr;
//     } while (i / i > 0 && foo % foo >= 1);
// }

int main(int argc, char *argv[]) {
    IfStatement();
    return 0;
}