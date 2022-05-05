import React from "react";
import { Stat, StatLabel, StatNumber, StatGroup } from "@chakra-ui/react";

type VisitorStatPops = {
    yesterdayVisitor: number;
    todayVisitor: number;
    sumVisitor: number;
};

export const VisitorStat: React.FC<VisitorStatPops> = (props) => {
    const { yesterdayVisitor, todayVisitor, sumVisitor } = props;

    return (
        <StatGroup w="50%" m="50px auto">
            <Stat textAlign="center">
                <StatLabel>昨日の訪問者</StatLabel>
                <StatNumber>{yesterdayVisitor}人</StatNumber>
            </Stat>
            <Stat textAlign="center">
                <StatLabel>今日の訪問者</StatLabel>
                <StatNumber>{todayVisitor}人</StatNumber>
            </Stat>
            <Stat textAlign="center">
                <StatLabel>これまでの訪問者</StatLabel>
                <StatNumber>{sumVisitor}人</StatNumber>
            </Stat>
        </StatGroup>
    );
};
