// ---- state ---- //
export interface User {
    id: string;
    username: string;
}

// ---- saga response ---- //
export interface UserResponse {
    fetchResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: User;
    };
    signupResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: User;
    };
    signinResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
        body: User;
    };
    signoutResponse: {
        pending: boolean;
        success: boolean;
        error: boolean;
    };
}
