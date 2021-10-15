// #import <Foundation/Foundation.h> //для NSObject и строк NSString
// #include <asl.h>
//
// @interface LoginViewController ()
// // 116 - 124
//
// @property(readwrite,strong) int* usernameView;
// @property(nonatomic,weak) int* dummyNameView;
// // @property(readonly,retain) int* testikView1;
//
// - (void)displayLocalVariable;
// + (void)terminalLocalVariable;
//
// @end

int main(char *argv) // или же просто main()
{
    int i = 9;
    int j = 8 + i;
    int value = 1;
    value *= value++;
    value /= value - value--;
    value %= value % --value;
    value += value - --value;
    value >>= value + ++value;
    int integerVar = 100u; // DECIMAL_LITERAL=129
    return 0;
}