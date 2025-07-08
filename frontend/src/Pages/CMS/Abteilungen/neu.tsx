import BackBtn from "../_components/back-button";
import AbteilungForm from "./_components/form";

export default function AbteilungAnlegen() {
  return (
    <div>
      <BackBtn href="/CMS/Abteilungen" />
      <h1 className="text-center">Neue Abteilung anglegen</h1>
      <AbteilungForm />
    </div>
  );
}
