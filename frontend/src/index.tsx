import React from "react";
import { render } from "react-dom";
import App from "./App";
import { store } from "./store/store";
import { Provider } from "react-redux";

const root = document.getElementById("root") as HTMLElement;
render(
    <React.StrictMode>
        <Provider store={store}>
            <App />
        </Provider>
    </React.StrictMode>,
    root
);
