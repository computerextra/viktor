import { lazy } from "react";
import BackBtn from "../_components/back-button";
const PartnerForm = lazy(() => import("./_components/form"));

export default function PartnerAnlegen() {
  return (
    <div>
      <BackBtn href="/CMS/Partner" />
      <h1 className="text-center">Neuen Partner anlegen</h1>
      <PartnerForm />
    </div>
  );
}
