import React from 'react';

import { createRoot } from 'react-dom/client';
import { Provider } from 'react-redux';
import store from './store';

import './index.css';

import App from './App';
import AuthListener from './features/auth/AuthListener';

const container = document.getElementById('root') as HTMLDivElement;
const root = createRoot(container!);

root.render(
  <Provider store={store}>
    <App />
  </Provider>
);
