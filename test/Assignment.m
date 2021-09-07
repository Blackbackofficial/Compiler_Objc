#import <Foundation/Foundation.h>

int main() {
   int a = 21;
   int c ;

   c =  a;
   NSLog(@"Line 1 - =  Operator Example, Value of c = %d\n", c );

   c +=  a;
   NSLog(@"Line 2 - += Operator Example, Value of c = %d\n", c );

   c -=  a;
   NSLog(@"Line 3 - -= Operator Example, Value of c = %d\n", c );

   c *=  a;
   NSLog(@"Line 4 - *= Operator Example, Value of c = %d\n", c );

   c /=  a;
   NSLog(@"Line 5 - /= Operator Example, Value of c = %d\n", c );

   c  = 200;
   c %=  a;
   NSLog(@"Line 6 - %= Operator Example, Value of c = %d\n", c );

   c <<=  2;
   NSLog(@"Line 7 - <<= Operator Example, Value of c = %d\n", c );

   c >>=  2;
   NSLog(@"Line 8 - >>= Operator Example, Value of c = %d\n", c );

   c &=  2;
   NSLog(@"Line 9 - &= Operator Example, Value of c = %d\n", c );

   c ^=  2;
   NSLog(@"Line 10 - ^= Operator Example, Value of c = %d\n", c );

   c |=  2;
   NSLog(@"Line 11 - |= Operator Example, Value of c = %d\n", c );

}