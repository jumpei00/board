type envType = "development" | "production";
type envMap = {
    development: string;
    production: string;
};

const ENV: envType = process.env.WEB_ENV as envType;

const urlMap: envMap = {
    development: "http://localhost.api",
    production: "",
};

export const ApiURL = urlMap[ENV]
