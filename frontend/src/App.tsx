import React, { useState } from 'react';
import { CssVarsProvider, useTheme } from '@mui/joy/styles';

import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';

import { About } from './pages/About';
// import { Home } from './pages/Home'
import { Login } from './pages/Login';
import { SignedIn } from './pages/SignedIn';
import AuthListener from './features/auth/AuthListener';
import { MainDesktop } from './pages/MainDesktop';
import { Outlet } from 'react-router-dom';
import { FirebaseOptions, initializeApp } from 'firebase/app';
import { browserLocalPersistence, getAuth } from 'firebase/auth';
import theme from './theme/theme';
import { ApidDocs } from './pages/ApiDocs';

import { useMediaQuery } from '@mui/material';

import { MainMobile } from './pages/MainMobile';
import { TeamPage } from './pages/Team';
import { Player } from './pages/Player';
// import './css/firebase_styling.global.css';
// Configure Firebase.
const config: FirebaseOptions = {
  apiKey: 'AIzaSyB2TjivCm9ZmtDJVH3Lr9qYcxL3zoKsoa0',
  authDomain: 'instacom-auth.firebaseapp.com',
  appId: '226538007405',
};
export const firebaseInstance = initializeApp(config);
export const auth = getAuth(firebaseInstance);

auth.setPersistence(browserLocalPersistence).then(() => {});
function Layout() {
  const isMobile = useMediaQuery(theme.breakpoints.down('sm'));

  return (
    <CssVarsProvider defaultColorScheme="dark" theme={theme}>
      <AuthListener />
      {isMobile ? <MainMobile /> : <MainDesktop />}
    </CssVarsProvider>
  );
}

function LoginLayout() {
  return (
    <CssVarsProvider defaultColorScheme="dark" theme={theme}>
      <Login></Login>
    </CssVarsProvider>
  );
}
const router = createBrowserRouter([
  {
    path: '/apidoc',
    element: <ApidDocs />,
  },
  {
    path: '/login',
    element: <LoginLayout />,
  },

  {
    element: <Layout />,
    children: [
      {
        path: '/login',
        element: <Login />,
      },
      {
        path: '/team',
        element: <TeamPage />,
      },
      {
        path: '/',
        element: <Player />,
      },
      {
        path: '/signedIn',
        element: <SignedIn />,
        loader: () => {
          return { auth: auth };
        },
      },
    ],
  },
]);

export default function App() {
  return <RouterProvider router={router} />;
}
