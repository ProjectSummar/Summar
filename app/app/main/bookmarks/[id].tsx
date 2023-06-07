import { Stack, useLocalSearchParams } from "expo-router";
import { Text } from "react-native";

const BookmarkPage = () => {
    const { id } = useLocalSearchParams();

    return (
        <>
            <Stack.Screen
                options={{
                    title: `Bookmark ${id}`,
                }}
            />
            <Text>{id}</Text>
        </>
    );
};

export default BookmarkPage;
