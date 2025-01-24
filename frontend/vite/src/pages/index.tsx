import React, { useState, useEffect } from 'react';
import axios from 'axios';
import {BarChart, ChartContainer, LineChart, PieChart} from '@mui/x-charts';
import { Card, CardContent, Typography } from '@mui/material';
import {Transaction, TransactionType} from "./transactions";
import {API_URL} from "../App";

export default function DashboardPage() {
  const [transactions, setTransactions] = useState<Transaction[]>([]);
  const [withdrawalAmount, setWithdrawalAmount] = useState(0)
  const [depositAmount, setDepositAmount] = useState(0)

  useEffect(() => {
    axios.get(`${API_URL}/transactions`)
      .then(response => {
        const fetchedTransactions = response.data.transactions;
        setTransactions(fetchedTransactions);
        let da = 0
        let wo = 0
        for (let i = 0; i < fetchedTransactions.length; i++) {
          if (fetchedTransactions[i].type == TransactionType.Withdrawal) {
            wo+=fetchedTransactions[i].amount
          } else if (fetchedTransactions[i].type == TransactionType.Deposit) {
            da+=fetchedTransactions[i].amount
          }
        }
        setWithdrawalAmount(wo)
        setDepositAmount(da)
    });
  }, []);

  const s = [
      { label: 'Expenses', data: [withdrawalAmount]},
      { label: 'Income', data: [depositAmount]},
    ]

  return (
      <Card sx={{ width: '100%'}}>
        <CardContent sx={{height: "400px"}}>
          <BarChart
            // xAxis={[{
            //   colorMap: {
            //     type: 'ordinal',
            //     colors: ['red', 'green']
            //   },
            // }]}
            yAxis={[{label: "Income vs Expenses"}]}
            series={s}
            barlabels="value"
            axisHighlight={{
              x: 'none',
              y: 'line'
            }}
          />
        </CardContent>
      </Card>
  );
}
