import { Bookmark } from "@src/types";
import { Link } from "expo-router";
import { Pressable, StyleSheet, Text, View } from "react-native";
import BookmarkContextMenu from "./bookmark-context-menu";

const BookmarkCard = ({ bookmark }: { bookmark: Bookmark }) => {
    return (
        <View style={styles.card}>
            <Link href={`/main/bookmark/${bookmark.id}`} asChild>
                <Pressable style={styles.pressable}>
                    <View>
                        <Text style={styles.title}>{bookmark.title}</Text>
                        <Text style={styles.url}>{bookmark.url}</Text>
                    </View>
                </Pressable>
            </Link>
            <BookmarkContextMenu />
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
