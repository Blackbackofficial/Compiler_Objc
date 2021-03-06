#import <Foundation/Foundation.h>

@interface Box : NSObject {
   double length;    // Length of a box
   double breadth;   // Breadth of a box
   double height;    // Height of a box
}

@property(nonatomic, readwrite) double height;  // Property

-(byref double) volume;
+(bycopy NSString) print;

@end

@implementation Box

@synthesize height;

    -(id)init {
       int i = 0;
       int j = 0;
       self = [super init];
       length = 1.0;
       if(length == 1.0){
          continue;
       }
       breadth = 1.0;
       return self;
    }

    -(double) volume {
       double i = length*breadth*height;
       return i;
    }

    +(NSString) print {
       double string = 234;
       return string;
    }
@end

int main(int argc, char *argv[]) {
   SEL sel = @selector(lowercaseString);
   auto a = [NSString new];
   NSAutoreleasePool * pool = [[NSAutoreleasePool alloc] init];
   Box *box1 = [[Box alloc]init];    // Create box1 object of type Box
   Box *box2 = [[Box alloc]init];    // Create box2 object of type Box

   double volume = 0.0;             // Store the volume of a box here

   // box 1 specification
   box1.height = 5.0;

   // box 2 specification
   box2.height = 10.0;

   // volume of box 1
   volume = [box1 volume];
   NSLog(@"i = %f", volume);

   // volume of box 2
   volume = [box2 volume];
   NSLog(@"i = %f", volume);

   [pool drain];
   return 0;
}
int qwerty = 89;

@protocol TestProtocol
-(void) test;
-(void) testWithName;
@end

@interface Worker : NSObject
{
    @private struct Data data;
}
@end

@interface Worker2 : NSObject
{
    @package struct Data1 data1;
}
@end

@interface Worker3 : NSObject
{
    @protected struct Data2 data2;
}
@end

@interface Worker4 : NSObject
{
    @private
    struct Data2 data2;
}
@end

@protocol myProtocol
- (void)myMandatoryMethod;
@optional
- (void)myOptionalMethod;
@end