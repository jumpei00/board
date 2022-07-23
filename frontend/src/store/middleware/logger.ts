import { createLogger } from "redux-logger";

const logger = createLogger({
    collapsed: true,
    diff: true,
});

export default logger
