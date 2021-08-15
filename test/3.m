#import "Foundation/Foundation.h"
@class Foo;
@interface Bar : NSObject {
    Foo *someFoo;
}
@end

@interface example : NSObject
{
    @public NSString* name;
}

@end

@implementation example @end

int main(int argc)
{
    example* me;
    if (example) {
        return me;
    }
    me->name = @"World";
}