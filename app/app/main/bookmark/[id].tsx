import { Ionicons } from "@expo/vector-icons";
import { useGetBookmark } from "@src/api/bookmark";
import { Bookmark } from "@src/types";
import { Stack, useLocalSearchParams } from "expo-router";
import { Text } from "react-native";

const BookmarkPage = () => {
    const { id } = useLocalSearchParams();

    if (!id || typeof id !== "string") return <Text>Loading...</Text>;

    const { data: bookmark } = useGetBookmark(id);

    if (!bookmark) return <Text>Loading...</Text>;

    return (
        <>
            <Stack.Screen
                options={{
                    title: bookmark.title,
                    headerRight: () => <SummariseButton bookmark={bookmark} />,
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

const SummariseButton = ({ bookmark }: { bookmark: Bookmark }) => {
    return (
        <Ionicons
            name={bookmark.summary.length === 0 ? "flash-outline" : "flash"}
            size={20}
        />
    );
};

export default BookmarkPage;
