#import <Foundation/Foundation.h>

@interface Box : NSObject {
   double length;    // Length of a box
   double breadth;   // Breadth of a box
   double height;    // Height of a box
}

@property(nonatomic, readwrite) double height;  // Property

-(double) volume;
+(NSString) print;

@end


@implementation Box

@synthesize height;

    -(id)init {
       int i = 0;
       int j = 0;
       self = [super init];
       length = 1.0;
       breadth = 1.0;
       return self;
    }

    -(double) volume {
       double i = length*breadth*height;
       return i;
    }

    +(NSString) print {
       double string = 234;
       NSLog (@"%f", string);
    }

@end

int main(int argc, char *argv[]) {
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
   NSLog(@"Volume of Box1 : %f", volume);

   // volume of box 2
   volume = [box2 volume];
   NSLog(@"Volume of Box2 : %f", volume);

   [pool drain];
   return 0;
}
int qwerty = 89;