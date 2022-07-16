import React, { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { VisitorStat } from "../../components/organisms/stat/VisitorStat";
import { ThreadPostForm } from "../../components/organisms/form/ThreadPostForm";
import { ThreadsBoardList } from "./organisms/ThreadBoardList";
import { RootState } from "../../store/store";
import { getAllThread } from "../../state/threads";
import { getVisitors } from "../../state/visitor";

export const Home: React.FC = () => {
    const user = useSelector((state: RootState) => state.user);
    const visitors = useSelector((state: RootState) => state.visitors);
    const threads = useSelector((state: RootState) => state.threads.threads);
    const dispatch = useDispatch();

    useEffect(() => {
        dispatch(getAllThread());
    }, []);

    useEffect(() => {
        dispatch(getVisitors());
    }, []);

    return (
        <>
            <VisitorStat
                yesterdayVisitor={visitors.yesterdayVisitor}
                todayVisitor={visitors.todayVisitor}
                sumVisitor={visitors.sumVisitor}
            ></VisitorStat>
            {user.username === "" || <ThreadPostForm loginUsername={user.username}></ThreadPostForm>}
            <ThreadsBoardList threads={threads}></ThreadsBoardList>
        </>
    );
};