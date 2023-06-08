import { useCreateBookmark } from "@src/api/bookmark";
import { Stack, useNavigation, useRouter } from "expo-router";
import { StatusBar } from "expo-status-bar";
import { Controller, useForm } from "react-hook-form";
import {
    Button,
    Keyboard,
    Pressable,
    StyleSheet,
    Text,
    TextInput,
    TouchableWithoutFeedback,
    View,
} from "react-native";

type CreateBookmarkInput = {
    url: string;
    title: string;
};

const Create = () => {
    const { mutate: createBookmark, isLoading } = useCreateBookmark();

    const {
        control,
        handleSubmit,
        reset: resetForm,
    } = useForm<CreateBookmarkInput>();

    const router = useRouter();

    const nav = useNavigation();

    const createBookmarkOnSubmit = handleSubmit((input) => {
        console.log("createBookmarkOnSubmit", input);

        createBookmark(input, {
            onSuccess: () => {
                router.push("/main/bookmark");
            },
            onSettled: () => {
                resetForm();
            },
        });
    });

    return (
        <>
            <StatusBar style="light" />
            <Stack.Screen
                options={{
                    headerLeft: () => <BackButton backFn={nav.goBack} />,
                }}
            />
            <TouchableWithoutFeedback
                onPress={Keyboard.dismiss}
                accessible={false}
            >
                <View style={styles.container}>
                    <View style={styles.inputForm}>
                        <Text>Bookmark Title</Text>
                        <Controller
                            control={control}
                            rules={{ required: true }}
                            render={({
                                field: { onChange, onBlur, value },
                            }) => (
                                <TextInput
                                    style={styles.inputField}
                                    placeholder="Enter bookmark title here"
                                    placeholderTextColor="gray"
                                    autoCapitalize="none"
                                    onBlur={onBlur}
                                    onChangeText={onChange}
                                    value={value}
                                />
                            )}
                            name="title"
                        />
                        <Text>Bookmark URL</Text>
                        <Controller
                            control={control}
                            rules={{ required: true }}
                            render={({
                                field: { onChange, onBlur, value },
                            }) => (
                                <TextInput
                                    style={styles.inputField}
                                    placeholder="Enter bookmark url here"
                                    placeholderTextColor="gray"
                                    autoCapitalize="none"
                                    onBlur={onBlur}
                                    onChangeText={onChange}
                                    value={value}
                                />
                            )}
                            name="url"
                        />
                        <Pressable
                            style={({ pressed }) => [
                                {
                                    backgroundColor: pressed ? "gray" : "black",
                                },
                                styles.button,
                            ]}
                            onPress={createBookmarkOnSubmit}
                            disabled={isLoading}
                        >
                            <Text style={styles.buttonText}>
                                Create Bookmark
                            </Text>
                        </Pressable>
                    </View>
                </View>
            </TouchableWithoutFeedback>
        </>
    );
};

const BackButton = ({ backFn }: { backFn: any }) => {
    return <Button title="Back" onPress={backFn} />;
};

const styles = StyleSheet.create({
    container: {
        alignItems: "center",
        marginTop: 30,
    },
    inputForm: {
        width: "70%",
        gap: 10,
    },
    inputField: {
        backgroundColor: "white",
        padding: 15,
        borderRadius: 10,
    },
    button: {
        marginTop: 10,
        borderRadius: 10,
    },
    buttonText: {
        width: "100%",
        textAlign: "center",
        color: "white",
        padding: 10,
        fontWeight: "bold",
    },
});

export default Create;
