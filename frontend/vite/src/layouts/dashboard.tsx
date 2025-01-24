import * as React from 'react';
import {Outlet} from 'react-router';
import {DashboardLayout} from '@toolpad/core/DashboardLayout';
import {PageContainer} from '@toolpad/core/PageContainer';
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogTitle,
  Divider,
  MenuItem,
  Select,
  SelectChangeEvent,
  Stack,
  TextField
} from "@mui/material";
import Typography from "@mui/material/Typography";
import {API_URL, APP_NAME} from "../App";
import PaidIcon from '@mui/icons-material/Paid';
import AccountBalanceIcon from '@mui/icons-material/AccountBalance';
import ReceiptIcon from "@mui/icons-material/Receipt";
import {NumericFormat} from "react-number-format";
import {DateTimePicker} from "@mui/x-date-pickers";
import {AdapterDayjs} from "@mui/x-date-pickers/AdapterDayjs";
import {LocalizationProvider} from "@mui/x-date-pickers/LocalizationProvider";
import currencyFormatter from 'currency-formatter'
import axios from 'axios'
import dayjs from "dayjs";
import {Account, AccountType} from "../pages/accounts";
import { TransactionType } from '../pages/transactions';
import {useEffect} from "react";


function CustomAppTitle({ handleNewAccountOpen, handleNewTransactionOpen }: { handleNewAccountOpen: () => void, handleNewTransactionOpen: () => void }) {
  return (
    <Stack direction="row" alignItems="center" spacing={2}>
      <PaidIcon fontSize="large" color="primary" />
      <Typography variant="h6">{APP_NAME}</Typography>
      <Divider orientation="vertical" />
      <Button
        variant="outlined"
        endIcon={<AccountBalanceIcon />}
        onClick={handleNewAccountOpen}
      >
        New Account
      </Button>
      <Button
        variant="outlined"
        endIcon={<ReceiptIcon />}
        onClick={handleNewTransactionOpen}
      >
        New Transaction
      </Button>
    </Stack>
  );
}

export interface Category {
  title: string;
}

export function NewTransactionDialog({ open, handleClose }: { open: boolean; handleClose: () => void }) {
  const [type, setType] = React.useState('');
  const [accounts, setAccounts] = React.useState<Account[]>([]);
  const [categories, setCategories] = React.useState<Category[]>([]);
  const [category, setCategory] = React.useState('')
  const [account, setAccount] = React.useState('')

  const handleTypeChange = (event: SelectChangeEvent) => {
    setType(event.target.value as string)
  }

  const handleAccountChange = (event: SelectChangeEvent) => {
    setAccount(event.target.value as string)
  }

  const handleCategoryChange = (event: SelectChangeEvent) => {
    setCategory(event.target.value as string)
  }

  useEffect(() => {
    axios.get(API_URL+"/accounts")
      .then(response => {
        setAccounts(response.data.accounts)
      }).catch(error => {
      console.error(error)
    })
    axios.get(API_URL+"/categories")
      .then(response => {
        setCategories(response.data.categories)
      }).catch(error => {
      console.error(error)
    })
  }, [])

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      PaperProps={{
        component: 'form',
        onSubmit: (event: React.FormEvent<HTMLFormElement>) => {
          event.preventDefault()
          const formData = new FormData(event.currentTarget);
          const formJson = Object.fromEntries((formData as any).entries());
          formJson.amount = +formJson.amount
          formJson.created_at = dayjs(formJson.created_at, 'DD/MM/YYYY HH:mm').toISOString()
          axios.post(API_URL+"/transaction", formJson)
            .then((response) => {
              console.log('Response:', response.data)
            })
            .catch((error) => {
              console.error('Error:', error)
            })
          handleClose();
        }
      }}
    >
      <DialogTitle>New Transaction</DialogTitle>
      <DialogContent>
        <Stack spacing={2}>
          <Select
            required
            autoFocus
            fullWidth
            name="type"
            id="new-transaction-dialog-type"
            value={type}
            // label="Type"
            onChange={handleTypeChange}
          >
            {
              Object.values(TransactionType).map(type => <MenuItem value={type}>{type}</MenuItem>)
            }
          </Select>
          <NumericFormat
            id="new-transaction-dialog-transaction-amount"
            name="amount"
            label="Amount"
            required
            fullWidth
            customInput={TextField}
          />
          <LocalizationProvider dateAdapter={AdapterDayjs}>
            <DateTimePicker
              disableFuture
              ampm={false}
              slotProps={{textField: {fullWidth: true, required: true}}}
              name="created_at"
              label="Date"
            />
          </LocalizationProvider>
          <Select
            required
            autoFocus
            fullWidth
            name="category"
            id="new-transaction-dialog-category"
            value={category}
            // label="Type"
            onChange={handleCategoryChange}
          >
            {
              Object.values(categories).map(category => <MenuItem value={category.title}>{category.title}</MenuItem>)
            }
          </Select>
          <Select
            required
            autoFocus
            fullWidth
            name="account"
            id="new-transaction-dialog-account"
            value={account}
            // label="Type"
            onChange={handleAccountChange}
          >
            {
              Object.values(accounts).map(account => <MenuItem value={account.name}>{account.name}</MenuItem>)
            }
          </Select>
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>Cancel</Button>
        <Button type="submit">Create</Button>
      </DialogActions>
    </Dialog>
  );
}

export function NewAccountDialog({ open, handleClose }: { open: boolean; handleClose: () => void }) {
  const [type, setType] = React.useState('');
  const [currency, setCurrency] = React.useState('')

  const handleTypeChange = (event: SelectChangeEvent) => {
    setType(event.target.value as string)
  }

  const handleCurrencyChange = (event: SelectChangeEvent) => {
    setCurrency(event.target.value as string)
  }

  return (
    <Dialog
      open={open}
      onClose={handleClose}
      PaperProps={{
        component: 'form',
        onSubmit: (event: React.FormEvent<HTMLFormElement>) => {
          event.preventDefault()
          const formData = new FormData(event.currentTarget);
          const formJson = Object.fromEntries((formData as any).entries());
          formJson.balance = +formJson.balance
          formJson.created_at = dayjs(formJson.created_at, 'DD/MM/YYYY HH:mm').toISOString()
          axios.post(API_URL+"/account", formJson)
            .then((response) => {
              console.log('Response:', response.data)
            })
            .catch((error) => {
              console.error('Error:', error)
            })
          handleClose();
        }
      }}
    >
      <DialogTitle>New Account</DialogTitle>
      <DialogContent>
        <Stack spacing={2}>
        <Select
          required
          autoFocus
          fullWidth
          name="type"
          id="new-account-dialog-type"
          value={type}
          // label="Type"
          onChange={handleTypeChange}
        >
          {
            Object.values(AccountType).map(type => <MenuItem value={type}>{type}</MenuItem>)
          }
        </Select>
        <TextField
          required
          id="new-account-dialog-account-name"
          fullWidth
          name="name"
          label="Name"
        />
        <NumericFormat
          id="new-account-dialog-account-balance"
          name="balance"
          label="Balance"
          required
          fullWidth
          customInput={TextField}
        />
        <LocalizationProvider dateAdapter={AdapterDayjs}>
          <DateTimePicker
            disableFuture
            ampm={false}
            slotProps={{textField: {fullWidth: true, required: true}}}
            name="created_at"
            label="Open date"
          />
        </LocalizationProvider>
        <Select
          fullWidth
          value={currency}
          name="currency"
          required
          onChange={handleCurrencyChange}
          label="Currency"
        >
          {currencyFormatter.currencies.map((c) => {
            return (
              <MenuItem value={c.code}>
                {c.code} - {c.symbol}
              </MenuItem>
            );
          })}
        </Select>
        </Stack>
      </DialogContent>
      <DialogActions>
        <Button onClick={handleClose}>Cancel</Button>
        <Button type="submit">Create</Button>
      </DialogActions>
    </Dialog>
  );
}

export default function Layout() {
  const [newAccountOpen, setNewAccountOpen] = React.useState(false);
  const [newTransactionOpen, setNewTransactionOpen] = React.useState(false);

  const handleNewAccountOpen = () => {
    setNewAccountOpen(true);
  };

  const handleNewAccountClose = () => {
    setNewAccountOpen(false);
  };

  const handleNewTransactionOpen = () => {
    setNewTransactionOpen(true);
  };

  const handleNewTransactionClose = () => {
    setNewTransactionOpen(false);
  };

  return (
    <DashboardLayout
      slots={{
        appTitle: () => <CustomAppTitle handleNewAccountOpen={handleNewAccountOpen} handleNewTransactionOpen={handleNewTransactionOpen} />,
      }}
    >
      <PageContainer>
        <Outlet />
        <NewAccountDialog open={newAccountOpen} handleClose={handleNewAccountClose} />
        <NewTransactionDialog open={newTransactionOpen} handleClose={handleNewTransactionClose} />
      </PageContainer>
    </DashboardLayout>
  );
}
