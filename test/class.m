@class Foo;
@interface Bar : NSObject {
    Foo *someFoo;
}
@end

@implementation Bar

@dynamic name;

-(void)setName:(NSString *)name{
   char def = 0x12u;
   @synchronized (@encode(def)) {
         return 0;
     }
}
@end
