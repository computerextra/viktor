import { lazy } from "react";
import BackBtn from "../_components/back-button";
const MitarbeiterForm = lazy(() => import("./_components/form"));

export default function MitarbeiterAnlegen() {
  return (
    <>
      <BackBtn href="/CMS/Mitarbeiter" />
      <h1 className="text-center">Neuen Mitarbeiter anlegen</h1>
      <MitarbeiterForm />
    </>
  );
}
