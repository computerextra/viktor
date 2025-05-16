import "@/index.css";
import Layout from "@/Layout.tsx";
import Anmelden from "@/Pages/Anmelden/Anmelden";
import Archiv from "@/Pages/Archiv/Archiv";
import Home from "@/Pages/Home/Home";
import Lieferant from "@/Pages/Lieferant/Lieferant";
import Mitarbeiter from "@/Pages/Mitarbeiter/Mitarbeiter";
import Suche from "@/Pages/Suche/Suche";
import Werkstatt from "@/Pages/Werkstatt/Werkstatt";

import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import { HashRouter, Route, Routes } from "react-router";
import NeuerMitarbeiter from "./Pages/Mitarbeiter/Neu";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <HashRouter basename={"/"}>
      <Routes>
        <Route path="/" element={<Layout />}>
          <Route index element={<Home />} />
          <Route path="Anmelden">
            <Route index element={<Anmelden />} />
          </Route>
          <Route path="Archiv">
            <Route index element={<Archiv />} />
          </Route>
          <Route path="Mitarbeiter">
            <Route index element={<Mitarbeiter />} />
            <Route path="Neu" element={<NeuerMitarbeiter />} />
          </Route>
          <Route path="Lieferant">
            <Route index element={<Lieferant />} />
          </Route>
          <Route path="Suche">
            <Route index element={<Suche />} />
          </Route>
          <Route path="Werkstatt">
            <Route index element={<Werkstatt />} />
          </Route>
        </Route>
      </Routes>
    </HashRouter>
  </StrictMode>
);
