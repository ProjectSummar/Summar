import {
    Keyboard,
    Pressable,
    StyleSheet,
    Text,
    TextInput,
    TouchableWithoutFeedback,
    View,
} from "react-native";
import { SafeAreaView } from "react-native-safe-area-context";
import { Link } from "expo-router";
import { Controller, useForm } from "react-hook-form";
import { useSignup } from "@src/api/auth";
import { useToast } from "@src/contexts/toast-context";

type SignupFormInput = {
    email: string;
    password: string;
};

const Signup = () => {
    const { errorToast, successToast } = useToast();

    const { mutate: signup, isLoading } = useSignup();

    const {
        control,
        handleSubmit,
        reset: resetForm,
    } = useForm<SignupFormInput>();

    const signupOnSubmit = handleSubmit((input) => {
        console.log("signup", input);
        signup(input, {
            onSuccess: () => successToast("Signed up successfully"),
            onError: () => errorToast("Error signing up"),
            onSettled: () => resetForm(),
        });
    });

    return (
        <TouchableWithoutFeedback onPress={Keyboard.dismiss} accessible={false}>
            <SafeAreaView style={styles.container}>
                <View style={styles.signupContainer}>
                    <Text style={styles.logo}>Summar</Text>
                    <View style={styles.input}>
                        <Text style={styles.inputLabel}>Email</Text>
                        <Controller
                            control={control}
                            rules={{ required: true }}
                            render={({
                                field: { onChange, onBlur, value },
                            }) => (
                                <TextInput
                                    style={styles.inputField}
                                    placeholder="Enter your email here"
                                    placeholderTextColor="gray"
                                    autoCapitalize="none"
                                    onBlur={onBlur}
                                    onChangeText={onChange}
                                    value={value}
                                />
                            )}
                            name="email"
                        />
                    </View>
                    <View style={styles.input}>
                        <Text style={styles.inputLabel}>Password</Text>
                        <Controller
                            control={control}
                            rules={{ required: true }}
                            render={({
                                field: { onChange, onBlur, value },
                            }) => (
                                <TextInput
                                    style={styles.inputField}
                                    placeholder="Enter your password here"
                                    placeholderTextColor="gray"
                                    autoCapitalize="none"
                                    onBlur={onBlur}
                                    onChangeText={onChange}
                                    value={value}
                                    secureTextEntry
                                />
                            )}
                            name="password"
                        />
                    </View>
                    <Pressable
                        style={({ pressed }) => [
                            {
                                backgroundColor: pressed ? "gray" : "black",
                            },
                            styles.signupButton,
                        ]}
                        onPress={signupOnSubmit}
                        disabled={isLoading}
                    >
                        <Text style={styles.signupButtonText}>Sign Up</Text>
                    </Pressable>
                    <Link style={styles.loginLink} href="/auth/login">
                        Already have an account? Log in here.
                    </Link>
                </View>
            </SafeAreaView>
        </TouchableWithoutFeedback>
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
        backgroundColor: "#fafafa",
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
