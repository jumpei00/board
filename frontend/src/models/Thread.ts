export interface Thread {
    threadKey: string;
    title: string;
    contributer: string;
    postDate: string;
    updateDate: string;
    views: number;
    sumComment: number;
}

export interface Threads {
    threads: Array<Thread>;
}
