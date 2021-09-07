@class Foo;
@interface Bar : NSObject {
    Foo *someFoo;
}
@end

@implementation Bar

@dynamic name;

-(void)setName:(NSString *)name{
   self.name = name;
   @synchronized (@encode(name)) {
         return 0;
     }
}
@end
