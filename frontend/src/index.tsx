import React from "react";
import { createRoot } from "react-dom/client";
import App from "./App";

const hugahugahuga = "hoge";
const state = {
    hugahugahuga: hugahugahuga
}
console.log(state)

const container = document.getElementById("root") as HTMLElement;
const root = createRoot(container);
root.render(
    <React.StrictMode>
        <App />
    </React.StrictMode>
);
