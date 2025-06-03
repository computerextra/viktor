import BackButton from "@/components/BackButton";
import MitarbeiterForm from "./_components/mitarbeiter-form";

export default function NeuerMitarbeiter() {
  return (
    <>
      <BackButton href="/Mitarbeiter" />
      <h1 className="text-center">Neuen Mitarbeiter anlegen</h1>

      <MitarbeiterForm />
    </>
  );
}
