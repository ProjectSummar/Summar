import { useCreateBookmark } from "@src/api/bookmark";
import { Stack, useNavigation, useRouter } from "expo-router";
import { StatusBar } from "expo-status-bar";
import { Controller, useForm } from "react-hook-form";
import {
    Button,
    Keyboard,
    StyleSheet,
    Text,
    TextInput,
    TouchableWithoutFeedback,
    View,
} from "react-native";

type CreateBookmarkInput = {
    url: string;
};

const Create = () => {
    const { mutate: createBookmark } = useCreateBookmark();

    const {
        control,
        handleSubmit,
        reset: resetForm,
    } = useForm<CreateBookmarkInput>();

    const router = useRouter();

    const nav = useNavigation();

    const createBookmarkOnSubmit = handleSubmit((input) => {
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
                    <Text style={styles.inputLabel}>Bookmark Url</Text>
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
                        name="url"
                    />
                </View>
            </TouchableWithoutFeedback>
        </>
    );
};

const BackButton = ({ backFn }: { backFn: any }) => {
    return <Button title="Back" onPress={backFn} />;
};

const styles = StyleSheet.create({
    container: {},
    inputLabel: {},
    inputField: {},
});

export default Create;
