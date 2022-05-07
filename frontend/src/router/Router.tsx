import React from "react";
import { Routes, Route } from "react-router-dom";
import { Header } from "../components/templates/header/Header";
import { Home } from "../pages/home/Page";
import { SingUp } from "../pages/singup/Page";
import { SignIn } from "../pages/signin/Page";
import { ThreadDetail } from "../pages/threadDetail/Page";
import { NotFind } from "../pages/404/Page";

export const Router: React.FC = () => {
    return (
        <Routes>
            <Route path="/" element={<Header></Header>}>
                <Route index element={<Home></Home>}></Route>
                <Route path="signup" element={<SingUp></SingUp>}></Route>
                <Route path="signin" element={<SignIn></SignIn>}></Route>
                <Route
                    path="detail"
                    element={
                        <ThreadDetail
                            hashID="1"
                            title="テスト"
                            contributer="jumpei00"
                            postDate="2022/1/1 12:00"
                            updateDate="2022/1/1 13:00"
                            views={200}
                            sumComment={100}
                        ></ThreadDetail>
                    }
                ></Route>
                <Route path="*" element={<NotFind></NotFind>}></Route>
            </Route>
        </Routes>
    );
}