import React from "react";
import { useSelector } from "react-redux";
import { VisitorStat } from "../../components/organisms/stat/VisitorStat";
import { ThreadPostForm } from "../../components/organisms/form/ThreadPostForm";
import { ThreadsBoardList } from "../../components/templates/threads/ThreadBoardList";
import { RootState } from "../../store/store";

export const Home: React.FC = () => {
    const visitors = useSelector((state: RootState) => state.visitors);
    const threads = useSelector((state: RootState) => state.threads.threads);

    return (
        <>
            <VisitorStat
                yesterdayVisitor={visitors.yesterdayVisitor}
                todayVisitor={visitors.todayVisitor}
                sumVisitor={visitors.sumVisitor}
            ></VisitorStat>
            <ThreadPostForm></ThreadPostForm>
            <ThreadsBoardList threads={threads}></ThreadsBoardList>
        </>
    );
};
