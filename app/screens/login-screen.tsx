import { useState } from "react";
import {
    StyleSheet,
    View,
    Dimensions,
    Text,
    TextInput,
    Pressable,
} from "react-native";
import Svg, { Image, Ellipse, ClipPath } from "react-native-svg";
import Animated, {
    useSharedValue,
    useAnimatedStyle,
    interpolate,
    withTiming,
    withDelay,
    withSequence,
    withSpring,
} from "react-native-reanimated";

const { width, height } = Dimensions.get("window"); // get dimensions of phone

function LoginScreen() {
    const imagePosition = useSharedValue(1);
    const formButtonScale = useSharedValue(1);
    const [isRegistering, setIsRegistering] = useState(false);
    const imageAnimatedStyle = useAnimatedStyle(() => {
        const interpolation = interpolate(
            imagePosition.value,
            [0, 1],
            [-height / 2, 0]
        );
        return {
            transform: [
                { translateY: withTiming(interpolation, { duration: 1000 }) },
            ],
        };
    });

    const buttonsAnimatedStyle = useAnimatedStyle(() => {
        const interpolation = interpolate(
            imagePosition.value,
            [0, 1],
            [250, 0]
        );
        return {
            opacity: withTiming(imagePosition.value, { duration: 500 }),
            transform: [
                { translateY: withTiming(interpolation, { duration: 1000 }) },
            ],
        };
    });

    const closeButtonContainerStyle = useAnimatedStyle(() => {
        const interpolation = interpolate(
            imagePosition.value,
            [0, 1],
            [180, 360]
        );
        return {
            opacity: withTiming(imagePosition.value === 1 ? 0 : 1, {
                duration: 800,
            }),
            transform: [
                {
                    rotate: withTiming(interpolation + "deg", {
                        duration: 1000,
                    }),
                },
            ],
        };
    });

    const formAnimatedStyle = useAnimatedStyle(() => {
        return {
            opacity:
                imagePosition.value === 0
                    ? withDelay(400, withTiming(1, { duration: 800 }))
                    : withTiming(0, { duration: 300 }),
        };
    });
    const formButtonAnimatedStyle = useAnimatedStyle(() => {
        return {
            transform: [{ scale: formButtonScale.value }],
        };
    });
    const loginHandler = () => {
        imagePosition.value = 0;
        if (isRegistering) {
            setIsRegistering(false);
        }
    };

    const registerHandler = () => {
        imagePosition.value = 0;
        if (!isRegistering) {
            setIsRegistering(true);
        }
    };

    return (
        <Animated.View style={styles.container}>
            <Animated.View
                style={[StyleSheet.absoluteFill, imageAnimatedStyle]}
            >
                <Svg height={height + 100} width={width}>
                    <ClipPath id="clipPathId">
                        <Ellipse cx={width / 2} rx={height} ry={height + 100} />
                    </ClipPath>
                    <Image
                        href={require("../assets/login-background.jpg")}
                        width={width + 100}
                        height={height + 100}
                        preserveAspectRatio="xMidYMid slice" // element to be centered both horizontally and vertically within its container while maintaining its aspect ratio.
                        clipPath="url(#clipPathId)"
                    />
                </Svg>
                <Animated.View
                    style={[
                        styles.closeButtonContainer,
                        closeButtonContainerStyle,
                    ]}
                >
                    <Text onPress={() => (imagePosition.value = 1)}>X</Text>
                </Animated.View>
            </Animated.View>
            <View style={styles.bottomContainer}>
                <Animated.View style={buttonsAnimatedStyle}>
                    <Pressable style={styles.button} onPress={loginHandler}>
                        <Text style={styles.buttonText}>LOG IN</Text>
                    </Pressable>
                </Animated.View>
                <Animated.View style={buttonsAnimatedStyle}>
                    <Pressable style={styles.button} onPress={registerHandler}>
                        <Text style={styles.buttonText}>SIGN UP</Text>
                    </Pressable>
                </Animated.View>
                <Animated.View
                    style={[styles.formInputContainer, formAnimatedStyle]}
                >
                    {isRegistering && (
                        <TextInput
                            placeholder="Email"
                            placeholderTextColor="black"
                            style={styles.textInput}
                        ></TextInput>
                    )}
                    <TextInput
                        placeholder="Username"
                        placeholderTextColor="black"
                        style={styles.textInput}
                    ></TextInput>
                    <TextInput
                        placeholder="Password"
                        placeholderTextColor="black"
                        style={styles.textInput}
                    ></TextInput>
                    <Animated.View
                        style={[styles.formButton, formButtonAnimatedStyle]}
                    >
                        <Pressable
                            onPress={() =>
                                (formButtonScale.value = withSequence(
                                    withSpring(1.5),
                                    withSpring(1)
                                ))
                            }
                        >
                            <Text style={styles.buttonText}>
                                {isRegistering ? "SIGN UP" : "LOG IN"}
                            </Text>
                        </Pressable>
                    </Animated.View>
                </Animated.View>
            </View>
        </Animated.View>
    );
}

const styles = StyleSheet.create({
    container: {
        flex: 1,
        justifyContent: "flex-end",
    },
    button: {
        backgroundColor: "rgba(1,20,255,0.8)",
        height: 55,
        alignItems: "center",
        justifyContent: "center",
        borderRadius: 35,
        marginHorizontal: 20,
        marginVertical: 10,
        borderWidth: 1,
        borderColor: "white",
    },
    buttonText: {
        fontSize: 20,
        fontWeight: "600",
        color: "white",
        letterSpacing: 0.5,
    },
    bottomContainer: {
        justifyContent: "center",
        height: height / 3,
    },
    textInput: {
        height: 50,
        borderWidth: 1,
        borderColor: "rgba(0,0,0,0.2)",
        marginHorizontal: 20,
        marginVertical: 10,
        borderRadius: 25,
        paddingLeft: 10,
    },
    formButton: {
        backgroundColor: "rgba(1,20,255,0.8)",
        height: 55,
        alignItems: "center",
        justifyContent: "center",
        borderRadius: 35,
        marginHorizontal: 20,
        marginVertical: 10,
        borderWidth: 1,
        borderColor: "white",
        shadowColor: "#000",
        shadowOffset: {
            width: 0,
            height: 4,
        },
        shadowOpacity: 0.25,
        shadowRadius: 3.84,
        elevation: 5,
    },
    formInputContainer: {
        marginBottom: 70,
        ...StyleSheet.absoluteFillObject,
        zIndex: -1,
        justifyContent: "center",
        backgroundColor: "white",
    },
    closeButtonContainer: {
        height: 40,
        width: 40,
        justifyContent: "center",
        alignSelf: "center",
        shadowColor: "#000",
        shadowOffset: {
            width: 0,
            height: 5,
        },
        shadowOpacity: 0.34,
        shadowRadius: 6.27,
        elevation: 1,
        backgroundColor: "white",
        alignItems: "center",
        borderRadius: 20,
        top: -20,
    },
});

export default LoginScreen;
