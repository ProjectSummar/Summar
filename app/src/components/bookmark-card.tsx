import { Bookmark } from "@src/types";
import { Link } from "expo-router";
import { Pressable, StyleSheet, Text, View } from "react-native";
import { Share } from "react-native";
import { useDeleteBookmark } from "@src/api/bookmark";
import { useRouter } from "expo-router";
import { useToast } from "@src/contexts/toast-context";
import { ContextMenu, ContextMenuOption } from "./context-menu";

const BookmarkCard = ({ bookmark }: { bookmark: Bookmark }) => {
    const { errorToast, successToast } = useToast();

    const router = useRouter();

    const { mutate: deleteBookmark } = useDeleteBookmark();

    const deleteBookmarkOnSelect = () => {
        console.log("deleting bookmark", bookmark.title);
        deleteBookmark(
            { id: bookmark.id },
            {
                onSuccess: () => successToast("Bookmark deleted successfully"),
                onError: () => errorToast("Error deleting bookmark"),
            },
        );
    };

    const shareBookmark = async () => {
        return await Share.share({ url: bookmark.url });
    };

    return (
        <View style={styles.card}>
            <Link href={`/main/bookmark/${bookmark.id}`} asChild>
                <Pressable style={styles.pressable}>
                    <View>
                        <Text style={styles.title}>{bookmark.title}</Text>
                        <Text style={styles.url} numberOfLines={1}>
                            {bookmark.url}
                        </Text>
                    </View>
                </Pressable>
            </Link>
            <ContextMenu>
                <ContextMenuOption
                    text="Delete Bookmark"
                    onSelect={deleteBookmarkOnSelect}
                    icon="trash"
                />
                <ContextMenuOption
                    text="Update Bookmark"
                    onSelect={() =>
                        router.push({
                            pathname: "/main/bookmark/update",
                            params: { id: bookmark.id },
                        })}
                    icon="md-pencil-sharp"
                />
                <ContextMenuOption
                    text="Share Bookmark"
                    onSelect={shareBookmark}
                    icon="share-outline"
                />
            </ContextMenu>
        </View>
    );
};

const styles = StyleSheet.create({
    card: {
        flexDirection: "row",
        alignItems: "center",
        justifyContent: "space-between",
        backgroundColor: "#FFFFFF",
        padding: 15,
        marginVertical: 1,
        shadowColor: "black",
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 5,
    },
    pressable: {
        flex: 1,
    },
    title: {
        fontSize: 20,
    },
    url: {},
});

export default BookmarkCard;
