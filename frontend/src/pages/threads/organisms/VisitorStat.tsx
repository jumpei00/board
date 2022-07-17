import React from "react";
import { Stat, StatLabel, StatNumber, StatGroup } from "@chakra-ui/react";

type VisitorStatPops = {
    yesterday: number;
    today: number;
    sum: number;
};

export const VisitorStat: React.FC<VisitorStatPops> = (props) => {
    const { yesterday, today, sum } = props;

    return (
        <StatGroup w="50%" m="50px auto">
            <Stat textAlign="center">
                <StatLabel>昨日の訪問者</StatLabel>
                <StatNumber>{yesterday}人</StatNumber>
            </Stat>
            <Stat textAlign="center">
                <StatLabel>今日の訪問者</StatLabel>
                <StatNumber>{today}人</StatNumber>
            </Stat>
            <Stat textAlign="center">
                <StatLabel>これまでの訪問者</StatLabel>
                <StatNumber>{sum}人</StatNumber>
            </Stat>
        </StatGroup>
    );
};
