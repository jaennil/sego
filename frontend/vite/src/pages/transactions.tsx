import * as React from 'react';
import Typography from '@mui/material/Typography';
import {DataGrid, GridColDef} from "@mui/x-data-grid";
import {useEffect} from "react";
import {Category} from "../layouts/dashboard";
import axios from "axios";
import {API_URL} from "../App";
import {Paper} from "@mui/material";

const columns: GridColDef[] = [
  { field: 'amount', headerName: 'Amount'},
  { field: 'type', headerName: 'Type'},
  { field: 'account', headerName: 'Account'},
  { field: 'category', headerName: 'Category'},
  { field: 'created_at', headerName: 'Created At', width: 170},
]

export interface Transaction {
  id: string;
  amount: number;
  type: string;
  account: string;
  category: string;
  created_at: string;
}

export default function TransactionsPage() {
  const [transactions, setTransactions] = React.useState<Transaction[]>([]);

  useEffect(() => {
    axios.get(API_URL+"/transactions")
      .then(response => {
        setTransactions(response.data.transactions)
      }).catch(error => {
      console.error(error)
    })
  }, [])

  return (
    <Paper sx={{ height: 400, width: '100%'}}>
      <DataGrid
        columns={columns}
        rows={transactions}
      />
    </Paper>
  )
}

export enum TransactionType {
  Deposit = "Deposit",
  Withdrawal = "Withdrawal",
}