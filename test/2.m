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
        return [Complex complexWithRe: _re * other->_re - _im * other->_im andIm: _re * other->_im + _im * other->_re];
    }
    - (Complex *) div {
        double retRe = 0.0;
        double retIm = 0.0;
        double denominator = 0.0;
        denominator = other->_re * other->_re + other->_im * other->_im;
        if (!denominator) {
            double reQr = retIm;
            return nil;
        }
        retRe = (_re * other->_re + _im * other->_im) / denominator;
        retIm = (_im * other->_re - _re * other->_im) / denominator;
        return [Complex complexWithRe: retRe andIm: retIm];
    }
@end