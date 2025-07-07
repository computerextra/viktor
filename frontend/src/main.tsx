import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import { HashRouter, Route, Routes } from 'react-router'
import Layout from './Layout'
import Overview from './Pages/CMS/Overview'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <HashRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/">
            <Route index element={<>Home</>} />
            <Route path="CMS">
              <Route index element={<Overview />} />
            </Route>
          </Route>
        </Route>
      </Routes>
    </HashRouter>
  </StrictMode>,
)
