import { createRoot } from 'react-dom/client'
import { BrowserRouter } from 'react-router-dom';
import './main.scss'
import App from './App.tsx'
import '@ant-design/v5-patch-for-react-19';

createRoot(document.getElementById('root')!).render(
  <BrowserRouter>
    <App />
  </BrowserRouter>,
)
