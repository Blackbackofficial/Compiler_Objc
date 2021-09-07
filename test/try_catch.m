#import <Foundation/Foundation.h>

int main(int argc, char *argv[]) // или же просто main()
{
    NSString *test = @"test";
    unichar a;
    int index = 5;
    int a = 12;
    @try {
       a = 0x34;
    }
    @catch (NSException *exception) {
       NSLog(@"%@", a);
       NSLog(@"Char at index %d cannot be found", index);
       NSLog(@"Max index is: %lu", a);
    }
    @finally {
       NSLog(@"Finally condition");
    }
}