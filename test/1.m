#import <Foundation/Foundation.h> //для NSObject и строк NSString

int main(int argc, char *argv[]) // или же просто main()
{
    NSAutoreleasePool *pool = [[NSAutoreleasePool alloc] init]; //создаем пул, он автоматически становится текущим
    int retVal;
    int i = 6;
    if (true) {
        int j = 6;
        j++;
        i--;
        NSLog (@"Hello World");
        while (i < 20) {
              int qwerty = i;
              qwerty++;
              if ((i % 2) != 0) {
                NSLog(@"i = %i", i);
              }

              NSLog(@"i = %i", i);
            }
    }
    [pool drain]; //освобождаем пул и все объекты, помещенные в него вызовами autorelease
    return retVal;
}