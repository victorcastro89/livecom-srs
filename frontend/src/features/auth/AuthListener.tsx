/* eslint-disable @typescript-eslint/no-unused-vars */
import React, { useEffect } from 'react';
import { redirect, useLocation } from 'react-router-dom';
import {
  setFirebaseUser,
  AuthStateType,
  FirebaseAuthStateType,
  setFirebaseUserError,
} from './authSlice';
import { useAppDispatch } from '../../hooks';
import { getAuth } from 'firebase/auth';

import { useNavigate } from 'react-router-dom';
import { firebaseInstance } from '../../App';
import store from '../../store';
import { createOrGetUser, logOut } from './authThunk';
import { Roles } from '../../factory/apiFactory';

const AuthListener = () => {
  const dispatch = useAppDispatch();
  const navigate = useNavigate();
  const location = useLocation();
  const currentPath = location.pathname;

  useEffect(() => {
    const auth = getAuth(firebaseInstance);
    const unsubscribe = auth.onAuthStateChanged((user: any) => {
      console.log('Auth State changed');

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
        store
          .dispatch(
            createOrGetUser({
              first_name: user.displayName || '',
              last_name: '',
              phone_number: user.phoneNumber || '',
              photo_url: user.photoURL || '',
              account_name: user.email?.split('@')[0] || '',
              role: Roles.OWNER,
            })
          )
          .unwrap()
          .then(() => {
            navigate(currentPath);
          })
          //TOTO : handle error
          .catch((err) => {
            console.log('error on Dispatch createOrGetUser', err);
            store.dispatch(
              setFirebaseUserError((err as string) || 'Failed to login')
            );
            console.log('Redirecting');
            dispatch(logOut());
            navigate('/login');
          });
      } else {
        console.log('Redirecting');
        dispatch(logOut());
        navigate('/login');
      }
    });

    return () => unsubscribe(); // Cleanup on unmount
  }, [dispatch]);

  return null; // This component does not render anything
};

export default AuthListener;
// Generic function to copy properties
