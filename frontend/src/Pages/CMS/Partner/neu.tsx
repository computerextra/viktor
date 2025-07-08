import BackBtn from "../_components/back-button";
import PartnerForm from "./_components/form";

export default function PartnerAnlegen() {
  return (
    <div>
      <BackBtn href="/CMS/Partner" />
      <h1 className="text-center">Neuen Partner anlegen</h1>
      <PartnerForm />
    </div>
  );
}
