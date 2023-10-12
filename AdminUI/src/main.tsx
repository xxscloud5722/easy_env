import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.less';
import Layout from './layout/Layout.tsx';

const root = document.getElementById('root');
if (root != null) {
  ReactDOM.createRoot(root)
    .render(
      <React.StrictMode>
        <Layout/>
      </React.StrictMode>
    );
}
