import { useGetBookmark } from "@src/api/bookmark";
import { Stack, useLocalSearchParams } from "expo-router";
import { Text } from "react-native";

const BookmarkPage = () => {
    const { id } = useLocalSearchParams();

    if (!id || typeof id !== "string") return <Text>Loading...</Text>;

    const { data: bookmark } = useGetBookmark(id);

    if (!bookmark) return <Text>Loading...</Text>;

    return (
        <>
            <Stack.Screen options={{ title: "" }} />
            <Text>{id}</Text>
            {Object.entries(bookmark).map(([key, value]) => (
                <Text key={key}>
                    {key}: {value}
                </Text>
            ))}
        </>
    );
};

export default BookmarkPage;
