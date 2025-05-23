import BackButton from "@/components/BackButton";
import { useParams } from "react-router";
import { AnsprechpartnerForm, LieferantenForm } from "./_components/forms";

export function NeuerLieferant() {
  return (
    <>
      <BackButton href="/Lieferant" />
      <h1 className="text-center">Neuen Lieferanten anlegen</h1>
      <LieferantenForm />
    </>
  );
}

export function NeuerAnsprechpartner() {
  const { id } = useParams();
  return (
    <>
      <BackButton href="/Lieferant" />
      <h1 className="text-center">Neuen Ansprechpartner anlegen</h1>
      {id && <AnsprechpartnerForm LieferantenId={parseInt(id)} />}
    </>
  );
}
