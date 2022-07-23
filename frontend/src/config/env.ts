type envType = "development" | "production";
type envMap = {
    development: string;
    production: string;
};

export const ENV: envType = process.env.REACT_APP_WEB_ENV as envType;

const urlMap: envMap = {
    development: "http://api.localhost.test",
    production: "",
};

export const ApiURL = urlMap[ENV]
