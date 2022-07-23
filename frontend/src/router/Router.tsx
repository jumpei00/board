import React from "react";
import { Routes, Route } from "react-router-dom";
import { Header } from "../components/organisms/header/Header";
import { Home } from "../pages/threads/Page";
import { SingUp } from "../pages/singup/Page";
import { SignIn } from "../pages/signin/Page";
import { ThreadContent } from "../pages/threads/comments/Page";
import { NotFind } from "../pages/404/Page";

export const Router: React.FC = () => {
    return (
        <Routes>
            <Route path="/" element={<Header></Header>}>
                <Route index element={<Home></Home>}></Route>
                <Route path="signup" element={<SingUp></SingUp>}></Route>
                <Route path="signin" element={<SignIn></SignIn>}></Route>
                <Route path="thread/:threadKey" element={<ThreadContent></ThreadContent>}></Route>
                <Route path="*" element={<NotFind></NotFind>}></Route>
            </Route>
        </Routes>
    );
};
