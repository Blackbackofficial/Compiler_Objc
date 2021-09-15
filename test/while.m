#import <Foundation/Foundation.h> //для NSObject и строк NSString

int main(int argc, char *argv[]) // или же просто main()
{
    int i = 6;
    i = i/i;
    i = i++;
    i = i++;
    i = i++;
    i = i++;
    i = i++;


    while (i < 20) {
        int j = 2;
        while (i >= 3) {
            j++;
            int t = 2;
            while (i == 3 && j != 6) {
                t = t * i;
                t++;
            }
        }
        i++;
    }
    return 0;
}