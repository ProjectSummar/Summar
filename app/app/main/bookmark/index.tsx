import { Ionicons } from "@expo/vector-icons";
import { useGetBookmarks } from "@src/api/bookmark";
import BookmarkCard from "@src/components/bookmark-card";
import { Link, Stack, useNavigation } from "expo-router";
import { FlatList } from "react-native";

const Index = () => {
    const drawer = useNavigation("/main") as any;

    const { data: bookmarks, isLoading } = useGetBookmarks();

    return (
        <>
            <Stack.Screen
                options={{
                    title: "Bookmarks",
                    headerRight: () => (
                        <Link href="/main/bookmark/create" asChild>
                            <Ionicons name="create-outline" size={30} />
                        </Link>
                    ),
                    headerLeft: () => (
                        <Ionicons
                            name="menu-sharp"
                            size={30}
                            onPress={drawer.toggleDrawer}
                        />
                    ),
                }}
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

export default Index;
