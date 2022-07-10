import { configureStore } from "@reduxjs/toolkit";
import { threadsReducer } from "../state/threads";
import { visitorsReducer } from "../state/visitor";
import { commentsReducer } from "../state/comments";
import { threadReducer } from "../pages/threadContent/redux/thread";
import { userReducer } from "../state/user";

export const store = configureStore({
    reducer: {
        user: userReducer,
        visitors: visitorsReducer,
        threads: threadsReducer,
        thread: threadReducer,
        comments: commentsReducer,
    },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
