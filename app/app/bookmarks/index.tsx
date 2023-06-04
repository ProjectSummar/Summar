import { StyleSheet, Text, TextInput, View, Image } from "react-native";
import React from "react";

const Index = () => {
    const cards = [
        {
            title: "Title 1",
            description: "Description 1",
            imageSource:
                "https://images.unsplash.com/flagged/photo-1562503542-2a1e6f03b16b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=774&q=80",
        },
        {
            title: "Title 2",
            description: "Description 2",
            imageSource:
                "https://images.unsplash.com/flagged/photo-1562503542-2a1e6f03b16b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=774&q=80",
        },
    ];

    return (
        <View>
            {cards.map((card, index) => (
                <View style={styles.card} key={index}>
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
            ))}
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

export default Index;
