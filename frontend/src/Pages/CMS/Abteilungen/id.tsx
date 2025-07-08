import { GetAbteilung, type GetAbteilungeRes } from "@/api/Abteilungen";
import { useEffect, useState } from "react";
import { useParams } from "react-router";
import BackBtn from "../_components/back-button";
import AbteilungForm from "./_components/form";

export default function AbteilungBearbeiten() {
  const { id } = useParams();
  const [Abteilung, setAbteilung] = useState<GetAbteilungeRes | undefined>(
    undefined
  );
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const res = await GetAbteilung(id);
      setAbteilung(res);
      setLoading(false);
    })();
  }, [id]);

  if (loading) return <>Lädt...</>;

  return (
    <>
      <BackBtn href="/CMS/Abteilungen" />
      <h1 className="text-center">Abteilung: {Abteilung?.name} bearbeiten</h1>
      <AbteilungForm id={id} />
    </>
  );
}
