import {
    createContext,
    ReactNode,
    useContext,
    useEffect,
    useState,
} from "react";
import { Ionicons } from "@expo/vector-icons";
import Animated, { SlideInUp, SlideOutUp } from "react-native-reanimated";
import { ColorValue, Text } from "react-native";
import { IconName } from "@src/types";
import { useSafeAreaInsets } from "react-native-safe-area-context";

type ToastFn = ({
    text,
    icon,
    color,
    bgColor,
}: {
    text: string;
    icon: IconName;
    color: ColorValue;
    bgColor: ColorValue;
}) => void;

type ToastContext = {
    text: string;
    icon: IconName;
    color: ColorValue;
    bgColor: ColorValue;
    show: boolean;
    toast: ToastFn;
};

const toastContext = createContext<ToastContext>({
    text: "",
    icon: "car",
    color: "red",
    bgColor: "red",
    show: false,
    toast: () => {},
});

const useToast = () => {
    const toast = useContext(toastContext).toast;

    const successToast = (text: string) => {
        toast({
            text,
            icon: "checkmark-circle-outline",
            bgColor: "green",
            color: "white",
        });
    };

    const errorToast = (text: string) => {
        toast({
            text,
            icon: "alert-circle-outline",
            bgColor: "red",
            color: "white",
        });
    };

    return { errorToast, successToast };
};

const ToastProvider = ({ children }: { children: ReactNode }) => {
    const [text, setText] = useState("");
    const [icon, setIcon] = useState<IconName>("car");
    const [color, setColor] = useState<ColorValue>("red");
    const [bgColor, setBgColor] = useState<ColorValue>("red");
    const [show, setShow] = useState(false);

    const toast = ({
        text,
        icon,
        color,
        bgColor,
    }: {
        text: string;
        icon: IconName;
        color: ColorValue;
        bgColor: ColorValue;
    }) => {
        setText(text);
        setIcon(icon);
        setColor(color);
        setBgColor(bgColor);
        setShow(true);
    };

    useEffect(() => {
        if (show) {
            const timer = setTimeout(() => setShow(false), 3000);

            return () => clearTimeout(timer);
        }
    }, [show]);

    return (
        <toastContext.Provider
            value={{ text, icon, color, bgColor, show, toast }}
        >
            {children}
            {show && (
                <ToastComponent
                    text={text}
                    icon={icon}
                    color={color}
                    bgColor={bgColor}
                />
            )}
        </toastContext.Provider>
    );
};

const ToastComponent = ({
    icon,
    text,
    bgColor,
    color,
}: {
    icon: IconName;
    text: string;
    bgColor: ColorValue;
    color: ColorValue;
}) => {
    const insets = useSafeAreaInsets();

    return (
        <Animated.View
            style={{
                position: "absolute",
                flexDirection: "row",
                alignItems: "center",
                justifyContent: "flex-start",
                gap: 10,
                width: "90%",
                marginTop: insets.top,
                marginHorizontal: "5%",
                padding: 15,
                borderRadius: 10,
                backgroundColor: bgColor,
                zIndex: 99,
            }}
            entering={SlideInUp.duration(1000)}
            exiting={SlideOutUp.duration(1000)}
        >
            <Ionicons name={icon} size={25} color={color} />
            <Text style={{ color: color, fontSize: 15, fontWeight: "bold" }}>
                {text}
            </Text>
        </Animated.View>
    );
};

export { ToastProvider, useToast };
