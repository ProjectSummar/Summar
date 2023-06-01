import { Pressable, StyleSheet, Text, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { StatusBar } from "expo-status-bar";
import { Link } from "expo-router";

const Login = () => {
    return (
        <>
            <StatusBar style="dark" />
            <SafeAreaView style={styles.container}>
                <View style={styles.loginContainer}>
                    <Text style={styles.loginHeader}>Log In</Text>
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
                            styles.loginButton,
                        ]}
                    >
                        <Text style={styles.loginButtonText}>Log In</Text>
                    </Pressable>
                    <Link style={styles.signupLink} href="/auth/signup">
                        Click here to Sign Up
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
        marginTop: 200,
    },
    loginContainer: {
        width: "80%",
        marginHorizontal: "auto",
    },
    loginHeader: {
        fontWeight: "bold",
        textAlign: "center",
        fontSize: 30,
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
    loginButton: {
        marginVertical: 20,
        borderRadius: 10,
    },
    loginButtonText: {
        textAlign: "center",
        color: "white",
        padding: 10,
        fontSize: 20,
        fontWeight: "bold",
    },
    signupLink: {
        textAlign: "center",
        textDecorationLine: "underline",
    },
});

export default Login;
