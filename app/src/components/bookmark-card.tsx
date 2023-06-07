import { Image, StyleSheet, Text, View } from "react-native";

const BookmarkCard = ({ card }: { card: any }) => {
    return (
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
    );
};

const styles = StyleSheet.create({
    card: {
        backgroundColor: "#FFFFFF",
        padding: 16,
        marginVertical: 1,
        shadowColor: "#000000",
        shadowOffset: { width: 0, height: 2 },
        shadowOpacity: 0.1,
        shadowRadius: 4,
    },
    cardContent: {
        flexDirection: "row",
    },
    cardText: {
        flex: 1,
        marginRight: 16,
    },
    cardImage: {
        width: 100,
        height: 100,
        borderRadius: 5,
    },
    title: {
        fontSize: 18,
        fontWeight: "bold",
        marginBottom: 8,
    },
    description: {
        fontSize: 16,
        color: "#808080",
    },
});

export default BookmarkCard;
