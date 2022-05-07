import React from "react";
import { VisitorStat } from "../../components/organisms/stat/VisitorStat";
import { ThreadPostForm } from "../../components/organisms/form/ThreadPostForm";
import { Thread } from "../../models/Thread";
import { ThreadsBoardList } from "../../components/templates/threads/ThreadBoardList";

export const Home: React.FC = () => {
    const testThreads: Thread[] = [
        {
            threadKey: "1",
            title: "test",
            contributer: "motohashi",
            postDate: "2022/1/1 12:00",
            updateDate: "2022/1/1 13:00",
            views: 10,
            sumComment: 20,
        },
        {
            threadKey: "2",
            title: "test",
            contributer: "motohashi",
            postDate: "2022/1/1 12:00",
            updateDate: "2022/1/1 13:00",
            views: 10,
            sumComment: 20,
        },
    ];

    return (
        <>
            <VisitorStat
                yesterdayVisitor={0}
                todayVisitor={0}
                sumVisitor={0}
            ></VisitorStat>
            <ThreadPostForm></ThreadPostForm>
            <ThreadsBoardList threads={testThreads}></ThreadsBoardList>
        </>
    );
};
