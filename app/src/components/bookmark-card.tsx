import { Bookmark } from "@src/types";
import { Link } from "expo-router";
import { Pressable, StyleSheet, Text, View } from "react-native";

const BookmarkCard = ({ bookmark }: { bookmark: Bookmark }) => {
    return (
        <Link href={`/main/bookmark/${bookmark.id}`} asChild>
            <Pressable>
                <View style={styles.card}>
                    <Text style={styles.title}>{bookmark.title}</Text>
                    <Text style={styles.url}>{bookmark.url}</Text>
                </View>
            </Pressable>
        </Link>
    );
};

const styles = StyleSheet.create({
    card: {
        backgroundColor: "#FFFFFF",
        padding: 15,
        marginVertical: 1,
        shadowColor: "black",
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 5,
    },
    title: {
        fontSize: 20,
    },
    url: {},
});

export default BookmarkCard;
