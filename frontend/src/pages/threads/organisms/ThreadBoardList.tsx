import React from "react";
import { Stack } from "@chakra-ui/react";
import { ThreadBoard } from "../../../components/organisms/thread/ThreadBoard";
import { Threads } from "../../../models/Thread";

export const ThreadsBoardList: React.FC<Threads> = (props) => {
    return (
        <Stack w="70%" m="50px auto" spacing={6}>
            {props.threads.map((thread) => (
                <ThreadBoard
                    key={thread.threadKey}
                    thread={thread}
                ></ThreadBoard>
            ))}
        </Stack>
    );
};
