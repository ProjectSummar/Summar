import { Ionicons } from "@expo/vector-icons";
import { useGetBookmark, useSummariseBookmark } from "@src/api/bookmark";
import { Stack, useLocalSearchParams } from "expo-router";
import { useState } from "react";
import { ActivityIndicator, ScrollView, Text, View } from "react-native";
import WebView from "react-native-webview";

const BookmarkPage = () => {
    const { id } = useLocalSearchParams();

    const { data: bookmark, isLoading } = useGetBookmark(id as string);

    const { mutate: summariseBookmark } = useSummariseBookmark();

    const [summaryView, setSummaryView] = useState(false);

    if (!bookmark || isLoading) return <Loading />;

    const displaySummary = summaryView && bookmark.summary.length !== 0;

    const summariseBookmarkOnPress = () => {
        if (bookmark.summary.length !== 0) {
            return;
        }

        summariseBookmark({ id: bookmark.id });
    };

    return (
        <>
            <Stack.Screen
                options={{
                    title: bookmark.title,
                    headerRight: () => (
                        <View style={{ flexDirection: "row", gap: 10 }}>
                            <Ionicons
                                name={
                                    bookmark.summary.length === 0
                                        ? "flash-outline"
                                        : "flash"
                                }
                                size={20}
                                onPress={summariseBookmarkOnPress}
                            />
                            <Ionicons
                                name={displaySummary ? "book" : "book-outline"}
                                size={20}
                                onPress={() =>
                                    setSummaryView(
                                        (summaryView) => !summaryView
                                    )
                                }
                            />
                        </View>
                    ),
                }}
            />
            {displaySummary ? (
                <ScrollView style={{ padding: 20 }}>
                    <Text style={{ fontSize: 20 }}>{bookmark.summary}</Text>
                </ScrollView>
            ) : (
                <WebView
                    originWhitelist={["*"]}
                    source={{ uri: bookmark.url }}
                    style={{ flex: 1 }}
                    startInLoadingState={true}
                    renderLoading={() => <Loading />}
                />
            )}
        </>
    );
};

const Loading = () => {
    return (
        <View
            style={{
                position: "absolute",
                height: "100%",
                width: "100%",
                justifyContent: "center",
                alignItems: "center",
                backgroundColor: "white",
            }}
        >
            <ActivityIndicator size="large" />
        </View>
    );
};

export default BookmarkPage;
