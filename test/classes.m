#import <UIKit/UIKit.h>

@interface SVIndefiniteAnimatedView : UIView

@property (nonatomic, assign) CGFloat strokeThickness;
@property (nonatomic, assign) CGFloat radius;
@property (nonatomic, strong) UIColor *strokeColor;

+ (void)setDefaultStyle:(SVProgressHUDStyle)style;                      // default is SVProgressHUDStyleLight
+ (void)setDefaultMaskType:(SVProgressHUDMaskType)maskType;             // default is SVProgressHUDMaskTypeNone
+ (void)setDefaultAnimationType:(SVProgressHUDAnimationType)type;       // default is SVProgressHUDAnimationTypeFlat
+ (void)setContainerView:(nullable UIView*)containerView;               // default is window level
+ (void)setMinimumSize:(CGSize)minimumSize;                             // default is CGSizeZero, can be used to avoid resizing for a larger message
+ (void)setRingThickness:(CGFloat)ringThickness;                        // default is 2 pt
+ (void)setRingRadius:(CGFloat)radius;                                  // default is 18 pt
+ (void)setRingNoTextRadius:(CGFloat)radius;                            // default is 24 pt
+ (void)setCornerRadius:(CGFloat)cornerRadius;                          // default is 14 pt
+ (void)setBorderColor:(nonnull UIColor*)color;                         // default is nil
+ (void)setBorderWidth:(CGFloat)width;                                  // default is 0
+ (void)setFont:(nonnull UIFont*)font;                                  // default is [UIFont preferredFontForTextStyle:UIFontTextStyleSubheadline]
+ (void)setForegroundColor:(nonnull UIColor*)color;                     // default is [UIColor blackColor], only used for SVProgressHUDStyleCustom
+ (void)setForegroundImageColor:(nullable UIColor*)color;               // default is nil == foregroundColor, only used for SVProgressHUDStyleCustom
+ (void)setBackgroundColor:(nonnull UIColor*)color;                     // default is [UIColor whiteColor], only used for SVProgressHUDStyleCustom
+ (void)setHudViewCustomBlurEffect:(nullable UIBlurEffect*)blurEffect;  // default is nil, only used for SVProgressHUDStyleCustom, can be combined with backgroundColor
+ (void)setBackgroundLayerColor:(nonnull UIColor*)color;                // default is [UIColor colorWithWhite:0 alpha:0.5], only used for SVProgressHUDMaskTypeCustom
+ (void)setImageViewSize:(CGSize)size;                                  // default is 28x28 pt
+ (void)setShouldTintImages:(BOOL)shouldTintImages;                     // default is YES
+ (void)setInfoImage:(nonnull UIImage*)image;                           // default is the bundled info image provided by Freepik
+ (void)setSuccessImage:(nonnull UIImage*)image;                        // default is the bundled success image provided by Freepik
+ (void)setErrorImage:(nonnull UIImage*)image;                          // default is the bundled error image provided by Freepik
+ (void)setViewForExtension:(nonnull UIView*)view;                      // default is nil, only used if #define SV_APP_EXTENSIONS is set
+ (void)setGraceTimeInterval:(NSTimeInterval)interval;                  // default is 0 seconds
+ (void)setMinimumDismissTimeInterval:(NSTimeInterval)interval;         // default is 5.0 seconds
+ (void)setMaximumDismissTimeInterval:(NSTimeInterval)interval;         // default is infinite
+ (void)setFadeInAnimationDuration:(NSTimeInterval)duration;            // default is 0.15 seconds
+ (void)setFadeOutAnimationDuration:(NSTimeInterval)duration;           // default is 0.15 seconds
+ (void)setMaxSupportedWindowLevel:(UIWindowLevel)windowLevel;          // default is UIWindowLevelNormal
+ (void)setHapticsEnabled:(BOOL)hapticsEnabled;						    // default is NO
+ (void)setMotionEffectEnabled:(BOOL)motionEffectEnabled;               // default is YES
@end