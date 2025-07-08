import { useParams } from "react-router";
import BackBtn from "../_components/back-button";
import PartnerForm from "./_components/form";

export default function PartnerBearbeiten() {
  const { id } = useParams();

  return (
    <>
      <BackBtn href="/CMS/Partner" />
      <h1 className="text-center">Partner bearbeiten</h1>
      <PartnerForm id={id} />
    </>
  );
}
