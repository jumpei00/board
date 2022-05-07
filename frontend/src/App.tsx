import React from "react";
import { ChakraProvider } from "@chakra-ui/react";
import theme from "./theme/Theme";
import { Header } from "./components/templates/header/Header";

const App: React.FC = () => {
    return (
        <ChakraProvider theme={theme}>
            <Header></Header>
        </ChakraProvider>
    );
};

export default App;
