#import "Foundation/Foundation.h"

    @interface example //: NSObject
    {
        @public NSString* name;
    }

    @end

    @implementation example @end

    int main()
    {
        example* me;
        if (example) {
            return me;
        }
        me->name = @"World";
    }