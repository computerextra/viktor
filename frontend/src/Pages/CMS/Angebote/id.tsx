import { type AngeboteRes, GetAngebot } from "@/api/Angebote";
import { useEffect, useState } from "react";
import { useParams } from "react-router";
import BackBtn from "../_components/back-button";
import AngebotForm from "./_components/form";

export default function AngebotBearbeiten() {
  const { id } = useParams();
  const [Angebot, setAngebot] = useState<AngeboteRes | undefined>(undefined);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    (async () => {
      if (id == null) return;
      setLoading(true);
      const res = await GetAngebot(id);
      setAngebot(res);
      setLoading(false);
    })();
  }, [id]);

  if (loading) return <>Lädt...</>;

  return (
    <>
      <BackBtn href="/CMS/Angebote" />
      <h1 className="text-center">Angebot: {Angebot?.title} bearbeiten</h1>
      <AngebotForm id={id} />
    </>
  );
}
