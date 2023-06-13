import { Ionicons } from "@expo/vector-icons";
import {
    Menu,
    MenuOption,
    MenuOptions,
    MenuTrigger,
} from "react-native-popup-menu";
import { Text } from "react-native";
import { IconName } from "@src/types";
import { ReactNode } from "react";

const ContextMenu = ({ children }: { children: ReactNode }) => {
    return (
        <Menu>
            <MenuTrigger>
                <Ionicons
                    name="ellipsis-vertical"
                    size={20}
                    style={{ padding: 10 }}
                />
            </MenuTrigger>

            <MenuOptions
                customStyles={{ optionsContainer: { borderRadius: 10 } }}
            >
                {children}
            </MenuOptions>
        </Menu>
    );
};

const ContextMenuOption = ({
    text,
    onSelect,
    icon,
}: {
    text: string;
    onSelect: () => any;
    icon: IconName;
}) => {
    return (
        <MenuOption
            onSelect={onSelect}
            customStyles={{
                optionWrapper: {
                    flexDirection: "row",
                    alignItems: "center",
                    justifyContent: "space-between",
                    padding: 10,
                },
            }}
        >
            <Text style={{ fontSize: 15 }}>{text}</Text>
            <Ionicons name={icon} size={20} />
        </MenuOption>
    );
};

export { ContextMenu, ContextMenuOption };
