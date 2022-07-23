// ---- state ---- //
export interface Visitor {
    yesterday: number;
    today: number;
    sum: number;
}

// ---- saga state ---- //
export interface VisitorResponse {
    fetchResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: Visitor
    };
    countupResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: Visitor
    };
}
