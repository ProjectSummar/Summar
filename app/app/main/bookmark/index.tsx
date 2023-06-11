import { Ionicons } from "@expo/vector-icons";
import { useGetBookmarks } from "@src/api/bookmark";
import BookmarkCard from "@src/components/bookmark-card";
import { Link, Stack } from "expo-router";
import { FlatList } from "react-native";

const Index = () => {
    const { data: bookmarks, isLoading } = useGetBookmarks();

    return (
        <>
            <Stack.Screen
                options={{ title: "Bookmarks", headerRight: CreateButton }}
            />
            <FlatList
                data={bookmarks}
                renderItem={({ item }) => <BookmarkCard bookmark={item} />}
                keyExtractor={(item) => item.id}
                refreshing={!bookmarks || isLoading}
            />
        </>
    );
};

const CreateButton = () => {
    return (
        <Link href="/main/bookmark/create" asChild>
            <Ionicons name="create-outline" size={30} />
        </Link>
    );
};

export default Index;
