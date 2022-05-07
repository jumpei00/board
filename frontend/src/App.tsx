import React from "react";
import { ChakraProvider } from "@chakra-ui/react";
import { BrowserRouter } from "react-router-dom";
import theme from "./theme/Theme";
import { Router } from "./router/Router";

const App: React.FC = () => {
    return (
        <ChakraProvider theme={theme}>
            <BrowserRouter>
                <Router></Router>
            </BrowserRouter>
        </ChakraProvider>
    );
};

export default App;
