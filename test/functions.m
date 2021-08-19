#import <Foundation/Foundation.h>
// NSLog(@"Excellent!\n");

int func1(int count)
{
    int j = 0;
    for (int i = count; i <= 1; i++) {
        NSLog(@"func1");
        j = i;
    }
    func2(j);
}

int func2(int count)
{
    int j = 0;
    for (int i = count; i <= 1; i++) {
        NSLog(@"func2");
        j = i;
    }
    func3(j);
}

int func3(int count)
{
    int j = 0;
    for (int i = count; i <= 1; i++) {
        NSLog(@"func2");
        j = i;
    }
    func4(j);
}

int func3(int count)
{
    int j = 0;
    for (int i = count; i <= 1; i++) {
        NSLog(@"func2");
        func3(j);
    }
}

int main(int argc, char *argv[])
{
    int start = 1;
    func3(start);
    NSLog(@"%s", "End");
}