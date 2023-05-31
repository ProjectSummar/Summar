import { Pressable, StyleSheet, Text, TextInput, View } from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { StatusBar } from "expo-status-bar";

const Signup = () => {
    return (
        <>
            <StatusBar style="dark" />
            <SafeAreaView style={styles.container}>
                <View style={styles.signupContainer}>
                    <Text style={styles.signupHeader}>Sign Up</Text>
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
    signupContainer: {
        width: "80%",
        marginHorizontal: "auto",
    },
    signupHeader: {
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
    signupButton: {
        marginTop: 10,
        borderRadius: 10,
    },
    signupButtonText: {
        textAlign: "center",
        color: "white",
        padding: 10,
        fontSize: 20,
        fontWeight: "bold",
    },
});

export default Signup;
