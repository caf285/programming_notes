import AppBar from '@mui/material/AppBar';
import Toolbar from '@mui/material/Toolbar';
import Box from '@mui/material/Box';
import Typography from '@mui/material/Typography';
import Button from '@mui/material/Button';
import IconButton from '@mui/material/IconButton';
import MenuIcon from '@mui/icons-material/Menu';

import styles from './Navigation.module.scss';

function Header() {
  return (
    <AppBar position="static" sx={{ backgroundColor: 'background.main' }}>
      <div className={styles.NavigationHeader}>
        <Toolbar>
          <IconButton
            size="large"
            edge="start"
            color="inherit"
            aria-label="menu"
            sx={{ mr: 2 }}
          >
            <MenuIcon />
          </IconButton>
          <Box>
            <Button href="/" color="inherit">
              <Typography variant="h6" component="div">
                Home
              </Typography>
            </Button>
          </Box>
        </Toolbar>
      </div>
    </AppBar>
  );
}

export default Header;
