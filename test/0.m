#import <Foundation/Foundation.h> //для NSObject и строк NSString

int main(int argc, char *argv[]) // или же просто main()
{
    int integerVar = 100; // DECIMAL_LITERAL=129
    float floatVar = 331.79; // FLOATING_POINT_LITERAL=131
    double doubleVar = 8.44e+11;
    char charVar = 'W'; // CHARACTER_LITERAL=126
    BOOL boolVar = YES; // 125
    NSString *my_string = @"Line1 \
                            Line2"; // STRING_LITERAL=127
    char hex = 0x12; // Hex HEX_LITERAL=128
    char octal = 0213; //OCTAL_LITERAL=130

    NSLog (@"Тип переменной int. Значение переменной равно %i", integerVar);
    NSLog (@"Тип переменной float. Значение переменной равно %f", floatVar);
    NSLog (@"Тип переменной double. Значение переменной равно %e", doubleVar);
    NSLog (@"Тип переменной double. Значение переменной равно %g", doubleVar);
    NSLog (@"Тип переменной char. Значение переменной равно %c", charVar);
    NSLog (@"Тип переменной BOOL. Значение переменной равно %c", boolVar);
    NSLog (@"Тип переменной NSString. Строковый литерал %c", my_string);
    NSLog (@"Тип переменной char. Литерал в шестнадцатеричной системе %c", hex);
    NSLog (@"Тип переменной char. Восьмеричный литерал %c", octal);

    return 0;
}