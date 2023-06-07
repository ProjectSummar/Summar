import { Link } from "expo-router";
import { Image, Pressable, StyleSheet, Text, View } from "react-native";

const BookmarkCard = ({ card }: { card: any }) => {
    return (
        <Link href={`/main/bookmarks/${card.id}`} asChild>
            <Pressable>
                <View style={styles.card}>
                    <View style={styles.cardContent}>
                        <View style={styles.cardText}>
                            <Text style={styles.title}>{card.title}</Text>
                            <Text style={styles.description}>
                                {card.description}
                            </Text>
                        </View>
                        <Image
                            source={{ uri: card.imageSource }}
                            style={styles.cardImage}
                        />
                    </View>
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
    cardContent: {
        display: "flex",
        flexDirection: "row",
    },
    cardText: {
        flex: 1,
        marginRight: 15,
    },
    cardImage: {
        width: 100,
        height: 100,
        borderRadius: 5,
    },
    title: {
        fontSize: 20,
        fontWeight: "bold",
        marginBottom: 10,
    },
    description: {
        fontSize: 15,
        color: "gray",
    },
});

export default BookmarkCard;
