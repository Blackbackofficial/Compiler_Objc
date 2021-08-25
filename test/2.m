#import "Complex.h"
@interface Complex : NSObject {
    - (Complex *)mul;
    - (Complex *)div;
}
@end

//файл "CategorizedComplex.m"
#import "CategorizedComplex.h"

@implementation Complex
    - (Complex *) mul {
        return [[Complex alloc] init]];
    }
    - (Complex *) div {
        double retRe = 6.0;
        double retIm = 3.0;
        double denominator = 7.0;
        denominator = retRe * retIm;
        if (!denominator) {
            double reQr = retIm;
            return 0;
        }
        retRe = (retRe + retIm) / denominator;
        retIm = (retRe - retIm) / denominator;
        return [Complex complexWithRe: retRe: retIm];
    }
@end