import { create } from "zustand";

type SearchStore = {
    query: string;
    setQuery: (newQuery: string) => void;
};

const useSearchStore = create<SearchStore>()((set) => ({
    query: "",
    setQuery: (newQuery: string) => set({ query: newQuery }),
}));

export { useSearchStore };
