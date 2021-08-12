#import "Complex.h"
@interface Complex (CategorizedComplex)
    - (Complex *)mul: (Complex *)other;
    - (Complex *)div: (Complex *)other;
@end

//файл "CategorizedComplex.m"
#import "CategorizedComplex.h"

@implementation Complex (CategorizedComplex)
    - (Complex *)mul: (Complex *)other
{
    return [Complex complexWithRe: _re * other->_re - _im * other->_im andIm: _re * other->_im + _im * other->_re];
}
    - (Complex *)div: (Complex *)other
{
    double retRe, retIm, denominator;
    denominator = other->_re * other->_re + other->_im * other->_im;
    if (!denominator) {
        return nil;
    }
    retRe = (_re * other->_re + _im * other->_im) / denominator;
    retIm = (_im * other->_re - _re * other->_im) / denominator;
    return [Complex complexWithRe: retRe andIm: retIm];
}
@end