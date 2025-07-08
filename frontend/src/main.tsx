import { lazy, StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { BrowserRouter, Route, Routes } from "react-router";
import "./index.css";

const Layout = lazy(() => import("./Layout"));
const Archive = lazy(() => import("./Pages/Archiv/page"));
const AbteilungBearbeiten = lazy(() => import("./Pages/CMS/Abteilungen/id"));
const AbteilungAnlegen = lazy(() => import("./Pages/CMS/Abteilungen/neu"));
const AbteilungenÜbersicht = lazy(() => import("./Pages/CMS/Abteilungen/page"));
const AngebotBearbeiten = lazy(() => import("./Pages/CMS/Angebote/id"));
const AngebotAnlegen = lazy(() => import("./Pages/CMS/Angebote/neu"));
const AngeboteÜbersicht = lazy(() => import("./Pages/CMS/Angebote/page"));
const JobBearbeiten = lazy(() => import("./Pages/CMS/Jobs/id"));
const JobAnlegen = lazy(() => import("./Pages/CMS/Jobs/neu"));
const JobÜbersicht = lazy(() => import("./Pages/CMS/Jobs/page"));
const MitarbeiterBearbeiten = lazy(() => import("./Pages/CMS/Mitarbeiter/id"));
const MitarbeiterAnlegen = lazy(() => import("./Pages/CMS/Mitarbeiter/neu"));
const MitarbeiterÜbersicht = lazy(() => import("./Pages/CMS/Mitarbeiter/page"));
const Overview = lazy(() => import("./Pages/CMS/Overview"));
const PartnerBearbeiten = lazy(() => import("./Pages/CMS/Partner/id"));
const PartnerAnlegen = lazy(() => import("./Pages/CMS/Partner/neu"));
const PartnerÜbersicht = lazy(() => import("./Pages/CMS/Partner/page"));
const EinkaufBearbeiten = lazy(() => import("./Pages/Einkauf/id"));
const Einkauf = lazy(() => import("./Pages/Einkauf/page"));
const Home = lazy(() => import("./Pages/Home"));
const Warenlieferung = lazy(() => import("./Pages/Warenlieferung/page"));

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <BrowserRouter>
      <Routes>
        <Route element={<Layout />}>
          <Route path="/">
            <Route index element={<Home />} />
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
              <Route path="Mitarbeiter">
                <Route index element={<MitarbeiterÜbersicht />} />
                <Route path="Neu" element={<MitarbeiterAnlegen />} />
                <Route path=":id" element={<MitarbeiterBearbeiten />} />
              </Route>
              <Route path="Partner">
                <Route index element={<PartnerÜbersicht />} />
                <Route path="Neu" element={<PartnerAnlegen />} />
                <Route path=":id" element={<PartnerBearbeiten />} />
              </Route>
            </Route>
            <Route path="Archiv" element={<Archive />} />
            <Route path="Warenlieferung" element={<Warenlieferung />} />
            <Route path="Einkauf">
              <Route index element={<Einkauf />} />
              <Route path=":id" element={<EinkaufBearbeiten />} />
            </Route>
          </Route>
        </Route>
      </Routes>
    </BrowserRouter>
  </StrictMode>
);
