/* eslint-disable @typescript-eslint/no-unused-vars */
import React, { useEffect } from 'react';
import { redirect } from 'react-router-dom';
import {
  setFirebaseUser,
  logout,
  AuthStateType,
  FirebaseAuthStateType,
} from './authSlice';
import { useAppDispatch } from '../../hooks';
import { getAuth } from 'firebase/auth';

import { useNavigate } from 'react-router-dom';
import { firebaseInstance } from '../../App';

const AuthListener = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  useEffect(() => {
    const auth = getAuth(firebaseInstance);
    const unsubscribe = auth.onAuthStateChanged((user: any) => {
      console.log('Auth State changed');

      console.log(user?.accessToken);
      let u: FirebaseAuthStateType = {
        email: user?.email ? user?.email : null,
        emailVerified: user?.emailVerified ? user?.emailVerified : null,
        isAnonymous: user?.isAnonymous ? user?.isAnonymous : null,
        displayName: user?.displayName ? user?.displayName : null,
        phoneNumber: user?.phoneNumber ? user?.phoneNumber : null,
        photoURL: user?.photoURL ? user?.photoURL : null,
        uid: user?.uid ? user?.uid : null,
        refreshToken: user?.refreshToken ? user?.refreshToken : null,
        accessToken: user?.accessToken ? user?.accessToken : null,
      };

      // if (u) {
      //   copyObjProperties(u, user);
      //   console.log(copyObjProperties(u, user));
      // }
      if (user?.uid) {
        dispatch(setFirebaseUser(u));
      } else {
        console.log('Redirecting');
        dispatch(logout());
        navigate('/login');
      }
    });

    return () => unsubscribe(); // Cleanup on unmount
  }, [dispatch]);

  return null; // This component does not render anything
};

export default AuthListener;
// Generic function to copy properties
