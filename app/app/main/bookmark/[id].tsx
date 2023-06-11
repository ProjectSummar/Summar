import { Ionicons } from "@expo/vector-icons";
import { useGetBookmark, useSummariseBookmark } from "@src/api/bookmark";
import { Stack, useLocalSearchParams } from "expo-router";
import { Text } from "react-native";

const BookmarkPage = () => {
    const { id } = useLocalSearchParams();

    const { data: bookmark, isLoading } = useGetBookmark(id as string);

    const { mutate: summariseBookmark } = useSummariseBookmark();

    if (!bookmark || isLoading) return <Text>Loading...</Text>;

    const summariseBookmarkOnPress = () => {
        summariseBookmark({ id: bookmark.id });
    };

    return (
        <>
            <Stack.Screen
                options={{
                    title: bookmark.title,
                    headerRight: () => (
                        <Ionicons
                            name={bookmark.summary.length === 0
                                ? "flash-outline"
                                : "flash"}
                            size={20}
                            onPress={summariseBookmarkOnPress}
                        />
                    ),
                }}
            />
            {Object.entries(bookmark).map(([key, value]) => (
                <Text key={key}>
                    {key}: {value}
                </Text>
            ))}
        </>
    );
};

export default BookmarkPage;
