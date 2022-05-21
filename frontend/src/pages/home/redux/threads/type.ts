export type createThreadPayload = {
    title: string;
    contributer: string;
};

export type editThreadPayload = {
    threadKey: string;
    title: string;
    contributer: string;
};

export type deleteThreadPayload = {
    threadKey: string;
    contributer: string;
};
