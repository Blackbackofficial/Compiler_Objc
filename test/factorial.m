#import <Foundation/Foundation.h>

int main(int argc, char *argv[]) {
    int a = 5;
    int f = 2;
    int i = 1;
    for(i; i <= 5; i++) {
        int d = i * f;
        f = d;
    }
    NSLog(@"%d", f);
}