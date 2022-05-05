import React from "react";
import {ChakraProvider} from "@chakra-ui/react"
import { Home } from "./pages/home/Page";
import theme from "./theme/Theme";

const App: React.FC = () =>  {
    return (
        <ChakraProvider theme={theme}>
            <Home></Home>
        </ChakraProvider>
    )
}

export default App;
