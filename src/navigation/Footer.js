import AppBar from '@mui/material/AppBar';
import Typography from '@mui/material/Typography';
import styles from './Navigation.module.scss';

function Footer() {
  return (
    <AppBar position="static" sx={{ backgroundColor: 'background.main' }}>
      <div className={styles.NavigationFooter}>
        <div></div>
        <Typography variant="body2" component="div">
          ©{new Date().getFullYear()} Chris French
        </Typography>
      </div>
    </AppBar>
  )
}

export default Footer;
