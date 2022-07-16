import axios from "axios";
import { ApiURL } from "../config/env";

export const boardApi = axios.create({
    baseURL: ApiURL,
    withCredentials: true
})