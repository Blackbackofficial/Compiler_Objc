#import <Foundation/Foundation.h>

@interface Person : NSObject

@property (strong) Person *bestFriend;

@property id firstName;
@property NSString *lastName;
@property (readonly) int age;

@end

#import "Person.h"

@implementation Person

- (NSString *) description {
    return self.firstName;
}

@end

#import <Foundation/Foundation.h>
#import "Person.h"

int main(int argc, const char * argv[]) {

    @autoreleasepool {

        Person *yong = [[Person alloc] init];
        yong.firstName = @"Yong";
        yong.lastName = @"Bakos";

        NSLog(@"Hello there, %@", yong);
    }
    return 0;
}