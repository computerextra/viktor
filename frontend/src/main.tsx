import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import "./index.css";
import { HashRouter, Route, Routes } from "react-router";
import Layout from "./Layout";
import Overview from "./Pages/CMS/Overview";
import AbteilungenÜbersicht from "./Pages/CMS/Abteilungen/page";
import AbteilungBearbeiten from "./Pages/CMS/Abteilungen/id";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <HashRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/">
            <Route index element={<>Home</>} />
            <Route path="CMS">
              <Route index element={<Overview />} />
              <Route path="Abteilungen">
                <Route index element={<AbteilungenÜbersicht />} />
                <Route path="Neu" element={<>Neue Abteilung</>} />
                <Route path=":id" element={<AbteilungBearbeiten />} />
              </Route>
            </Route>
          </Route>
        </Route>
      </Routes>
    </HashRouter>
  </StrictMode>
);
