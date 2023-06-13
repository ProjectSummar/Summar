import { Ionicons } from "@expo/vector-icons";
import { useGetBookmarks } from "@src/api/bookmark";
import BookmarkCard from "@src/components/bookmark-card";
import { useSearchStore } from "@src/stores/search-store";
import { Link, Stack, useNavigation } from "expo-router";
import { FlatList, TextInput, View } from "react-native";
import { useMemo } from "react";

const Index = () => {
    const drawer = useNavigation("/main") as any;

    const { data: bookmarks, isLoading } = useGetBookmarks();

    const query = useSearchStore((state) => state.query);
    const setQuery = useSearchStore((state) => state.setQuery);

    const filteredBookmarks = useMemo(() => {
        if (query === "") return bookmarks;

        return bookmarks?.filter((bookmark) =>
            bookmark.title.toLowerCase().match(query)
        );
    }, [bookmarks, query]);

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
            <View
                style={{
                    flexDirection: "row",
                    alignItems: "center",
                    margin: 10,
                    paddingHorizontal: 10,
                    backgroundColor: "#fafafa",
                    borderRadius: 10,
                    shadowColor: "black",
                    shadowOffset: { width: 0, height: 2 },
                    shadowOpacity: 0.1,
                    shadowRadius: 5,
                }}
            >
                <Ionicons name="search" size={25} color="gray" />
                <TextInput
                    style={{
                        padding: 15,
                        flex: 1,
                    }}
                    placeholder="Search"
                    placeholderTextColor="gray"
                    autoCapitalize="none"
                    value={query}
                    onChangeText={setQuery}
                    clearButtonMode="always"
                />
            </View>
            <FlatList
                data={filteredBookmarks}
                renderItem={({ item }) => <BookmarkCard bookmark={item} />}
                keyExtractor={(item) => item.id}
                refreshing={!bookmarks || isLoading}
            />
        </>
    );
};

export default Index;
