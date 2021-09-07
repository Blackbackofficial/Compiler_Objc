#import <Foundation/Foundation.h>

int git(int t) {
    static int i = 5;

    register int i = 5;

    i += 10;
    return i;
}

int main() {
   NSLog(@"Storage size for float : %d \n", sizeof(float));
   NSLog(@"Storage size for float : %d \n", sizeof(short));
   NSLog(@"Storage size for float : %d \n", sizeof(signed));
   NSLog(@"Storage size for float : %d \n", sizeof(unsigned));
   NSArray *objectsToShare = @[@"F", @"R"]; // 79
   volatile int a = 0;

   int n = (*frac).numerator;     // these two expressions
   int z = frac->numerator;       // are equivalent

   return 0;
}

typedef enum ShapeType {
    kCircle,
    kRectangle,
    kOblateSpheroid
};
enum ACCOUNT_TYPE {
        SAVINGS,
        LOAN,
        FIXED_DEPOSIT
};

struct SavingsData {
        NSMutableString* accountName;
        NSMutableString* accountType;
        NSMutableString* balance;
        NSMutableString* currency;
};

struct LoanData {
        NSMutableString* accountName;
        NSMutableString* balance;
};

extern const struct CommonResourceAttributes {
    __weak NSString *typeKey;
    __unsafe_unretained NSString *identifierKey;
} CommonResourceAttributes;

struct FixedDepositData {
        NSMutableString* accountName;
        NSMutableString* depositAmount;
        NSMutableString* roi;
};

union AccountData {
        struct SavingsData *savingsData;
        struct LoanData *loanData;
        struct FixedDepositData *fixedDepositData;
} AccountData;

typedef NS_ENUM(NSUInteger, ShapeType) {
    kCircle,
    kRectangle,
    kOblateSpheroid
};

typedef NS_OPTIONS(NSUInteger, Fix) {
    sCircle,
    sRectangle,
    sOblateSpheroid
};