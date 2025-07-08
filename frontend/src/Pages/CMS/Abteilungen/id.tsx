import { lazy } from "react";
import { useParams } from "react-router";
import BackBtn from "../_components/back-button";

const AbteilungForm = lazy(() => import("./_components/form"));

export default function AbteilungBearbeiten() {
  const { id } = useParams();

  return (
    <>
      <BackBtn href="/CMS/Abteilungen" />
      <h1 className="text-center">Abteilung bearbeiten</h1>
      <AbteilungForm id={id} />
    </>
  );
}
