import { lazy } from "react";
import BackBtn from "../_components/back-button";
const AbteilungForm = lazy(() => import("./_components/form"));

export default function AbteilungAnlegen() {
  return (
    <div>
      <BackBtn href="/CMS/Abteilungen" />
      <h1 className="text-center">Neue Abteilung anglegen</h1>
      <AbteilungForm />
    </div>
  );
}
