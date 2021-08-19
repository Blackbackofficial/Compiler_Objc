#import <Foundation/Foundation.h>
//
// void func1(int count)
// {
//     int j = 0;
//     int i;
//     for(i = count; i <= 1; i++) {
//         NSLog(@"func1");
//         j = i;
//     }
//     func2(j);
// }
//
// void func2(int count)
// {
//     int j = 0;
//     int i;
//     for (i = count; i <= 1; i++) {
//         NSLog(@"func2");
//         j = i;
//     }
//     func3(j);
// }
//
// void func3(int count)
// {
//     int j = 0;
//     int i;
//     for (i = count; i <= 1; i++) {
//         NSLog(@"func3");
//         j = i;
//     }
//     func4(j);
// }
//
// void func4(int count)
// {
//     int i;
//     if (count != 0) {
//         for (i = count; i <= 1; i++) {
//             NSLog(@"func4");
//             func4(0);
//         }
//     }
// }

int main(int argc, char *argv[])
{
    int start = 1;
    func1(start);
    func2(start);
    func3(start);
    NSLog(@"End");
    return 0;
}