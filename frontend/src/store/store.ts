import { configureStore } from "@reduxjs/toolkit";
import createSagaMiddleware from "@redux-saga/core";
import { threadReducer } from "../state/threads/modules";
import { visitorsReducer } from "../state/visitor/modules";
import { commentReducer } from "../state/comments/modules";
import { userReducer } from "../state/user/modules";
import rootSaga from "./middleware/saga";
import { ENV } from "../config/env";
import logger from "./middleware/logger";

// saga configure
const sagaMiddleware = createSagaMiddleware();

// middlewares
const middlewares = []

middlewares.push(sagaMiddleware)
if (ENV === "development") {
    middlewares.push(logger)
}

export const store = configureStore({
    reducer: {
        user: userReducer,
        visitor: visitorsReducer,
        thread: threadReducer,
        comment: commentReducer,
    },
    middleware: middlewares
});

// saga run
sagaMiddleware.run(rootSaga);

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
