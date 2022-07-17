import { configureStore, MiddlewareArray } from "@reduxjs/toolkit";
import createSagaMiddleware from "@redux-saga/core";
import { threadReducer } from "../state/threads/modules";
import { visitorsReducer } from "../state/visitor/modules";
import { commentReducer } from "../state/comments/modules";
import { userReducer } from "../state/user/modules";
import rootSaga from "./middleware/saga";

// saga configure
const sagaMiddleware = createSagaMiddleware()

export const store = configureStore({
    reducer: {
        user: userReducer,
        visitor: visitorsReducer,
        thread: threadReducer,
        comment: commentReducer,
    },
    middleware: new MiddlewareArray().concat(sagaMiddleware)
});

// saga run
sagaMiddleware.run(rootSaga)

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
