import { useGetBookmark } from "@src/api/bookmark";
import BookmarkPageContextMenu from "@src/components/bookmark-page-context-menu";
import { Stack, useLocalSearchParams } from "expo-router";
import { useState } from "react";
import { ActivityIndicator, ScrollView, Text, View } from "react-native";
import WebView from "react-native-webview";

const BookmarkPage = () => {
    const { id } = useLocalSearchParams();

    const { data: bookmark, isLoading } = useGetBookmark(id as string);

    const [summaryView, setSummaryView] = useState(false);

    if (!bookmark || isLoading) return <Loading />;

    const displaySummary = summaryView && bookmark.summary.length !== 0;

    return (
        <>
            <Stack.Screen
                options={{
                    title: bookmark.title,
                    headerRight: () => (
                        <BookmarkPageContextMenu
                            bookmark={bookmark}
                            setSummaryView={setSummaryView}
                        />
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
