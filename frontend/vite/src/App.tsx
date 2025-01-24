import * as React from 'react';
import DashboardIcon from '@mui/icons-material/Dashboard';
import ReceiptIcon from '@mui/icons-material/Receipt';
import { Outlet } from 'react-router';
import { ReactRouterAppProvider } from '@toolpad/core/react-router';
import type { Navigation } from '@toolpad/core/AppProvider';
import logo from './assets/logo.png'
import AccountBalanceIcon from "@mui/icons-material/AccountBalance";
import {LocalizationProvider} from "@mui/x-date-pickers/LocalizationProvider";
import { AdapterDayjs } from '@mui/x-date-pickers/AdapterDayjs';
import { CssBaseline } from '@mui/material';

export const APP_NAME = "Money Manager Ex"
export const API_URL = "http://127.0.0.1:8080/api"

const NAVIGATION: Navigation = [
  {
    title: 'Dashboard',
    icon: <DashboardIcon />,
  },
  {
    segment: 'accounts',
    title: 'Accounts',
    icon: <AccountBalanceIcon />,
  },
  {
    segment: 'transactions',
    title: 'Transactions',
    icon: <ReceiptIcon />,
  },
];

export default function App() {
  return (
    <CssBaseline enableColorScheme>
      <ReactRouterAppProvider
        navigation={NAVIGATION}
      >
        <Outlet />
      </ReactRouterAppProvider>
    </CssBaseline>
  );
}
