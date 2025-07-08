import BackBtn from "../_components/back-button";
import AngebotForm from "./_components/form";

export default function AngebotAnlegen() {
  return (
    <div>
      <BackBtn href="/CMS/Angebote" />
      <h1 className="text-center">Neues Angebot anlegen</h1>
      <AngebotForm />
    </div>
  );
}
