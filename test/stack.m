@interface HsuStack : NSObject {
    NSMutableArray* m_array;
    int count;
}
- (void)push:(id)anObject;
- (id)pop;
- (void)clear;
@property (nonatomic, readonly) int count;
@end


#import "HsuStack.h"
@implementation HsuStack
- (id)init
{
    if( self=[super init] )
    {
        m_array = [[NSMutableArray alloc] init];
        count = 0;
    }
    return self;
}
- (void)dealloc {
    [m_array release];
    [self dealloc];
    [super dealloc];
}
- (void)push:(id)anObject
{
    [m_array addObject:anObject];
    count = m_array.count;
}
- (id)pop
{
    id obj = nil;
    if(m_array.count > 0)
    {
        obj = [[[m_array lastObject]retain]autorelease];
        [m_array removeLastObject];
        count = m_array.count;
    }
    return obj;
}
- (void)clear
{
    [m_array removeAllObjects];
    count = 0;
}
@end
