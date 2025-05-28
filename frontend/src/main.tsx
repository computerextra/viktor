import "@/index.css";
import Layout from "@/Layout.tsx";
import Anmelden from "@/Pages/Anmelden/Anmelden";
import Archiv from "@/Pages/Archiv/Archiv";
import Home from "@/Pages/Home/Home";
import Mitarbeiter from "@/Pages/Mitarbeiter/Mitarbeiter";
import Suche from "@/Pages/Suche/Suche";
import Werkstatt from "@/Pages/Werkstatt/Werkstatt";

import LieferantOverview from "@/Pages/Lieferant/Lieferant";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { HashRouter, Route, Routes } from "react-router";
import Abrechnung from "./Pages/Einkauf/Abrechnung";
import Eingabe from "./Pages/Einkauf/Eingabe";
import KanbanBoard from "./Pages/Kanban/Board";
import KanbanBoards from "./Pages/Kanban/Overview";
import {
  AnsprechpartnerDetails,
  LieferantenDetails,
} from "./Pages/Lieferant/Details";
import { EditAnsprechpartner, EditLieferant } from "./Pages/Lieferant/Edit";
import { NeuerAnsprechpartner, NeuerLieferant } from "./Pages/Lieferant/Neu";
import MitarbeiterDetails from "./Pages/Mitarbeiter/Details";
import MitarbeiterBearbeiten from "./Pages/Mitarbeiter/Edit";
import NeuerMitarbeiter from "./Pages/Mitarbeiter/Neu";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <HashRouter basename={"/"}>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="Eingabe" element={<Eingabe />} />
          <Route path="Abrechnung" element={<Abrechnung />} />
          <Route path="Anmelden">
            <Route index element={<Anmelden />} />
          </Route>
          <Route path="Archiv">
            <Route index element={<Archiv />} />
          </Route>
          <Route path="Mitarbeiter">
            <Route index element={<Mitarbeiter />} />
            <Route path="Neu" element={<NeuerMitarbeiter />} />
            <Route path=":id" element={<MitarbeiterDetails />} />
            <Route path=":id/Bearbeiten" element={<MitarbeiterBearbeiten />} />
          </Route>
          <Route path="Lieferant">
            <Route index element={<LieferantOverview />} />
            <Route path="Neu" element={<NeuerLieferant />} />
            <Route path=":id">
              <Route index element={<LieferantenDetails />} />
              <Route path="Bearbeiten" element={<EditLieferant />} />
              <Route path="Neu" element={<NeuerAnsprechpartner />} />
              <Route path=":ap">
                <Route index element={<AnsprechpartnerDetails />} />
                <Route path="Bearbeiten" element={<EditAnsprechpartner />} />
              </Route>
            </Route>
          </Route>
          <Route path="Suche">
            <Route index element={<Suche />} />
          </Route>
          <Route path="Werkstatt">
            <Route index element={<Werkstatt />} />
          </Route>
          <Route path="/Kanban">
            <Route index element={<KanbanBoards />} />
            <Route path=":id" element={<KanbanBoard />} />
          </Route>
        </Route>
      </Routes>
    </HashRouter>
  </StrictMode>
);
