import { useUpdateBookmarkTitle } from "@src/api/bookmark";
import { useToast } from "@src/contexts/toast-context";
import {
    Stack,
    useLocalSearchParams,
    useNavigation,
    useRouter,
} from "expo-router";
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

type UpdateBookmarkTitleInput = {
    title: string;
};

const Update = () => {
    const { errorToast, successToast } = useToast();

    const { id } = useLocalSearchParams();

    const { mutate: updateBookmarkTitle, isLoading } = useUpdateBookmarkTitle();

    const {
        control,
        handleSubmit,
        reset: resetForm,
    } = useForm<UpdateBookmarkTitleInput>();

    const router = useRouter();

    const nav = useNavigation();

    const updateBookmarkTitleOnSubmit = handleSubmit((input) => {
        console.log("updateBookmarkTitleOnSubmit", input);

        updateBookmarkTitle(
            { id: id as string, title: input.title },
            {
                onSuccess: () => {
                    successToast("Bookmark updated successfully");
                    router.push("/main/bookmark");
                },
                onError: () => errorToast("Error updating bookmark"),
                onSettled: () => {
                    resetForm();
                },
            }
        );
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
                        <Pressable
                            style={({ pressed }) => [
                                {
                                    backgroundColor: pressed ? "gray" : "black",
                                },
                                styles.button,
                            ]}
                            onPress={updateBookmarkTitleOnSubmit}
                            disabled={isLoading}
                        >
                            <Text style={styles.buttonText}>
                                Update Bookmark
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

export default Update;
