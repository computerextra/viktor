import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { HashRouter, Route, Routes } from "react-router";
import "./index.css";
import Layout from "./Layout";
import AbteilungBearbeiten from "./Pages/CMS/Abteilungen/id";
import AbteilungAnlegen from "./Pages/CMS/Abteilungen/neu";
import AbteilungenÜbersicht from "./Pages/CMS/Abteilungen/page";
import AngebotBearbeiten from "./Pages/CMS/Angebote/id";
import AngebotAnlegen from "./Pages/CMS/Angebote/neu";
import AngeboteÜbersicht from "./Pages/CMS/Angebote/page";
import JobBearbeiten from "./Pages/CMS/Jobs/id";
import JobAnlegen from "./Pages/CMS/Jobs/neu";
import JobÜbersicht from "./Pages/CMS/Jobs/page";
import Overview from "./Pages/CMS/Overview";

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
                <Route path="Neu" element={<AbteilungAnlegen />} />
                <Route path=":id" element={<AbteilungBearbeiten />} />
              </Route>
              <Route path="Angebote">
                <Route index element={<AngeboteÜbersicht />} />
                <Route path="Neu" element={<AngebotAnlegen />} />
                <Route path=":id" element={<AngebotBearbeiten />} />
              </Route>
              <Route path="Jobs">
                <Route index element={<JobÜbersicht />} />
                <Route path="Neu" element={<JobAnlegen />} />
                <Route path=":id" element={<JobBearbeiten />} />
              </Route>
            </Route>
          </Route>
        </Route>
      </Routes>
    </HashRouter>
  </StrictMode>
);
