import React from "react";
import { ChakraProvider } from "@chakra-ui/react";
import theme from "./theme/Theme";
import { ThreadDetail } from "./pages/threadDetail/Page";
// import { Home } from "./pages/home/Page";

const App: React.FC = () => {
    return (
        <ChakraProvider theme={theme}>
            {/* <Home></Home> */}
            <ThreadDetail
                hashID="1"
                title="test"
                contributer="motohashi"
                postDate="2020/1/1 12:00:00"
                updateDate="2022/1/1 13:00"
                views={10}
                sumComment={20}
            ></ThreadDetail>
        </ChakraProvider>
    );
};

export default App;
