import { Share } from "react-native";
import { Bookmark } from "@src/types";
import { useDeleteBookmark } from "@src/api/bookmark";
import { useRouter } from "expo-router";
import { useToast } from "@src/contexts/toast-context";
import { ContextMenu, ContextMenuOption } from "./context-menu";

const BookmarkCardContextMenu = ({ bookmark }: { bookmark: Bookmark }) => {
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
    );
};

export default BookmarkCardContextMenu;
