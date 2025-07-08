import { lazy } from "react";
import { useParams } from "react-router";

const EinkaufForm = lazy(() => import("./_components/form"));

export default function EinkaufBearbeiten() {
  const { id } = useParams();

  return (
    <>
      <EinkaufForm id={id} />
    </>
  );
}
