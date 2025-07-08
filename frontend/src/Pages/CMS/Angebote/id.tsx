import { lazy } from "react";
import { useParams } from "react-router";
import BackBtn from "../_components/back-button";

const AngebotForm = lazy(() => import("./_components/form"));

export default function AngebotBearbeiten() {
  const { id } = useParams();

  return (
    <>
      <BackBtn href="/CMS/Angebote" />
      <h1 className="text-center">Angebot bearbeiten</h1>
      <AngebotForm id={id} />
    </>
  );
}
