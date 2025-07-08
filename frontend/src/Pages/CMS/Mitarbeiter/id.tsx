import { useParams } from "react-router";
import BackBtn from "../_components/back-button";
import MitarbeiterForm from "./_components/form";

export default function MitarbeiterBearbeiten() {
  const { id } = useParams();

  return (
    <>
      <BackBtn href="/CMS/Mitarbeiter" />
      <h1 className="text-center">Mitarbeiter bearbeiten</h1>
      <MitarbeiterForm id={id} />
    </>
  );
}
