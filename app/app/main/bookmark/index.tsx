import { useGetBookmarks } from "@src/api/bookmark";
import BookmarkCard from "@src/components/bookmark-card";
import { Link, Stack } from "expo-router";
import { Button, Text, View } from "react-native";

const Index = () => {
    const { data: bookmarks, isLoading } = useGetBookmarks();

    if (!bookmarks || isLoading) return <Text>Loading...</Text>;

    return (
        <>
            <Stack.Screen
                options={{ title: "Bookmarks", headerRight: CreateButton }}
            />
            <View>
                {bookmarks?.map((bookmark, index) => (
                    <BookmarkCard
                        key={index}
                        bookmark={bookmark}
                    />
                ))}
            </View>
        </>
    );
};

const CreateButton = () => {
    return (
        <Link href="/main/bookmark/create" asChild>
            <Button title="+" />
        </Link>
    );
};

export default Index;
