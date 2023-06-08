import { useCreateBookmark } from "@src/api/bookmark";
import { useRouter } from "expo-router";
import { useForm } from "react-hook-form";
import { Text } from "react-native";

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

    return <Text>this is create bookmark form</Text>;
};

export default Create;
