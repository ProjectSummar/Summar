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
import { Link, useRouter } from "expo-router";
import { useForm, Controller } from "react-hook-form";
import { useLogin } from "@src/api/auth";

type LoginFormInput = {
    email: string;
    password: string;
};

const Login = () => {
    const { mutate: login, isLoading } = useLogin();

    const {
        control,
        handleSubmit,
        reset: resetForm,
    } = useForm<LoginFormInput>();

    const router = useRouter();

    const loginOnSubmit = handleSubmit((input) => {
        console.log("login:", input);
        login(input, {
            onSuccess: () => {
                router.push("/bookmarks");
            },
            onSettled: () => {
                resetForm();
            },
        });
    });

    return (
        <TouchableWithoutFeedback onPress={Keyboard.dismiss} accessible={false}>
            <SafeAreaView style={styles.container}>
                <View style={styles.loginContainer}>
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
                            styles.loginButton,
                        ]}
                        onPress={loginOnSubmit}
                        disabled={isLoading}
                    >
                        <Text style={styles.loginButtonText}>Log In</Text>
                    </Pressable>
                    <Link style={styles.signupLink} href="/auth/signup">
                        Don't have an account? Sign up here.
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
    loginContainer: {
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
