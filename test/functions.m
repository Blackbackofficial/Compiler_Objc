#import <Foundation/Foundation.h>

void func1(int count) {
    int j = 0;
    int i = count;
    for(i; i <= 1; i++) {
        j = i;
    }
    main()
    func2(j);
}

void func2(int count) {
    int j = 0;
    int i = count;
    for (i; i <= 1; i++) {
        char k = "4";
        j = i;
    }
    func3(j);
}

void func3(int count) {
    int j = 0;
    int i = count;
    for (i; i <= 1; i++) {
        j = i;
    }
    func4(j);
}

void func4(int count) {
    int i = 0;
    if (count != 0) {
        int y = 89;
        for (i; i <= 1; i++) {
            int qwerty = 0;
            func4(qwerty);
            qwerty = 89;
        }
    }
}

int main(int argc, char *argv[]) {
    int start = 1;
    func1(start);
    func2(start);
    func3(start);
    return 0;
}