import { GetAbteilung, type GetAbteilungeRes } from "@/api/Abteilungen";
import { useEffect, useState } from "react";
import { useParams } from "react-router";

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

  return <>{Abteilung?.name}</>;
}
