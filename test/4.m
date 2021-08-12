#import <Foundation/Foundation.h>

@interface Person //: NSObject

@property (strong) Person *bestFriend;

@property id firstName;
@property NSString *lastName;
@property (readonly) int age;

@end

#import "Person.h"

@implementation Person

- (NSString *) description {
    return [NSString stringWithFormat:@"%@ %@, who is %d years old.", self.firstName, self.lastName, self.age];
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

        Person *lindsey = [[Person alloc] init];
        lindsey.firstName = @"Lindsey";
        lindsey.lastName = @"Bakos!";
        yong.bestFriend = lindsey;

        NSArray *array = nil;
        NSLog(@"What is nil: %@", nil);
        NSLog(@"What is the 10th element in the array: %@", [array objectAtIndex:9]);

        if (array) NSLog(@"array is not nil");
        else NSLog(@"array is nil");

        NSString *aString = @"Revolution!";
        NSLog(@"Where does aString point to? %p", aString);
        NSLog(@"Where doe nil point to? %p", nil);

        yong.firstName = @[@"Tupac", @"Shakur"];
        NSLog(@"yong.firstName = %@", yong.firstName);
        NSLog(@"There are %d items in yong.firstName", [yong.firstName count]);

        SEL firstNameSelector = @selector(firstName);

        NSLog(@"[yong firstName] = %@", [yong performSelector:firstNameSelector]);
        NSLog(@"[yong firstName] = %@", [yong performSelector:@selector(firstName)]);
        NSLog(@"[yong firstName] = %@", [yong performSelector:NSSelectorFromString(@"firstName")]);


    }
    return 0;
}