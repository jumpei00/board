import { createSlice, PayloadAction } from "@reduxjs/toolkit";
import { Thread } from "../../../../models/Thread";
import { getThreadByThreadKeyPayload } from "./type";

const initialState: Thread = {
    threadKey: "",
    title: "",
    contributer: "",
    postDate: "",
    updateDate: "",
    views: 0,
    sumComment: 0,
};

export const threadSlice = createSlice({
    name: "thread",
    initialState,
    reducers: {
        getThreadByThreadKey: (state, action: PayloadAction<getThreadByThreadKeyPayload>) => {
            for (const ts of action.payload.threads) {
                if (ts.threadKey === action.payload.threadKey) {
                    return ts;
                }
            }
        },
    },
});

export const { getThreadByThreadKey } = threadSlice.actions;
export const threadReducer = threadSlice.reducer;
