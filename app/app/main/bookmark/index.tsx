import BookmarkCard from "@src/components/bookmark-card";
import { Link, Stack } from "expo-router";
import { Button, View } from "react-native";

const Index = () => {
    const cards = [
        {
            id: 1,
            title: "Title 1",
            description: "Description 1",
            imageSource:
                "https://images.unsplash.com/flagged/photo-1562503542-2a1e6f03b16b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=774&q=80",
        },
        {
            id: 2,
            title: "Title 2",
            description: "Description 2",
            imageSource:
                "https://images.unsplash.com/flagged/photo-1562503542-2a1e6f03b16b?ixlib=rb-4.0.3&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D&auto=format&fit=crop&w=774&q=80",
        },
    ];

    return (
        <>
            <Stack.Screen
                options={{ title: "Bookmarks", headerRight: CreateButton }}
            />
            <View>
                {cards.map((card, index) => (
                    <BookmarkCard
                        key={index}
                        card={card}
                    />
                ))}
            </View>
        </>
    );
};

const CreateButton = () => {
    return (
        <Link href="/main/bookmark/create" asChild>
            <Button title="+" />
        </Link>
    );
};

export default Index;
