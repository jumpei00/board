import axios, { AxiosRequestConfig, AxiosResponse } from "axios";
import { camelizeKeys, decamelizeKeys } from "humps";
import { ApiURL } from "../config/env";

const boardApi = axios.create({
    baseURL: ApiURL,
    withCredentials: true,
});

boardApi.interceptors.request.use((config: AxiosRequestConfig) => {
    if (config.params) {
        config.params = decamelizeKeys(config.params);
    }
    if (config.data) {
        config.data = decamelizeKeys(config.data);
    }
    return config;
});

boardApi.interceptors.response.use((response: AxiosResponse) => {
    if (response.data && response.headers["content-type"].match("application/json")) {
        response.data = camelizeKeys(response.data)
    }
    return response
})

export default boardApi;
