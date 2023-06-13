import { useGetBookmark } from "@src/api/bookmark";
import { Stack, useLocalSearchParams } from "expo-router";
import { useState } from "react";
import { ActivityIndicator, ScrollView, Text, View } from "react-native";
import WebView from "react-native-webview";
import { Share } from "react-native";
import { useToast } from "@src/contexts/toast-context";
import { useSummariseBookmark } from "@src/api/bookmark";
import { ContextMenu, ContextMenuOption } from "@src/components/context-menu";

const BookmarkPage = () => {
    const { id } = useLocalSearchParams();

    const { data: bookmark, isLoading } = useGetBookmark(id as string);

    const [summaryView, setSummaryView] = useState(false);

    const { errorToast, successToast } = useToast();

    const { mutate: summariseBookmark } = useSummariseBookmark();

    if (!bookmark || isLoading) return <Loading />;

    const displaySummary = summaryView && bookmark.summary.length !== 0;

    const summariseBookmarkOnPress = () => {
        if (bookmark.summary.length !== 0) {
            return;
        }

        summariseBookmark(
            { id: bookmark.id },
            {
                onSuccess: () =>
                    successToast("Summarised bookmark successfully"),
                onError: () => errorToast("Error summarising bookmark"),
            },
        );
    };

    const toggleSummaryView = () => {
        if (bookmark.summary.length === 0) {
            errorToast("No summary available");
            return;
        }
        setSummaryView((summaryView) => !summaryView);
    };

    const shareBookmark = async () => {
        return await Share.share({ url: bookmark.url });
    };

    return (
        <>
            <Stack.Screen
                options={{
                    title: bookmark.title,
                    headerRight: () => (
                        <ContextMenu>
                            <ContextMenuOption
                                text="Summarise Bookmark"
                                onSelect={summariseBookmarkOnPress}
                                icon="flash-outline"
                            />
                            <ContextMenuOption
                                text="Toggle Summary View"
                                onSelect={toggleSummaryView}
                                icon="book-outline"
                            />
                            <ContextMenuOption
                                text="Share Bookmark"
                                onSelect={shareBookmark}
                                icon="share-outline"
                            />
                        </ContextMenu>
                    ),
                }}
            />
            {displaySummary
                ? (
                    <ScrollView style={{ padding: 20 }}>
                        <Text style={{ fontSize: 20 }}>{bookmark.summary}</Text>
                    </ScrollView>
                )
                : (
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
