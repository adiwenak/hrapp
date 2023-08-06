import React from 'react'
import ReactDOM from 'react-dom/client'
import Root from "./routes/root";
import { RouterProvider, createBrowserRouter } from 'react-router-dom'
import ErrorPage from './error-page';
import CheckInPage from './routes/checkin';
import TimesheetPage from './routes/timesheet';
import RosterPage from './routes/roster';
import PayslipPage from './routes/payslip';
import SettingsPage from './routes/settings';
import ApprovalPage from './routes/approvals';
import '@fontsource/roboto/400.css'
import CssBaseline from '@mui/material/CssBaseline';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    errorElement: <ErrorPage />,
    children: [
      {
        index: true,
        path: "checkin",
        element: <CheckInPage />
      },
      {
        path: "timesheet",
        element: <TimesheetPage />
      },
      {
        path: "payslip",
        element: <PayslipPage />
      },
      {
        path: "approvals",
        element: <ApprovalPage />
      },
      {
        path: "roster",
        element: <RosterPage />
      },
      {
        path: "settings",
        element: <SettingsPage />
      }
    ]
  },
])

ReactDOM.createRoot(document.getElementById('root')!).render(
  <React.StrictMode>
      <CssBaseline />
      <RouterProvider router={router} />
  </React.StrictMode>,
)
