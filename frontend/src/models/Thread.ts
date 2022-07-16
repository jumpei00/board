// ---- state ---- //
export interface Thread {
    threadKey: string;
    title: string;
    contributor: string;
    views: number;
    commentSum: number;
    createDate: string;
    updateDate: string;
}

export interface Threads {
    threads: Array<Thread>;
}

// ---- saga response ---- //
export interface FetchThreadResponse {
    threads: Array<Thread>
}

export interface CreateThreadResponse {
    threadKey: string;
    title: string;
    contributor: string;
    views: number;
    commentSum: number;
    createDate: string;
    updateDate: string;
}

export interface UpdateThreadResponse {
    threadKey: string;
    title: string;
    contributor: string;
    views: number;
    commentSum: number;
    createDate: string;
    updateDate: string;
}

export interface ThreadResponse {
    fetchResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: FetchThreadResponse;
    };
    createResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: CreateThreadResponse;
    };
    updateResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: UpdateThreadResponse;
    };
    deleteResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
    };
}
