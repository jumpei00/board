import { testSaga } from "redux-saga-test-plan";
import { getMe } from "../../api/user";
import { userActions, userSagaActions } from "./modules";
import { fetchUser } from "./saga";

describe("fetchUser", () => {
    it("can fetch user", () => {
        const resMock = {
            data: {
                id: "id",
                username: "username"
            }
        };

        testSaga(fetchUser)
            .next()
            .call(getMe)
            .next(resMock)
            .put(userSagaActions.getmeDone())
            .next(resMock)
            .put(userActions.storeUser(resMock.data))
            .next()
            .isDone();
    });
});
