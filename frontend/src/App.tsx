import React from 'react';
import { CssVarsProvider } from '@mui/joy/styles';

import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { Navbar } from './components/Navbar';
import { About } from './pages/About';
// import { Home } from './pages/Home'
import { Login } from './pages/Login';
import { SignedIn } from './pages/SignedIn';
import AuthListener from './features/auth/AuthListener';
import { Home } from './pages/Home';
import { Outlet } from "react-router-dom";
import { FirebaseOptions, initializeApp } from 'firebase/app';
import { getAuth } from 'firebase/auth';
import theme from './theme/theme';
import { ApidDocs } from './pages/ApiDocs';
// Configure Firebase.
const config:FirebaseOptions = {
  apiKey: 'AIzaSyB2TjivCm9ZmtDJVH3Lr9qYcxL3zoKsoa0',
  authDomain: 'instacom-auth.firebaseapp.com',
  appId: '226538007405',
};
export const firebaseInstance =initializeApp(config);
export const auth = getAuth(firebaseInstance);

function Layout() {
  return (
      <CssVarsProvider defaultMode="dark"
  
      theme={theme}

 
     >
      <AuthListener />
      <Navbar />
      <Outlet></Outlet>
      </CssVarsProvider>

  );
}

const router = createBrowserRouter([   {
  path: '/apidoc',
  element: <ApidDocs />,
},
  {
    element: <Layout />,
    children: [
      {
        path: '/',
        element: <Login />,
      },
   
      {
        path: '/signedIn',
        element: <SignedIn />,
        loader: () => {
          return {auth:auth}
        }
      },
  
    ],
  },
]);


export default function App() {
  return <RouterProvider router={router} />;
}
