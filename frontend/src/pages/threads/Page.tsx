import React, { useEffect } from "react";
import { useSelector, useDispatch } from "react-redux";
import { VisitorStat } from "./organisms/VisitorStat";
import { ThreadPostForm } from "../../components/organisms/form/ThreadPostForm";
import { ThreadsBoardList } from "./organisms/ThreadBoardList";
import { RootState } from "../../store/store";
import { userSagaActions } from "../../state/user/modules";
import { threadSagaActions } from "../../state/threads/modules";
import { visitorSagaActions } from "../../state/visitor/modules";
import { commentActions } from "../../state/comments/modules";

export const Home: React.FC = () => {
    const userState = useSelector((state: RootState) => state.user.userState);
    const visitorState = useSelector((state: RootState) => state.visitor.visitorState);
    const threadState = useSelector((state: RootState) => state.thread.threadState);
    const dispatch = useDispatch();

    useEffect(() => {
        dispatch(userSagaActions.getme());
        dispatch(threadSagaActions.getAll());
        dispatch(visitorSagaActions.getStat());
        dispatch(commentActions.clearComment());
    }, []);

    return (
        <>
            <VisitorStat
                yesterday={visitorState.yesterday}
                today={visitorState.today}
                sum={visitorState.sum}
            ></VisitorStat>
            {userState.username === "" || <ThreadPostForm></ThreadPostForm>}
            <ThreadsBoardList threads={threadState.threads}></ThreadsBoardList>
        </>
    );
};
