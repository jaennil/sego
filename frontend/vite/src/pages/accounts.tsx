import Typography from "@mui/material/Typography";
import * as React from "react";
import {DataGrid, GridColDef} from "@mui/x-data-grid";
import {useEffect} from "react";
import {API_URL} from "../App";
import axios from "axios";
import {Paper} from "@mui/material";

const columns: GridColDef[] = [
  { field: 'name', headerName: 'Name'},
  { field: 'type', headerName: 'Type'},
  { field: 'balance', headerName: 'Balance'},
  { field: 'currency', headerName: 'Currency Code', minWidth: 115},
  { field: 'created_at', headerName: 'Opened At', minWidth: 170},
]

export interface Account {
  id: string;
  name: string;
  type: string;
  balance: number;
  created_at: string;
  currency_code: string;
}

export enum AccountType {
  CreditCard = "Credit Card",
  Cash = "Cash",
}

export default function AccountsPage() {
  const [accounts, setAccounts] = React.useState<Account[]>([]);

  useEffect(() => {
    axios.get(API_URL+"/accounts")
      .then(response => {
        setAccounts(response.data.accounts)
      }).catch(error => {
        console.error(error)
    })
  }, [])

  return (
    <Paper sx={{ height: 400, width: '100%'}}>
      <DataGrid
        columns={columns}
        rows={accounts.map(account => {
          account.id = account.name
          return account
        })}
      />
    </Paper>
  )
}
