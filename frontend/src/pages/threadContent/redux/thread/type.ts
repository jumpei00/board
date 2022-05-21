import { Thread } from "../../../../models/Thread";

export type getThreadByThreadKeyPayload = {
    threads: Array<Thread>
    threadKey: string;
};
