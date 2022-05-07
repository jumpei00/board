import React from "react";
import { Stack } from "@chakra-ui/react";
import { ThreadBoard } from "../../organisms/thread/ThreadBoard";
import { Threads } from "../../../models/Thread";

export const ThreadsBoardList: React.FC<Threads> = (props) => {
    const { threads } = props;

    return (
        <Stack w="70%" m="50px auto" spacing={6}>
            {threads.map((thread) => (
                <ThreadBoard
                    key={thread.threadKey}
                    threadKey={thread.threadKey}
                    title={thread.title}
                    contributer={thread.contributer}
                    postDate={thread.postDate}
                    updateDate={thread.updateDate}
                    views={thread.views}
                    sumComment={thread.sumComment}
                ></ThreadBoard>
            ))}
        </Stack>
    );
};
