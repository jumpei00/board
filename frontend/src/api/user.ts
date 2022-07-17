import { SignInPayload, SignUpPayload } from "../state/user/modules";
import { boardApi } from "./init";

export const getMe = () => {
    return boardApi.get(`/api/user/me`)
};

export const signup = (payload: SignUpPayload) => {
    return boardApi.post(`/api/user/signup`, payload);
}

export const signin = (payload: SignInPayload) => {
    return boardApi.post(`/api/user/signin`, payload)
}

export const signout = () => {
    return boardApi.delete(`/api/user/signout`)
}
