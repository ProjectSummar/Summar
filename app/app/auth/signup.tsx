import { Pressable, StyleSheet, Text, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { StatusBar } from "expo-status-bar";
import { Link } from "expo-router";

const Signup = () => {
    return (
        <>
            <StatusBar style="dark" />
            <SafeAreaView style={styles.container}>
                <View style={styles.signupContainer}>
                    <Text style={styles.logo}>Summar</Text>
                    <View style={styles.input}>
                        <Text style={styles.inputLabel}>Email</Text>
                        <TextInput
                            style={styles.inputField}
                            placeholder="Enter your email here"
                        />
                    </View>
                    <View style={styles.input}>
                        <Text style={styles.inputLabel}>Password</Text>
                        <TextInput
                            style={styles.inputField}
                            placeholder="Enter your password here"
                        />
                    </View>
                    <Pressable
                        style={({ pressed }) => [
                            {
                                backgroundColor: pressed ? "gray" : "black",
                            },
                            styles.signupButton,
                        ]}
                    >
                        <Text style={styles.signupButtonText}>Sign Up</Text>
                    </Pressable>
                    <Link style={styles.loginLink} href="/auth/login">
                        Click here to Log In
                    </Link>
                </View>
            </SafeAreaView>
        </>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "flex-start",
        marginTop: 150,
    },
    signupContainer: {
        width: "80%",
        marginHorizontal: "auto",
    },
    logo: {
        fontWeight: "bold",
        textAlign: "center",
        fontSize: 50,
        marginBottom: 20,
    },
    input: {
        marginBottom: 10,
    },
    inputLabel: {
        marginBottom: 10,
    },
    inputField: {
        backgroundColor: "white",
        padding: 15,
        borderRadius: 10,
    },
    signupButton: {
        marginVertical: 20,
        borderRadius: 10,
    },
    signupButtonText: {
        textAlign: "center",
        color: "white",
        padding: 10,
        fontSize: 20,
        fontWeight: "bold",
    },
    loginLink: {
        textAlign: "center",
        textDecorationLine: "underline",
    },
});

export default Signup;
