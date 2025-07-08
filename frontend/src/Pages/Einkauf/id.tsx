import { useParams } from "react-router";
import EinkaufForm from "./_components/form";

export default function EinkaufBearbeiten() {
  const { id } = useParams();

  if (id == null) return <></>;

  return (
    <>
      <EinkaufForm id={id} />
    </>
  );
}
