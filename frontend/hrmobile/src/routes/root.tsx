import { Link as RouterLink, LinkProps as RouterLinkProps, Outlet } from "react-router-dom";
import Box from '@mui/material/Box';
import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import Drawer from '@mui/material/Drawer';
import { forwardRef, useState } from "react";
import List from '@mui/material/List';
import ListItemButton from '@mui/material/ListItemButton';
import ListItemText from '@mui/material/ListItemText';
import ListItem from '@mui/material/ListItem';
import Divider from '@mui/material/Divider';
import ListItemIcon from '@mui/material/ListItemIcon';
import MenuIcon from '@mui/icons-material/Menu';
import LogoutIcon from '@mui/icons-material/Logout';
import WorkHistoryIcon from '@mui/icons-material/WorkHistory';
import PendingActionsIcon from '@mui/icons-material/PendingActions';
import SettingsSuggestIcon from '@mui/icons-material/SettingsSuggest';
import TodayIcon from '@mui/icons-material/Today';
import MoreTimeIcon from '@mui/icons-material/MoreTime';
import ArticleIcon from '@mui/icons-material/Article';

interface ListItemLinkProps {
  icon?: React.ReactElement;
  title: string;
  to: string;
}

function ListItemLink(props: ListItemLinkProps) {
  const { icon, title, to } = props;

  return (
    <li>
      <ListItem component={Link} to={to}>
        {icon ? <ListItemIcon>{icon}</ListItemIcon> : null}
        <ListItemText primary={title} />
      </ListItem>
    </li>
  );
}

const Link = forwardRef<HTMLAnchorElement, RouterLinkProps>(function Link(
  itemProps,
  ref,
) {
  return <RouterLink ref={ref} {...itemProps} role={undefined} />;
});

export default function Root() {

  const [drawerState, setDrawerState] = useState(false)
  
  const toggleDrawer =
    (open: boolean) =>
    (event: React.KeyboardEvent | React.MouseEvent) => {
      if (
        event.type === 'keydown' &&
        ((event as React.KeyboardEvent).key === 'Tab' ||
          (event as React.KeyboardEvent).key === 'Shift')
      ) {
        return;
      }

      setDrawerState(open);
    };

    const topMenus = [
      {
        icon: <MoreTimeIcon />,
        title: "CheckIn",
        link: "checkin"
      },
      {
        icon: <PendingActionsIcon />,
        title: "Approvals",
        link: "approvals"
      },
      {
        icon: <WorkHistoryIcon />,
        title: "Timesheet",
        link: "timesheet"
      },
      {
        icon: <ArticleIcon />,
        title: "Payslip",
        link: "payslip"
      },
      {
        icon: <TodayIcon />,
        title: "Roster",
        link: "roster"
      },
      {
        icon: <SettingsSuggestIcon />,
        title: "Settings",
        link: "settings"
      },
    ]

    const bottomMenus = [
      {
        icon: <LogoutIcon />,
        title: "Logout",
        link: "logout"
      },
    ]
  
    const list = () => (
      <Box
        sx={{ width: 250 }}
        role="presentation"
        onClick={toggleDrawer(false)}
        onKeyDown={toggleDrawer(false)}
      >
        <List aria-label="main mailbox folders">
          {topMenus.map(({icon, title, link}) => (
            <ListItemLink to={link} title={title} icon={icon} />
          ))}
        </List>
        <Divider />
        <List>
          {bottomMenus.map(({icon, title}) => (
            <ListItem key={title} disablePadding>
              <ListItemButton>
                <ListItemIcon>
                  {icon}
                </ListItemIcon>
                <ListItemText primary={title} />
              </ListItemButton>
            </ListItem>
          ))}
        </List>
      </Box>
    );

  return (
    <Box sx={{ flexGrow: 1 }}>
      <AppBar position="static">
        <Toolbar>
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
            onClick={toggleDrawer(true)}
          >
            <MenuIcon />
          </IconButton>
          <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
            HRAPP
          </Typography>
          <Button color="inherit">Logout</Button>
        </Toolbar>
      </AppBar>
      <Drawer
          anchor="left"
          open={drawerState}
          onClose={toggleDrawer(false)}
        >
        {list()}
      </Drawer>
      <Box>
        <Outlet />
      </Box>
    </Box>
  )
}