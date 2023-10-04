// react
import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import reportWebVitals from './reportWebVitals';

// top level theme and navigation
import { createTheme } from '@mui/material/styles';
import { ThemeProvider } from '@mui/system';
import Header from './navigation/Header.js';
import Body from './navigation/Body.js';
import Footer from './navigation/Footer.js';
import styles from './navigation/Navigation.module.scss';

// react root
const root = ReactDOM.createRoot(document.getElementById('root'));

// mui root theme
const theme = createTheme({
  palette: {
    background: {
      main: '#333333',
    },
  },
});

// render
root.render(
  <React.StrictMode>
    <ThemeProvider theme={theme}>
      <div className={styles.Navigation}>
        <Header />
        <Body />
        <Footer />
      </div>
    </ThemeProvider>
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();
