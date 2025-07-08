import BackBtn from "../_components/back-button";
import MitarbeiterForm from "./_components/form";

export default function MitarbeiterAnlegen() {
  return (
    <>
      <BackBtn href="/CMS/Mitarbeiter" />
      <h1 className="text-center">Neuen Mitarbeiter anlegen</h1>
      <MitarbeiterForm />
    </>
  );
}
